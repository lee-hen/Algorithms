package max_pq

import (
	"math"
)


// Proposition P. The height of a complete binary tree of size N is  lg N .
// Proof: The stated result is easy to prove by induction or by noting that the height increases by 1 when N is a power of 2.
// Proposition Q. In an N-key priority queue, the heap algorithms require no more than 1 + lg N compares for insert and no more than 2 lg N compares for remove the maximum.
// Proof: By PROPOSITION P, both operations involve moving along a path between the root and the bottom of the heap whose number of links is no more than lg N. The remove the maximum operation requires two compares for each node on the path (except at the bottom): one to find the child with the larger key, the other to decide whether that child needs to be promoted.
// Proposition Q (continued). In an index priority queue of size N, the number of compares required is proportional to at most log N for insert, change priority, delete, and remove the minimum.
// Proof: Immediate from inspection of the code and the fact that all paths in a heap are of length at most ~lg N.

type MaxPQ struct {
	*Heap
}

func NewMaxPQ(keys []int) *MaxPQ {
	heap := NewHeap(MaxHeapFunc, len(keys)+1)
	heap.values[0] = math.MinInt32

	for i := 0; i < len(keys); i++ {
		heap.values[i+1] = keys[i]
	}

	for k := len(keys) / 2; k >= 1; k-- {
		heap.Sink(k)
	}

	return &MaxPQ{
		heap,
	}
}

func (pq *MaxPQ) IsEmpty() bool {
	return pq.length() == 0
}

func (pq *MaxPQ) Max() int {
	return pq.values[1]
}

func (pq *MaxPQ) Insert(x int) {
	pq.values = append(pq.values, x)
	pq.Swim(pq.length())
}

func (pq *MaxPQ) Size() int {
	return pq.length()
}

func (pq *MaxPQ) DelMax() int {
	max := pq.values[1]
	pq.exchange(1, pq.length())
	pq.values = pq.values[:pq.length()]
	pq.Sink(1)
	return max
}

type Heap struct {
	comp   HeapFunc
	values []int
}

type HeapFunc func(int, int) bool

var MaxHeapFunc = func(a, b int) bool { return a < b }

func NewHeap(fn HeapFunc, n int) *Heap {
	return &Heap{
		comp:   fn,
		values: make([]int, n),
	}
}

func (h *Heap) Swim(k int) {
	for k > 1 && h.less(k/2, k) {
		h.exchange(k, k/2)
		k = k/2
	}
}

func (h *Heap) Sink(k int) {
	for 2*k <= h.length() {
		j := 2*k
		if j < h.length() && h.less(j, j+1) {
			j++
		}
		if !h.less(k, j) {
			break
		}
		h.exchange(k, j)
		k = j
	}
}

func (h *Heap) length() int {
	return len(h.values)-1
}

func (h *Heap) less(i, j int) bool {
	return h.comp(h.values[i], h.values[j])
}

func (h *Heap) exchange(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *Heap) isMaxHeap() bool {
	if h.values[0] != math.MinInt32 {
		return false
	}

	return h.isMaxHeapOrdered(1)
}

func (h *Heap) isMaxHeapOrdered(k int) bool {
	if k > h.length() {
		return true
	}

	left, right := 2*k, 2*k + 1

	if left  <= h.length() && h.less(k, left)  {
		return false
	}

	if right <= h.length() && h.less(k, right) {
		return false
	}

	return h.isMaxHeapOrdered(left) && h.isMaxHeapOrdered(right)
}
