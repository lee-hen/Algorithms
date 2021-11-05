package index_max_pq

import (
	"log"
)

type IndexMaxPQ struct {
	*Heap
}

func NewIndexMaxPQ() *IndexMaxPQ {
	heap := NewHeap(MaxHeapFunc)

	return &IndexMaxPQ{
		heap,
	}
}

func (pq *IndexMaxPQ) IsEmpty() bool {
	return pq.length() == 0
}

func (pq *IndexMaxPQ) Contains(idx int) bool {
	pq.validateIndex(idx)

	if i, ok := pq.inverseIndices[idx]; ok && i != -1 {
		return true
	}

	return false
}

func (pq *IndexMaxPQ) Size() int {
	return pq.length()
}

func (pq *IndexMaxPQ) PQ() []int {
	return pq.indices[1:]
}

func (pq *IndexMaxPQ) Insert(idx int, x float64) {
	if pq.Contains(idx) {
		log.Fatalln("index is already in the priority queue")
	}

	pq.indices = append(pq.indices, idx)
	pq.inverseIndices[idx] = pq.length()
	pq.priorities[idx] = x

	pq.Swim(pq.length())
}

func (pq *IndexMaxPQ) Max() int {
	if pq.length() == 0 {
		log.Fatalln("Priority queue underflow")
	}

	return pq.indices[1]
}

func (pq *IndexMaxPQ) MaxPriority() float64 {
	if pq.length() == 0 {
		log.Fatalln("Priority queue underflow")
	}

	return pq.priorities[pq.indices[1]]
}

func (pq *IndexMaxPQ) DelMax() int {
	if pq.length() == 0 {
		log.Fatalln("Priority queue underflow")
	}

	max := pq.indices[1]

	pq.exchange(1, pq.length())
	pq.indices = pq.indices[:pq.length()]

	delete(pq.inverseIndices, max)
	delete(pq.priorities, max)

	pq.Sink(1)
	return max
}

func (pq *IndexMaxPQ) PriorityOf(idx int) float64 {
	if !pq.Contains(idx) {
		log.Fatalln("index is not in the priority queue")
	}
	return pq.priorities[idx]
}

func (pq *IndexMaxPQ) ChangePriority(idx int, x float64) {
	if !pq.Contains(idx) {
		log.Fatalln("index is not in the priority queue")
	}

	pq.priorities[idx] = x
	pq.Swim(pq.inverseIndices[idx])
	pq.Sink(pq.inverseIndices[idx])
}

func (pq *IndexMaxPQ) IncreasePriority(idx int, x float64) {
	if !pq.Contains(idx) {
		log.Fatalln("index is not in the priority queue")
	}

	if pq.priorities[idx] == x {
		log.Fatalln("Calling increasePriority() with a x equal to the priority in the priority queue")
	} else if pq.priorities[idx] > x {
		log.Fatalln("Calling increasePriority() with a x that is strictly less than the priority in the priority queue")
	}

	pq.priorities[idx] = x
	pq.Swim(pq.inverseIndices[idx])
}

func (pq *IndexMaxPQ) DecreasePriority(idx int, x float64) {
	if !pq.Contains(idx) {
		log.Fatalln("index is not in the priority queue")
	}

	if pq.priorities[idx] == x {
		log.Fatalln("Calling DecreasePriority() with a x equal to the priority in the priority queue")
	} else if pq.priorities[idx] < x {
		log.Fatalln("Calling DecreasePriority() with a x that is strictly greater than the priority in the priority queue")
	}

	pq.priorities[idx] = x
	pq.Sink(pq.inverseIndices[idx])
}

func (pq *IndexMaxPQ) Delete(idx int) {
	if !pq.Contains(idx) {
		log.Fatalln("index is not in the priority queue")
	}

	i := pq.inverseIndices[idx]
	pq.exchange(i, pq.length())
	pq.indices = pq.indices[:pq.length()]

	delete(pq.inverseIndices, idx)
	delete(pq.priorities, idx)

	pq.Sink(1)
}

func (pq *IndexMaxPQ) validateIndex(i int) {
	if i < 0 {
		log.Fatalf("index is negative: %d\n", i)
	}
}

type Heap struct {
	comp   HeapFunc
	//indices  binary heap using 1-based indexing indices[i] = idx
	indices []int
	//inverseIndices inverse of indices -> inverseIndices[indices[i]] = i,  indices[inverseIndices[i]] = idx
	inverseIndices map[int]int

	priorities map[int]float64
}

type HeapFunc func(float64, float64) bool

var MaxHeapFunc = func(a, b float64) bool { return a < b }

func NewHeap(fn HeapFunc) *Heap {
	return &Heap{
		comp:   fn,
		indices: []int{-1},
		inverseIndices: make(map[int]int),
		priorities: make(map[int]float64),
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
	return len(h.indices)-1
}

func (h *Heap) less(i, j int) bool {
	return h.comp(h.priorities[h.indices[i]], h.priorities[h.indices[j]])
}

func (h *Heap) exchange(i, j int) {
	h.indices[i], h.indices[j] = h.indices[j], h.indices[i]

	h.inverseIndices[h.indices[i]] = i
	h.inverseIndices[h.indices[j]] = j
}
