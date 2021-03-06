package main

import (
	"context"
	"fmt"
	"time"

	pbc "napoleon/controller"
	pbl "napoleon/leader"
	pb "napoleon/napolet"
	pbtype "napoleon/types"

	"github.com/golang/protobuf/proto"
	clientv3 "go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func ReadRPC(cli *clientv3.Client, squeue chan Sched) {
	tunnel := make(chan string)
	for i := 0; i < noOfReadRPCs; i++ {
		go TalkToNapolet(cli, tunnel, squeue)
	}
	for {
		start := time.Now()
		mu.RLock()
		length := len(partition.Ipaddress)
		mu.RUnlock()
		for i := 0; i < length; i++ {
			mu.RLock()
			if len(partition.Ipaddress) < i {
				break
			}

			zap.L().Debug("Ip address to pass to read napolet",
				zap.String("Ipaddress", partition.Ipaddress[i]),
			)
			tunnel <- partition.Ipaddress[i]

			mu.RUnlock()
		}
		_ = time.Since(start)
		zap.L().Debug("")
		time.Sleep(5 * time.Second)
		// time.Sleep(5*time.Minute - elapsed)
	}

}

func GetStat(client pb.PingClient) (*pb.Stat, error) {
	zap.L().Debug("Getting Stat from napolet")
	ctx, cancel := context.WithTimeout(context.Background(), ContextTimeout)
	defer cancel()
	stat, err := client.GetStat(ctx, &pb.Dummy{})
	if err != nil {
		zap.L().Warn("Error in calling RPC", zap.Error(err))
		return nil, err
	}
	return stat, nil
}

func CreateConnectionSlave(cli *clientv3.Client, ipaddress string) (*pb.Stat, error) {
	zap.L().Debug("Trying to connect with the napolet",
		zap.String("IPaddress", ipaddress),
	)
	conn, err := grpc.Dial(ipaddress, grpc.WithInsecure())
	if err != nil {
		zap.L().Error("Failed to dial", zap.Error(err))
		return nil, err
	}
	// zap.L().Debug("Connected with the napolet", zap.String("IPaddress", ipaddress))
	defer conn.Close()
	client := pb.NewPingClient(conn)
	stat, err := GetStat(client)
	if err != nil {
		return nil, err
	}
	zap.L().Debug("Connected with the napolet", zap.String("IPaddress", ipaddress))
	return stat, nil
}

func SendStateUpdate(client pbl.LeaderClient, pmid string, smem, scpu uint32) {
	zap.L().Debug("Sending Stat to Leader",
		zap.String("Physical Machine Id", pmid),
		zap.Uint32("Slack Memory", smem),
		zap.Uint32("Slack CPU", scpu),
	)
	ctx, cancel := context.WithTimeout(context.Background(), ContextTimeout)
	defer cancel()
	_, err := client.SendStateUpdate(ctx, &pbl.StateUpdateRequest{PMId: pmid, SlackCpu: scpu, SlackMemory: smem})
	if err != nil {
		zap.L().Error("Error in calling RPC", zap.Error(err))
		return
	}
}

// func RecoverFromClosedConnections() {
// 	if r := recover(); r != nil {
// 		fmt.Println(r)
// 	}
// }

func CreateConnectionLeader(pmid string, smem, scpu uint32) {
	mu.RLock()
	zap.L().Debug("Trying to connect with the Leader",
		zap.String("Physical Machine Id", pmid),
	)
	conn, err := grpc.Dial(leader+":"+fmt.Sprintf("%d", LeaderPort), grpc.WithInsecure())
	mu.RUnlock()
	if err != nil {
		zap.L().Error("Failed to dial", zap.Error(err))
		return
	}
	zap.L().Debug("Connected with the Leader", zap.String("Physical Machine Id", pmid))
	defer conn.Close()
	client := pbl.NewLeaderClient(conn)
	SendStateUpdate(client, pmid, smem, scpu)
}

func InformLeader(stat *pb.Stat, squeue chan Sched) {
	var smem, scpu uint32
	if stat.PM.SlackMemory == 0 || stat.PM.SlackCpu == 0 {
		smem = 0
		scpu = 0
	} else {
		smem = stat.PM.SlackMemory
		scpu = stat.PM.SlackCpu
	}
	CreateConnectionLeader(stat.PM.PMId, smem, scpu)
	if smem == 0 || scpu == 0 {

		value, error := proto.Marshal(stat)
		zap.L().Error("Error in Marshalling", zap.Error(error))
		eleme := Sched{types: 3, method: string(value)}
		squeue <- eleme
	}
}

func TalkToNapolet(cli *clientv3.Client, tunnel chan string, squeue chan Sched) {
	for {
		ipaddress := <-tunnel
		stat, err := CreateConnectionSlave(cli, ipaddress)
		if err != nil {
			zap.L().Warn("Not connected with IPaddress", zap.String("Ipaddress", ipaddress))
			continue
		}
		for _, vm := range stat.VMS {
			zap.L().Debug("Changing the state of VM depending upon Stat from Napolet",
				zap.String("VMId", vm.VMId),
			)
			// for changing state of migrating vm
			value := getKeyValue(GetKeyResp(context.Background(), cli, vm.VMId))
			var temp pbc.VMStatusResponse
			error := proto.Unmarshal([]byte(value), &temp)
			if error != nil {
				zap.L().Error("Unmarshalling Error", zap.Error(error))
			}
			// get src PM
			value = getKeyValue(GetKeyResp(context.Background(), cli, vm.VMId+"_on"))

			if temp.Status == pbc.VMStatusResponse_MIGRATING {
				// change vm to pm mapping
				insertVMToPMMapping(cli, vm.VMId, stat.PM.PMId)
				// delete vm on pm src
				DeleteVMOnPM(cli, vm.VMId, value)
				// add vm on pm del
				insertPMToVMMapping(cli, vm.VMId, stat.PM.PMId)
				// delete vm migrating key
				destpmid := getKeyValue(GetKeyResp(context.Background(), cli, "migrating_"+vm.VMId))
				DeleteVMToPMMappingMigrating(cli, vm.VMId)
				// delete pm on migrating key
				DeletePMToVMMappingMigrating(cli, vm.VMId, destpmid)
				// change state of vm
				changeStateOfVM(cli, vm.VMId, pbc.VMStatusResponse_CREATED)
			}

			if vm.State == "Created" {
				changeStateOfVM(cli, vm.VMId, pbc.VMStatusResponse_CREATED)
			} else if vm.State == "Suspended" {
				changeStateOfVM(cli, vm.VMId, pbc.VMStatusResponse_SUSPENDED)
			}
		}
		InformLeader(stat, squeue)
	}
}

func DeleteVMOnPM(cli *clientv3.Client, vmid, pmkey string) {
	zap.L().Debug("Trying to delete Migrated VM mapping on Source PM",
		zap.String("VMID", vmid),
		zap.String("PMID", pmkey),
	)
	var list pbtype.VMONPM
	key := "on_" + pmkey

	// Because another thread can be changing the pm to vm mapping
	// TODO: LOCKING ON KEY LEVEL SHOULD BE DONE
	pmtovm.Lock()
	defer pmtovm.Unlock()
	resp := GetKeyResp(context.Background(), cli, key)
	if isKeyPresent(resp) {
		value := getKeyValue(resp)
		error := proto.Unmarshal([]byte(value), &list)
		if error != nil {
			zap.L().Error("Unmarshalling Error", zap.Error(error))
		}
		// list.VMId = append(list.VMId, vmkey)
		for i, targetVMId := range list.VMId {
			if targetVMId == vmid {
				// removing the target vmid
				list.VMId = append(list.VMId[:i], list.VMId[i+1:]...)
				break
			}
		}
	}

	out, err := proto.Marshal(&list)
	if err != nil {
		zap.L().Error("Marshalling Error", zap.Error(err))
	}
	value := string(out)
	getResp := InsertKey(cli, key, value)
	zap.L().Debug("The returned response is", zap.Any("Response", getResp))
	zap.L().Info("Key is inserted",
		zap.String("Key", key),
	)
}

func DeleteVMToPMMappingMigrating(cli *clientv3.Client, vmid string) {
	zap.L().Debug("Deleting VM to PM Migrating Relation",
		zap.String("Virtual machine Id", vmid),
	)
	ctx, cancel := context.WithTimeout(context.Background(), ContextTimeout)
	defer cancel()
	key := "migrating_" + vmid
	// count keys about to be deleted
	gresp, err := cli.Get(ctx, key)
	if err != nil {
		zap.L().Error("Get error", zap.Error(err))
	}

	// delete the keys
	dresp, err := cli.Delete(ctx, key)
	if err != nil {
		zap.L().Error("Delete error", zap.Error(err))
	}

	zap.L().Info("Deleted all keys:", zap.Bool("Value", int64(len(gresp.Kvs)) == dresp.Deleted))
}

func DeletePMToVMMappingMigrating(cli *clientv3.Client, vmkey, pmkey string) {
	zap.L().Debug("Deleting relation between PM to VM for migrating",
		zap.String("Physical Machine Id", pmkey),
		zap.String("Virtual machine Id", vmkey),
	)
	var list pbtype.VMONPM
	key := "migrating_" + pmkey
	pmmigrate.Lock()
	defer pmmigrate.Unlock()
	resp := GetKeyResp(context.Background(), cli, key)
	if isKeyPresent(resp) {
		value := getKeyValue(resp)
		error := proto.Unmarshal([]byte(value), &list)
		if error != nil {
			zap.L().Error("Unmarshalling Error", zap.Error(error))
		}
		//list.VMId = append(list.VMId, vmkey)
		for i, targetVMId := range list.VMId {
			if targetVMId == vmkey {
				// removing the target vmid
				list.VMId = append(list.VMId[:i], list.VMId[i+1:]...)
				break
			}
		}
	}

	out, err := proto.Marshal(&list)
	if err != nil {
		zap.L().Error("Marshalling Error", zap.Error(err))
	}
	value := string(out)
	getResp := InsertKey(cli, key, value)
	zap.L().Debug("The returned response is", zap.Any("Response", getResp))
	zap.L().Info("Key is inserted",
		zap.String("Key", key),
	)
}
