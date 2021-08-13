package main

import "fmt"

// Priority queue implementation with an unsorted array

type MaxHeap []string

func (pq *MaxHeap) IsEmpty() bool {
	return len(*pq) == 0
}

func (pq *MaxHeap) Size() int {
	return len(*pq)
}

func (pq *MaxHeap) DelMax() string {
	max := 0
	for i := 1; i < len(*pq); i++  {
		if less((*pq)[max], (*pq)[i]) {
			max = i
		}
	}
	pq.exchange(max, len(*pq)-1)

	maxKey := (*pq)[len(*pq)-1]
	*pq = (*pq)[:len(*pq)-1]
	return maxKey
}

func (pq *MaxHeap) Insert(key string) {
	*pq = append(*pq, key)
}

func (pq *MaxHeap) exchange(i, j int) {
	(*pq)[i], (*pq)[j] = (*pq)[j], (*pq)[i]
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
