package stable_min_pq

import "github.com/lee-hen/Algorithms/util"

type StableMinPQ struct {
	*Heap
	timestamp int
}

func NewStableMinPQ(keys []Tuple) *StableMinPQ {
	heap := NewHeap(MinHeapFunc, len(keys)+1)
	heap.time[0] = 1

	var t = 1
	for i := 0; i < len(keys); i++ {
		heap.items[i+1] = keys[i]
		heap.time[i+1] = t
		t++
	}

	for k := len(keys) / 2; k >= 1; k-- {
		heap.Sink(k)
	}

	return &StableMinPQ{
		heap,
		t,
	}
}

func (pq *StableMinPQ) IsEmpty() bool {
	return pq.length() == 0
}

func (pq *StableMinPQ) Size() int {
	return pq.length()
}

func (pq *StableMinPQ) Min() Tuple {
	return pq.items[1]
}

func (pq *StableMinPQ) Insert(x Tuple) {
	pq.items = append(pq.items, x)
	pq.timestamp++
	pq.time = append(pq.time, pq.timestamp)
	pq.Swim(pq.length())
}

func (pq *StableMinPQ) DelMin() Tuple {
	min := pq.items[1]
	pq.exchange(1, pq.length())
	n := pq.length()
	pq.items = pq.items[:n]
	pq.time = pq.time[:n]
	pq.Sink(1)
	return min
}

type Heap struct {
	comp   HeapFunc
	items []Tuple
	time []int
}

type HeapFunc func(string, string) bool

var MinHeapFunc = func(a, b string) bool { return a > b }

func NewHeap(fn HeapFunc, n int) *Heap {
	return &Heap{
		comp:   fn,
		items: make([]Tuple, n),
		time: make([]int, n),
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
	return len(h.items)-1
}

func (h *Heap) greater(i, j int) bool {
	if h.items[i].Name == h.items[j].Name {
		return util.Greater(h.time[i], h.time[j])
	}

	return h.comp(h.items[i].Name, h.items[j].Name)
}

func (h *Heap) exchange(i, j int) {
	h.items[i], h.items[j] = h.items[j], h.items[i]
	h.time[i], h.time[j] = h.time[j], h.time[i]
}

type Tuple struct {
	Id int
	Name string
}

func (h *Heap) isMinHeap() bool {
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
