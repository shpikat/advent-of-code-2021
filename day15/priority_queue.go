package day15

import (
	"container/heap"
)

type Position struct {
	risk, row, column int
}

type PriorityQueue []Position

func NewPriorityQueue() (pq PriorityQueue) {
	pq = make([]Position, 0, 16)
	heap.Init(&pq)
	return
}

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].risk < pq[j].risk }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(Position))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}
