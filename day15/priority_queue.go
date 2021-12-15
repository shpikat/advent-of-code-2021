package day15

import (
	"container/heap"
)

type Position struct {
	risk, row, column int
}

type Item struct {
	Position
	index int
}

type PriorityQueue []*Item

func NewPriorityQueue() (pq PriorityQueue) {
	pq = make([]*Item, 0, 16)
	heap.Init(&pq)
	return
}

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].risk < pq[j].risk }
func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	last := len(*pq)
	*pq = append(*pq, &Item{x.(Position), last})
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*pq = old[0 : n-1]
	return x.Position
}

func (pq *PriorityQueue) UpdateOrAdd(p Position) {
	for _, item := range *pq {
		if item.row == p.row && item.column == p.column {
			item.risk = p.risk
			heap.Fix(pq, item.index)
			return
		}
	}

	heap.Push(pq, p)
}
