// Code generated by protoc-gen-go. DO NOT EDIT.
// source: controller.proto

package controller

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type VMStatusResponse_Status int32

const (
	VMStatusResponse_PENDING        VMStatusResponse_Status = 0
	VMStatusResponse_CREATING       VMStatusResponse_Status = 1
	VMStatusResponse_CREATED        VMStatusResponse_Status = 2
	VMStatusResponse_SUSPENDED      VMStatusResponse_Status = 3
	VMStatusResponse_MIGRATING      VMStatusResponse_Status = 4
	VMStatusResponse_DELETE_PENDING VMStatusResponse_Status = 5
	VMStatusResponse_DELETING       VMStatusResponse_Status = 6
)

var VMStatusResponse_Status_name = map[int32]string{
	0: "PENDING",
	1: "CREATING",
	2: "CREATED",
	3: "SUSPENDED",
	4: "MIGRATING",
	5: "DELETE_PENDING",
	6: "DELETING",
}

var VMStatusResponse_Status_value = map[string]int32{
	"PENDING":        0,
	"CREATING":       1,
	"CREATED":        2,
	"SUSPENDED":      3,
	"MIGRATING":      4,
	"DELETE_PENDING": 5,
	"DELETING":       6,
}

func (x VMStatusResponse_Status) String() string {
	return proto.EnumName(VMStatusResponse_Status_name, int32(x))
}

func (VMStatusResponse_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{5, 0}
}

type VMCreateRequest struct {
	VMName               string   `protobuf:"bytes,1,opt,name=VMName,proto3" json:"VMName,omitempty"`
	Vcpus                uint32   `protobuf:"varint,2,opt,name=vcpus,proto3" json:"vcpus,omitempty"`
	Memory               uint32   `protobuf:"varint,3,opt,name=memory,proto3" json:"memory,omitempty"`
	Storage              uint32   `protobuf:"varint,4,opt,name=storage,proto3" json:"storage,omitempty"`
	ImageName            string   `protobuf:"bytes,5,opt,name=ImageName,proto3" json:"ImageName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VMCreateRequest) Reset()         { *m = VMCreateRequest{} }
func (m *VMCreateRequest) String() string { return proto.CompactTextString(m) }
func (*VMCreateRequest) ProtoMessage()    {}
func (*VMCreateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{0}
}

func (m *VMCreateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMCreateRequest.Unmarshal(m, b)
}
func (m *VMCreateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMCreateRequest.Marshal(b, m, deterministic)
}
func (m *VMCreateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMCreateRequest.Merge(m, src)
}
func (m *VMCreateRequest) XXX_Size() int {
	return xxx_messageInfo_VMCreateRequest.Size(m)
}
func (m *VMCreateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VMCreateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VMCreateRequest proto.InternalMessageInfo

func (m *VMCreateRequest) GetVMName() string {
	if m != nil {
		return m.VMName
	}
	return ""
}

func (m *VMCreateRequest) GetVcpus() uint32 {
	if m != nil {
		return m.Vcpus
	}
	return 0
}

func (m *VMCreateRequest) GetMemory() uint32 {
	if m != nil {
		return m.Memory
	}
	return 0
}

func (m *VMCreateRequest) GetStorage() uint32 {
	if m != nil {
		return m.Storage
	}
	return 0
}

func (m *VMCreateRequest) GetImageName() string {
	if m != nil {
		return m.ImageName
	}
	return ""
}

type VMCreateResponse struct {
	VMId                 string   `protobuf:"bytes,1,opt,name=VMId,proto3" json:"VMId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VMCreateResponse) Reset()         { *m = VMCreateResponse{} }
func (m *VMCreateResponse) String() string { return proto.CompactTextString(m) }
func (*VMCreateResponse) ProtoMessage()    {}
func (*VMCreateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{1}
}

func (m *VMCreateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMCreateResponse.Unmarshal(m, b)
}
func (m *VMCreateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMCreateResponse.Marshal(b, m, deterministic)
}
func (m *VMCreateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMCreateResponse.Merge(m, src)
}
func (m *VMCreateResponse) XXX_Size() int {
	return xxx_messageInfo_VMCreateResponse.Size(m)
}
func (m *VMCreateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VMCreateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VMCreateResponse proto.InternalMessageInfo

func (m *VMCreateResponse) GetVMId() string {
	if m != nil {
		return m.VMId
	}
	return ""
}

type VMDeleteRequest struct {
	VMId                 string   `protobuf:"bytes,1,opt,name=VMId,proto3" json:"VMId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VMDeleteRequest) Reset()         { *m = VMDeleteRequest{} }
func (m *VMDeleteRequest) String() string { return proto.CompactTextString(m) }
func (*VMDeleteRequest) ProtoMessage()    {}
func (*VMDeleteRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{2}
}

func (m *VMDeleteRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMDeleteRequest.Unmarshal(m, b)
}
func (m *VMDeleteRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMDeleteRequest.Marshal(b, m, deterministic)
}
func (m *VMDeleteRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMDeleteRequest.Merge(m, src)
}
func (m *VMDeleteRequest) XXX_Size() int {
	return xxx_messageInfo_VMDeleteRequest.Size(m)
}
func (m *VMDeleteRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VMDeleteRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VMDeleteRequest proto.InternalMessageInfo

func (m *VMDeleteRequest) GetVMId() string {
	if m != nil {
		return m.VMId
	}
	return ""
}

type VMDeleteResponse struct {
	Accepted             bool     `protobuf:"varint,1,opt,name=accepted,proto3" json:"accepted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VMDeleteResponse) Reset()         { *m = VMDeleteResponse{} }
func (m *VMDeleteResponse) String() string { return proto.CompactTextString(m) }
func (*VMDeleteResponse) ProtoMessage()    {}
func (*VMDeleteResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{3}
}

func (m *VMDeleteResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMDeleteResponse.Unmarshal(m, b)
}
func (m *VMDeleteResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMDeleteResponse.Marshal(b, m, deterministic)
}
func (m *VMDeleteResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMDeleteResponse.Merge(m, src)
}
func (m *VMDeleteResponse) XXX_Size() int {
	return xxx_messageInfo_VMDeleteResponse.Size(m)
}
func (m *VMDeleteResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VMDeleteResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VMDeleteResponse proto.InternalMessageInfo

func (m *VMDeleteResponse) GetAccepted() bool {
	if m != nil {
		return m.Accepted
	}
	return false
}

type VMStatusRequest struct {
	VMId                 string   `protobuf:"bytes,1,opt,name=VMId,proto3" json:"VMId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *VMStatusRequest) Reset()         { *m = VMStatusRequest{} }
func (m *VMStatusRequest) String() string { return proto.CompactTextString(m) }
func (*VMStatusRequest) ProtoMessage()    {}
func (*VMStatusRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{4}
}

func (m *VMStatusRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMStatusRequest.Unmarshal(m, b)
}
func (m *VMStatusRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMStatusRequest.Marshal(b, m, deterministic)
}
func (m *VMStatusRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMStatusRequest.Merge(m, src)
}
func (m *VMStatusRequest) XXX_Size() int {
	return xxx_messageInfo_VMStatusRequest.Size(m)
}
func (m *VMStatusRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VMStatusRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VMStatusRequest proto.InternalMessageInfo

func (m *VMStatusRequest) GetVMId() string {
	if m != nil {
		return m.VMId
	}
	return ""
}

type VMStatusResponse struct {
	Vm                   *VMCreateRequest        `protobuf:"bytes,1,opt,name=vm,proto3" json:"vm,omitempty"`
	VMId                 string                  `protobuf:"bytes,2,opt,name=VMId,proto3" json:"VMId,omitempty"`
	Status               VMStatusResponse_Status `protobuf:"varint,3,opt,name=status,proto3,enum=VMStatusResponse_Status" json:"status,omitempty"`
	Ipaddress            string                  `protobuf:"bytes,4,opt,name=ipaddress,proto3" json:"ipaddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                `json:"-"`
	XXX_unrecognized     []byte                  `json:"-"`
	XXX_sizecache        int32                   `json:"-"`
}

func (m *VMStatusResponse) Reset()         { *m = VMStatusResponse{} }
func (m *VMStatusResponse) String() string { return proto.CompactTextString(m) }
func (*VMStatusResponse) ProtoMessage()    {}
func (*VMStatusResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_ed7f10298fa1d90f, []int{5}
}

func (m *VMStatusResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VMStatusResponse.Unmarshal(m, b)
}
func (m *VMStatusResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VMStatusResponse.Marshal(b, m, deterministic)
}
func (m *VMStatusResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VMStatusResponse.Merge(m, src)
}
func (m *VMStatusResponse) XXX_Size() int {
	return xxx_messageInfo_VMStatusResponse.Size(m)
}
func (m *VMStatusResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VMStatusResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VMStatusResponse proto.InternalMessageInfo

func (m *VMStatusResponse) GetVm() *VMCreateRequest {
	if m != nil {
		return m.Vm
	}
	return nil
}

func (m *VMStatusResponse) GetVMId() string {
	if m != nil {
		return m.VMId
	}
	return ""
}

func (m *VMStatusResponse) GetStatus() VMStatusResponse_Status {
	if m != nil {
		return m.Status
	}
	return VMStatusResponse_PENDING
}

func (m *VMStatusResponse) GetIpaddress() string {
	if m != nil {
		return m.Ipaddress
	}
	return ""
}

func init() {
	proto.RegisterEnum("VMStatusResponse_Status", VMStatusResponse_Status_name, VMStatusResponse_Status_value)
	proto.RegisterType((*VMCreateRequest)(nil), "VMCreateRequest")
	proto.RegisterType((*VMCreateResponse)(nil), "VMCreateResponse")
	proto.RegisterType((*VMDeleteRequest)(nil), "VMDeleteRequest")
	proto.RegisterType((*VMDeleteResponse)(nil), "VMDeleteResponse")
	proto.RegisterType((*VMStatusRequest)(nil), "VMStatusRequest")
	proto.RegisterType((*VMStatusResponse)(nil), "VMStatusResponse")
}

func init() { proto.RegisterFile("controller.proto", fileDescriptor_ed7f10298fa1d90f) }

var fileDescriptor_ed7f10298fa1d90f = []byte{
	// 399 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0xcf, 0x4a, 0xf3, 0x40,
	0x14, 0xc5, 0x9b, 0xb4, 0x4d, 0x93, 0xdb, 0xaf, 0xfd, 0xe2, 0x20, 0x12, 0x8a, 0x8b, 0x12, 0x50,
	0xba, 0x0a, 0x5a, 0x9f, 0x40, 0x9a, 0x50, 0x02, 0xa6, 0xc8, 0xb4, 0x66, 0x2b, 0x31, 0x1d, 0x8a,
	0xd0, 0x34, 0x31, 0x33, 0x2d, 0xf8, 0x0a, 0x6e, 0x7c, 0x02, 0xdf, 0x55, 0xe6, 0x4f, 0x1a, 0x52,
	0xc5, 0x5d, 0xce, 0x9d, 0x73, 0x7e, 0x73, 0xe7, 0xde, 0x80, 0x9d, 0xe6, 0x3b, 0x56, 0xe6, 0xdb,
	0x2d, 0x29, 0xbd, 0xa2, 0xcc, 0x59, 0xee, 0x7e, 0x6a, 0xf0, 0x3f, 0x8e, 0x66, 0x25, 0x49, 0x18,
	0xc1, 0xe4, 0x6d, 0x4f, 0x28, 0x43, 0x17, 0x60, 0xc4, 0xd1, 0x22, 0xc9, 0x88, 0xa3, 0x8d, 0xb5,
	0x89, 0x85, 0x95, 0x42, 0xe7, 0xd0, 0x3d, 0xa4, 0xc5, 0x9e, 0x3a, 0xfa, 0x58, 0x9b, 0x0c, 0xb0,
	0x14, 0xdc, 0x9d, 0x91, 0x2c, 0x2f, 0xdf, 0x9d, 0xb6, 0x28, 0x2b, 0x85, 0x1c, 0xe8, 0x51, 0x96,
	0x97, 0xc9, 0x86, 0x38, 0x1d, 0x71, 0x50, 0x49, 0x74, 0x09, 0x56, 0x98, 0x25, 0x1b, 0x22, 0xae,
	0xe8, 0x8a, 0x2b, 0xea, 0x82, 0x7b, 0x0d, 0x76, 0xdd, 0x10, 0x2d, 0xf2, 0x1d, 0x25, 0x08, 0x41,
	0x27, 0x8e, 0xc2, 0xb5, 0xea, 0x47, 0x7c, 0xbb, 0x57, 0xbc, 0x71, 0x9f, 0x6c, 0x49, 0xdd, 0xf8,
	0x6f, 0x36, 0x8f, 0xe3, 0x2a, 0x9b, 0xc2, 0x8d, 0xc0, 0x4c, 0xd2, 0x94, 0x14, 0x8c, 0x48, 0xaf,
	0x89, 0x8f, 0x5a, 0x62, 0x97, 0x2c, 0x61, 0x7b, 0xfa, 0x17, 0xf6, 0x43, 0xe7, 0xdc, 0xca, 0xa7,
	0xb8, 0x63, 0xd0, 0x0f, 0x99, 0xb0, 0xf5, 0xa7, 0xb6, 0x77, 0x32, 0x56, 0xac, 0x1f, 0xb2, 0x23,
	0x4a, 0xaf, 0x51, 0xe8, 0x06, 0x0c, 0x2a, 0x38, 0x62, 0x80, 0xc3, 0xa9, 0xe3, 0x9d, 0x82, 0x3d,
	0x25, 0x95, 0x8f, 0x0f, 0xf0, 0xb5, 0x48, 0xd6, 0xeb, 0x92, 0x50, 0x2a, 0x86, 0x6b, 0xe1, 0xba,
	0xe0, 0x16, 0x60, 0x48, 0x3f, 0xea, 0x43, 0xef, 0x31, 0x58, 0xf8, 0xe1, 0x62, 0x6e, 0xb7, 0xd0,
	0x3f, 0x30, 0x67, 0x38, 0xb8, 0x5f, 0x71, 0xa5, 0xf1, 0x23, 0xa1, 0x02, 0xdf, 0xd6, 0xd1, 0x00,
	0xac, 0xe5, 0xd3, 0x92, 0x5b, 0x03, 0xdf, 0x6e, 0x73, 0x19, 0x85, 0x73, 0x2c, 0xad, 0x1d, 0x84,
	0x60, 0xe8, 0x07, 0x0f, 0xc1, 0x2a, 0x78, 0xae, 0x60, 0x5d, 0x0e, 0x13, 0x35, 0xae, 0x8c, 0xe9,
	0x97, 0x06, 0x7a, 0x1c, 0xa1, 0x5b, 0x30, 0xe5, 0x8b, 0xe3, 0x08, 0xfd, 0x78, 0xfe, 0xe8, 0xcc,
	0x3b, 0x5d, 0xab, 0xdb, 0xe2, 0x11, 0xb9, 0x1b, 0x15, 0x69, 0xec, 0x53, 0x44, 0x9a, 0xab, 0x93,
	0x11, 0xf9, 0x3c, 0x15, 0x69, 0xec, 0x4a, 0x44, 0x9a, 0xc3, 0x73, 0x5b, 0x2f, 0x86, 0xf8, 0xd7,
	0xef, 0xbe, 0x03, 0x00, 0x00, 0xff, 0xff, 0xbd, 0x31, 0x09, 0xfa, 0xff, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// VMClient is the client API for VM service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type VMClient interface {
	CreateVM(ctx context.Context, in *VMCreateRequest, opts ...grpc.CallOption) (*VMCreateResponse, error)
	DeleteVM(ctx context.Context, in *VMDeleteRequest, opts ...grpc.CallOption) (*VMDeleteResponse, error)
	StatusVM(ctx context.Context, in *VMStatusRequest, opts ...grpc.CallOption) (*VMStatusResponse, error)
}

type vMClient struct {
	cc *grpc.ClientConn
}

func NewVMClient(cc *grpc.ClientConn) VMClient {
	return &vMClient{cc}
}

func (c *vMClient) CreateVM(ctx context.Context, in *VMCreateRequest, opts ...grpc.CallOption) (*VMCreateResponse, error) {
	out := new(VMCreateResponse)
	err := c.cc.Invoke(ctx, "/VM/CreateVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) DeleteVM(ctx context.Context, in *VMDeleteRequest, opts ...grpc.CallOption) (*VMDeleteResponse, error) {
	out := new(VMDeleteResponse)
	err := c.cc.Invoke(ctx, "/VM/DeleteVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *vMClient) StatusVM(ctx context.Context, in *VMStatusRequest, opts ...grpc.CallOption) (*VMStatusResponse, error) {
	out := new(VMStatusResponse)
	err := c.cc.Invoke(ctx, "/VM/StatusVM", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// VMServer is the server API for VM service.
type VMServer interface {
	CreateVM(context.Context, *VMCreateRequest) (*VMCreateResponse, error)
	DeleteVM(context.Context, *VMDeleteRequest) (*VMDeleteResponse, error)
	StatusVM(context.Context, *VMStatusRequest) (*VMStatusResponse, error)
}

// UnimplementedVMServer can be embedded to have forward compatible implementations.
type UnimplementedVMServer struct {
}

func (*UnimplementedVMServer) CreateVM(ctx context.Context, req *VMCreateRequest) (*VMCreateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVM not implemented")
}
func (*UnimplementedVMServer) DeleteVM(ctx context.Context, req *VMDeleteRequest) (*VMDeleteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteVM not implemented")
}
func (*UnimplementedVMServer) StatusVM(ctx context.Context, req *VMStatusRequest) (*VMStatusResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method StatusVM not implemented")
}

func RegisterVMServer(s *grpc.Server, srv VMServer) {
	s.RegisterService(&_VM_serviceDesc, srv)
}

func _VM_CreateVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMCreateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).CreateVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VM/CreateVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).CreateVM(ctx, req.(*VMCreateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_DeleteVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMDeleteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).DeleteVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VM/DeleteVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).DeleteVM(ctx, req.(*VMDeleteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _VM_StatusVM_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VMStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(VMServer).StatusVM(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/VM/StatusVM",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(VMServer).StatusVM(ctx, req.(*VMStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _VM_serviceDesc = grpc.ServiceDesc{
	ServiceName: "VM",
	HandlerType: (*VMServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVM",
			Handler:    _VM_CreateVM_Handler,
		},
		{
			MethodName: "DeleteVM",
			Handler:    _VM_DeleteVM_Handler,
		},
		{
			MethodName: "StatusVM",
			Handler:    _VM_StatusVM_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "controller.proto",
}
