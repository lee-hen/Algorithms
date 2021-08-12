package index_min_pq

import (
	"log"
)

// Proposition Q (continued). In an index priority queue of size N, the number of compares required is proportional to at most log N for insert, change priority, delete, and remove the minimum.
// Proof: Immediate from inspection of the code and the fact that all paths in a heap are of length at most ~lg N.

type IndexMinPQ struct {
	*Heap

	maxN int
}

func NewIndexMinPQ(maxN int) *IndexMinPQ {
	heap := NewHeap(MinHeapFunc)

	for idx := 0; idx <= maxN+1; idx++ {
		heap.inverseIndices[idx] = -1
	}

	return &IndexMinPQ{
		heap,
		maxN,
	}
}

func (pq *IndexMinPQ) IsEmpty() bool {
	return pq.length() == 0
}

func (pq *IndexMinPQ) Contains(idx int) bool {
	pq.validateIndex(idx)

	if i, ok := pq.inverseIndices[idx]; ok && i != -1 {
		return true
	}

	return false
}

func (pq *IndexMinPQ) Size() int {
	return pq.length()
}

func (pq *IndexMinPQ) PQ() []int {
	return pq.indices[1:]
}

func (pq *IndexMinPQ) Insert(idx int, x float64) {
	if pq.Contains(idx) {
		log.Fatalln("index is already in the priority queue")
	}

	pq.indices = append(pq.indices, idx)
	pq.inverseIndices[idx] = pq.length()
	pq.priorities[idx] = x

	pq.Swim(pq.length())
}

func (pq *IndexMinPQ) Min() int {
	if pq.length() == 0 {
		log.Fatalln("Priority queue underflow")
	}

	return pq.indices[1]
}

func (pq *IndexMinPQ) MinPriority() float64 {
	if pq.length() == 0 {
		log.Fatalln("Priority queue underflow")
	}

	return pq.priorities[pq.indices[1]]
}

func (pq *IndexMinPQ) DelMin() int {
	if pq.length() == 0 {
		log.Fatalln("Priority queue underflow")
	}

	min := pq.indices[1]

	pq.exchange(1, pq.length())
	pq.indices = pq.indices[:pq.length()]

	delete(pq.inverseIndices, min)
	delete(pq.priorities, min)

	pq.Sink(1)
	return min
}

func (pq *IndexMinPQ) PriorityOf(idx int) float64 {
	if !pq.Contains(idx) {
		log.Fatalln("index is not in the priority queue")
	}
	return pq.priorities[idx]
}

func (pq *IndexMinPQ) ChangePriority(idx int, x float64) {
	if !pq.Contains(idx) {
		log.Fatalln("index is not in the priority queue")
	}

	pq.priorities[idx] = x
	pq.Swim(pq.inverseIndices[idx])
	pq.Sink(pq.inverseIndices[idx])
}

func (pq *IndexMinPQ) DecreasePriority(idx int, x float64) {
	if !pq.Contains(idx) {
		log.Fatalln("index is not in the priority queue")
	}

	if pq.priorities[idx] == x {
		log.Fatalln("Calling increasePriority() with a x equal to the priority in the priority queue")
	} else if pq.priorities[idx] < x {
		log.Fatalln("Calling increasePriority() with a x that is strictly greater than the priority in the priority queue")
	}

	pq.priorities[idx] = x
	pq.Swim(pq.inverseIndices[idx])
}

func (pq *IndexMinPQ) IncreasePriority(idx int, x float64) {
	if !pq.Contains(idx) {
		log.Fatalln("index is not in the priority queue")
	}

	if pq.priorities[idx] == x {
		log.Fatalln("Calling increasePriority() with a x equal to the priority in the priority queue")
	} else if pq.priorities[idx] > x {
		log.Fatalln("Calling increasePriority() with a x that is strictly less than the priority in the priority queue")
	}
	pq.priorities[idx] = x
	pq.Sink(pq.inverseIndices[idx])
}


func (pq *IndexMinPQ) Delete(idx int) {
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

func (pq *IndexMinPQ) validateIndex(i int) {
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

var MinHeapFunc = func(a, b float64) bool { return a > b }

func NewHeap(fn HeapFunc) *Heap {
	return &Heap{
		comp:   fn,
		indices: []int{-1},
		inverseIndices: make(map[int]int),
		priorities: make(map[int]float64),
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
	return len(h.indices)-1
}

func (h *Heap) greater(i, j int) bool {
	return h.comp(h.priorities[h.indices[i]], h.priorities[h.indices[j]])
}

func (h *Heap) exchange(i, j int) {
	h.indices[i], h.indices[j] = h.indices[j], h.indices[i]

	h.inverseIndices[h.indices[i]] = i
	h.inverseIndices[h.indices[j]] = j
}
