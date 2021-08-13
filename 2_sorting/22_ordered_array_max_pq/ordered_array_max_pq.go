package main

import (
	"fmt"
)

// priority queue implementation with an ordered array.

type MaxHeap []string

func (pq *MaxHeap) IsEmpty() bool {
	return len(*pq) == 0
}

func (pq *MaxHeap) Size() int {
	return len(*pq)
}

func (pq *MaxHeap) DelMax() string {
	max := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return max
}

func (pq *MaxHeap) Insert(key string) {
	i := len(*pq) - 1
	*pq = append(*pq, key)

	for i >= 0 && less(key, (*pq)[i]) {
		(*pq)[i+1] = (*pq)[i]
		i--
	}
	(*pq)[i+1] = key
}

func less(x, y string) bool {
	return x < y
}

func main() {
	pq := make(MaxHeap, 0)
	pq.Insert("this")
	pq.Insert("is")
	pq.Insert("a")
	pq.Insert("test")

	for !pq.IsEmpty() {
		fmt.Println(pq.DelMax())
	}
}
