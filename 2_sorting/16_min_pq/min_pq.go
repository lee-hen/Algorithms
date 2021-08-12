package min_pq

import "math"

type MinPQ struct {
	*Heap
}

func NewMinPQ(keys []int) *MinPQ {
	heap := NewHeap(MinHeapFunc, len(keys)+1)
	heap.values[0] = math.MaxInt32

	for i := 0; i < len(keys); i++ {
		heap.values[i+1] = keys[i]
	}

	for k := len(keys) / 2; k >= 1; k-- {
		heap.Sink(k)
	}

	return &MinPQ{
		heap,
	}
}

func (pq *MinPQ) IsEmpty() bool {
	return pq.length() == 0
}

func (pq *MinPQ) Min() int {
	return pq.values[1]
}

func (pq *MinPQ) Insert(x int) {
	pq.values = append(pq.values, x)
	pq.Swim(pq.length())
}

func (pq *MinPQ) Size() int {
	return pq.length()
}

func (pq *MinPQ) DelMin() int {
	min := pq.values[1]
	pq.exchange(1, pq.length())
	pq.values = pq.values[:pq.length()]
	pq.Sink(1)
	return min
}

type Heap struct {
	comp   HeapFunc
	values []int
}

type HeapFunc func(int, int) bool

var MinHeapFunc = func(a, b int) bool { return a > b }

func NewHeap(fn HeapFunc, n int) *Heap {
	return &Heap{
		comp:   fn,
		values: make([]int, n),
	}
}

func (h *Heap) Swim(k int) {
	for k > 1 && h.greater(k/2, k) {
		h.exchange(k, k/2)
		k = k/2
	}
}

func (h *Heap) Sink(k int) {
	for 2*k <= h.length() {
		j := 2*k
		if j < h.length() && h.greater(j, j+1) {
			j++
		}
		if !h.greater(k, j) {
			break
		}
		h.exchange(k, j)
		k = j
	}
}

func (h *Heap) length() int {
	return len(h.values)-1
}

func (h *Heap) greater(i, j int) bool {
	return h.comp(h.values[i], h.values[j])
}

func (h *Heap) exchange(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}

func (h *Heap) isMinHeap() bool {
	if h.values[0] != math.MaxInt32 {
		return false
	}

	return h.isMinHeapOrdered(1)
}

func (h *Heap) isMinHeapOrdered(k int) bool {
	if k > h.length() {
		return true
	}

	left, right := 2*k, 2*k + 1

	if left  <= h.length() && h.greater(k, left)  {
		return false
	}

	if right <= h.length() && h.greater(k, right) {
		return false
	}

	return h.isMinHeapOrdered(left) && h.isMinHeapOrdered(right)
}
