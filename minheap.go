package main

import "container/heap"

type Item struct {
	url   string
	count int
}

type PriorityQueue []*Item

var size int

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].count < pq[j].count
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push pushs item into PriorityQueue
func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
	if pq.Len() > size {
		heap.Pop(pq)
	}
}

// Pop pops the bigest item from PriorityQueue
func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

// toList converts PriorityQueue to list
func (pq *PriorityQueue) toList() []Item {
	list := make([]Item, 0)
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		list = append(list, *item)
	}
	return list
}

// NewMinHeap creates a new min heap
func NewMinHeap(k int) *PriorityQueue {
	size = k
	pq := make(PriorityQueue, 0)
	return &pq
}
