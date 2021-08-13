package min_pq

type MinPQ struct {
	*Heap
}

func NewMinPQ(keys []CubeSum) *MinPQ {
	heap := NewHeap(MinHeapFunc, len(keys)+1)

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

func (pq *MinPQ) Min() CubeSum {
	return pq.values[1]
}

func (pq *MinPQ) Insert(x CubeSum) {
	pq.values = append(pq.values, x)
	pq.Swim(pq.length())
}

func (pq *MinPQ) Size() int {
	return pq.length()
}

func (pq *MinPQ) DelMin() CubeSum {
	min := pq.values[1]
	pq.exchange(1, pq.length())
	pq.values = pq.values[:pq.length()]
	pq.Sink(1)
	return min
}

type CubeSum struct {
	I, J, Sum int
}

type Heap struct {
	comp   HeapFunc
	values []CubeSum
}

type HeapFunc func(int, int) bool

var MinHeapFunc = func(a, b int) bool { return a > b }

func NewHeap(fn HeapFunc, n int) *Heap {
	return &Heap{
		comp:   fn,
		values: make([]CubeSum, n),
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
	return h.comp(h.values[i].Sum, h.values[j].Sum)
}

func (h *Heap) exchange(i, j int) {
	h.values[i], h.values[j] = h.values[j], h.values[i]
}
