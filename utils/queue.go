package utils

import "container/heap"

// Normal queue
type Queue[T any] []T

func (q Queue[T]) IsEmpty() bool { return len(q) == 0 }

func (q *Queue[T]) Push(elem T) {
	*q = append(*q, elem)
}

func (q *Queue[T]) Pop() T {
	n := len(*q)
	elem := (*q)[n-1]
	*q = (*q)[0 : n-1]
	return elem
}

// Uses container/heap for implementation
// Always pops the element with the lowest priority first
type PriorityQueue[T comparable] []*PQItem[T]

type PQItem[T comparable] struct {
	Value    T
	Priority int
	index    int
}

func (pq PriorityQueue[T]) IsEmpty() bool { return len(pq) == 0 }

func (pq PriorityQueue[T]) Len() int { return len(pq) }

func (pq PriorityQueue[T]) Less(i, j int) bool {
	// We want Pop to give us the lowest priority
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue[T]) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue[T]) Push(x any) {
	n := len(*pq)
	item := x.(*PQItem[T])
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue[T]) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue[T]) PushItem(value T, priority int) {
	i := len(*pq)
	item := &PQItem[T]{Value: value, Priority: priority, index: i}
	heap.Push(pq, item)
}

func (pq *PriorityQueue[T]) PopItem() T {
	return heap.Pop(pq).(*PQItem[T]).Value
}

func NewPriorityQueue[T comparable]() *PriorityQueue[T] {
	pq := &PriorityQueue[T]{}
	heap.Init(pq)
	return pq
}
