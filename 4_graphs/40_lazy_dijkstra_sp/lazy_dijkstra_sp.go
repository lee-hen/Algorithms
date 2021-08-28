package lazy_dijkstra_sp

import (
	directedEdge "github.com/lee-hen/Algorithms/4_graphs/22_directed_edge"
	graph "github.com/lee-hen/Algorithms/4_graphs/24_edge_weighted_digraph"
	"github.com/lee-hen/Algorithms/util"

	"container/heap"
	"log"
	"math"
)

type minHeap struct {
	sp *LazyDijkstraSP
}

func (h *minHeap) Len() int           { return len(h.sp.pq) }
func (h *minHeap) Less(i, j int) bool { return h.sp.less(h.sp.pq[i], h.sp.pq[j]) < 0 }
func (h *minHeap) Swap(i, j int)      { h.sp.pq[i], h.sp.pq[j] = h.sp.pq[j], h.sp.pq[i] }

func (h *minHeap) Push(x interface{}) {
	h.sp.pq = append(h.sp.pq, x.(*directedEdge.Edge))
}

func (h *minHeap) Pop() interface{} {
	old := h.sp.pq
	n := len(old)
	x := old[n-1]
	h.sp.pq = old[0 : n-1]
	return x
}

func (h *minHeap) IsEmpty() bool {
	return len(h.sp.pq) == 0
}

func (h *minHeap) Size() int {
	return len(h.sp.pq)
}

// relax vertex v
func (h *minHeap) relax(g *graph.EdgeWeightedDigraph, v int) {
	sp := h.sp
	sp.marked[v] = true

	for _, e := range g.Adj(v) {
		_, w := e.From(), e.To()
		if e.Weight()+sp.distTo[v] < sp.distTo[w] {
			sp.distTo[w] = e.Weight() + sp.distTo[v]
			sp.edgeTo[w] = e

			heap.Push(h, e)
		}
	}
}

type LazyDijkstraSP struct {
	marked map[int]bool               // has vertex v been relaxed?
	distTo []float64                  // distTo[v] = length of shortest s->v path
	edgeTo map[int]*directedEdge.Edge // edgeTo[v] = last edge on shortest s->v path

	pq []*directedEdge.Edge
}

// New
// single-source shortest path problem from s
func New(g *graph.EdgeWeightedDigraph, s int) *LazyDijkstraSP {
	for _, e := range g.Edges() {
		if e.Weight() < 0 {
			log.Fatalln("edge ", e, "has negative weight")
		}
	}

	h := minHeap{&LazyDijkstraSP{}}
	h.sp.pq = make([]*directedEdge.Edge, 0)

	h.sp.marked = make(map[int]bool)
	h.sp.distTo = make([]float64, g.V, g.V)
	h.sp.edgeTo = make(map[int]*directedEdge.Edge)

	// initialize
	for v := 0; v < g.V; v++ {
		h.sp.distTo[v] = math.MaxFloat64
	}
	h.sp.distTo[s] = 0.0

	h.relax(g, s)

	// run Dijkstra's algorithm
	for !h.IsEmpty() {
		e := heap.Pop(&h) // smallest edge on pq

		// two endpoints
		_, w := e.(*directedEdge.Edge).From(), e.(*directedEdge.Edge).To()

		// lazy, so w might already have been relaxed
		if !h.sp.marked[w] {
			h.relax(g, w)
		}
	}

	h.sp.check(g, s)

	return h.sp
}

func (sp *LazyDijkstraSP) less(e, f *directedEdge.Edge) int {
	dist1 := sp.distTo[e.From()] + e.Weight()
	dist2 := sp.distTo[f.From()] + f.Weight()

	if dist1 == dist2 {
		return 0
	}

	if dist1-dist2 > 0.0 {
		return 1
	}

	return -1
}

// DistTo
// Returns the weight of a shortest path from the source vertex s to vertex v.
func (sp *LazyDijkstraSP) DistTo(v int) float64 {
	return sp.distTo[v]
}

// HasPathTo
// Returns true if there is a path from the source vertex s to vertex v.
func (sp *LazyDijkstraSP) HasPathTo(v int) bool {
	return sp.distTo[v] < math.MaxFloat64
}

// PathTo
// Returns a shortest path from the source vertex s to vertex v.
func (sp *LazyDijkstraSP) PathTo(v int) util.DirectedEdgeStack {
	if !sp.HasPathTo(v) {
		return nil
	}

	path := make(util.DirectedEdgeStack, 0)
	for e := sp.edgeTo[v]; e != nil; e = sp.edgeTo[e.From()] {
		path.Push(e)
	}

	return path
}

// check optimality conditions:
// (i) for all edges e:            distTo[e.to()] <= distTo[e.from()] + e.weight()
// (ii) for all edge e on the SPT: distTo[e.to()] == distTo[e.from()] + e.weight()
func (sp *LazyDijkstraSP) check(g *graph.EdgeWeightedDigraph, s int) bool {
	// check that edge weights are non-negative
	for _, e := range g.Edges() {
		if e.Weight() < 0 {
			log.Fatalln("negative edge weight detected")
			return false
		}
	}

	// check that distTo[v] and edgeTo[v] are consistent
	if sp.distTo[s] != 0.0 || sp.edgeTo[s] != nil {
		log.Fatalln("distTo[s] and edgeTo[s] inconsistent")
		return false
	}

	for v := 0; v < g.V; v++ {
		if v == s {
			continue
		}

		if sp.edgeTo[v] == nil && sp.distTo[v] != math.MaxFloat64 {
			log.Fatalln("distTo[] and edgeTo[] inconsistent")
			return false
		}
	}

	// check that all edges e = v->w satisfy distTo[w] <= distTo[v] + e.weight()
	for v := 0; v < g.V; v++ {
		for _, e := range g.Adj(v) {
			w := e.To()
			if sp.distTo[v]+e.Weight() < sp.distTo[w] {
				log.Fatalln("edge", e, "not relaxed")
				return false
			}
		}
	}

	// check that all edges e = v->w on SPT satisfy distTo[w] == distTo[v] + e.weight()
	for w := 0; w < g.V; w++ {
		if sp.edgeTo[w] == nil {
			continue
		}
		e := sp.edgeTo[w]
		v := e.From()
		if w != e.To() {
			return false
		}

		if sp.distTo[v]+e.Weight() != sp.distTo[w] {
			log.Fatalln("edge", e, "on shortest path not tight")
			return false
		}
	}

	return true
}
