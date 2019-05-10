// This example demonstrates a priority queue built using the heap interface.
package main

import (
	"container/heap"
)

// An Item is something we manage in a priority queue.
type Item struct {
	PMId                    string
	scpu, smemory, priority uint32
	index                   int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	if pq[i].priority != pq[j].priority {
		return pq[i].priority > pq[j].priority
	}
	return pq[i].PMId > pq[j].PMId
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, scpu, smemory uint32) {
	item.scpu = scpu
	item.smemory = smemory
	item.priority = scpu * smemory
	heap.Fix(pq, item.index)
}

// This example creates a PriorityQueue with some items, adds and manipulates an item,
// and then removes the items in priority order.
// func main() {
// 	// Some items and their priorities.
// 	// items := map[string]int{
// 	// 	"banana": 3, "apple": 2, "pear": 4,
// 	// }

// 	// Create a priority queue, put the items in it, and
// 	// establish the priority queue (heap) invariants.
// 	pq := make(PriorityQueue, 0)
// 	// i := 0
// 	// for value, priority := range items {
// 	// 	pq[i] = &Item{
// 	// 		value:    value,
// 	// 		priority: priority,
// 	// 		index:    i,
// 	// 	}
// 	// 	i++
// 	// }
// 	// heap.Init(&pq)

// 	// Insert a new item and then modify its priority.
// 	item := &Item{
// 		scpu:     3,
// 		smemory:  4,
// 		priority: 1,
// 	}
// 	heap.Push(&pq, item)
// 	pq.update(item, 4, 5)

// 	// Take the items out; they arrive in decreasing priority order.
// 	for pq.Len() > 0 {
// 		item := heap.Pop(&pq).(*Item)
// 		fmt.Printf("%.2d:%s ", item.priority, item.scpu)
// 	}
// }
