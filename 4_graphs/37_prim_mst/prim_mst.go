package prim_mst

import (
	UF "github.com/lee-hen/Algorithms/1_fundamentals/12_uf"
	minPQ "github.com/lee-hen/Algorithms/2_sorting/22_index_min_pq"
	edge "github.com/lee-hen/Algorithms/4_graphs/21_edge"
	graph "github.com/lee-hen/Algorithms/4_graphs/23_edge_weighted_graph"

	"log"
	"math"
)

// Proposition N. The eager version of Prim’s algorithm uses extra space proportional to V and time proportional to E log V (in the worst case) to compute the MST of a connected edge-weighted graph with E edges and V vertices.
// Proof: The number of edges on the priority queue is at most V, and there are three vertex-indexed arrays, which implies the space bound. The algorithm uses V insert operations, V delete the minimum operations, and (in the worst case) E change priority operations.
// These counts, coupled with the fact that our heap-based implementation of the index priority queue implements all these operations in time proportional to log V (see page 321), imply the time bound.

const floatingPointEpsilon = 1E-12

type PrimMST struct {
	edgeTo map[int]*edge.Edge
	distTo []float64

	marked map[int]bool
	pq *minPQ.IndexMinPQ
}

// New
// Compute a minimum spanning tree (or forest) of an edge-weighted graph.
// G the edge-weighted graph
func New(g *graph.EdgeWeightedGraph) *PrimMST {
	pm := PrimMST{}

	pm.edgeTo = make(map[int]*edge.Edge)
	pm.distTo = make([]float64, g.V, g.V)
	pm.marked = make(map[int]bool)
	pm.pq = minPQ.NewIndexMinPQ(g.V)
	for v := 0; v < g.V; v++ {
		pm.distTo[v] = math.MaxFloat64
	}

	for v := 0; v < g.V; v++ {
		if !pm.marked[v] {
			pm.prim(g, v)
		}
	}

	pm.check(g)
	return &pm
}

// run Prim's algorithm in graph G, starting from vertex s
func (pm *PrimMST) prim(g *graph.EdgeWeightedGraph, s int) {
	pm.distTo[s] = 0.0
	pm.pq.Insert(s, pm.distTo[s])

	for !pm.pq.IsEmpty() {
		v := pm.pq.DelMin()
		pm.scan(g, v)
	}
}

// scan vertex v
func (pm *PrimMST) scan(g *graph.EdgeWeightedGraph, v int) {
	pm.marked[v] = true
	for _, e := range g.Adj(v) {
		w := e.Other(v)
		if pm.marked[w] {
			continue  // v-w is obsolete edge
		}

		if e.Weight() < pm.distTo[w] {
			pm.distTo[w] = e.Weight()
			pm.edgeTo[w] = e

			if pm.pq.Contains(w) {
				pm.pq.DecreasePriority(w, pm.distTo[w])
			} else {
				pm.pq.Insert(w, pm.distTo[w])
			}
		}

	}
}

// Edges
// Returns the edges in a minimum spanning tree (or forest).
func (pm *PrimMST) Edges() []*edge.Edge {
	mst := make([]*edge.Edge, 0)
	for v := range pm.edgeTo {
		e := pm.edgeTo[v]
		if e != nil {
			mst = append(mst, e)
		}
	}
	return mst
}

// Weight
// Returns the sum of the edge weights in a minimum spanning tree (or forest).
func (pm *PrimMST) Weight() float64 {
	weight := 0.0

	for _, e := range pm.Edges() {
		weight += e.Weight()
	}

	return weight
}

// check optimality conditions (takes time proportional to E V lg* V)
func (pm *PrimMST) check(g *graph.EdgeWeightedGraph) bool {
	// check weight
	totalWeight := 0.0
	for _, e := range pm.Edges() {
		totalWeight += e.Weight()
	}

	if math.Abs(totalWeight - pm.Weight()) > floatingPointEpsilon {
		log.Fatalf("Weight of edges does not equal weight(): %f vs. %f\n", totalWeight, pm.Weight())
		return false
	}

	// check that it is acyclic
	uf := UF.NewUF(g.V)
	for _, e := range pm.Edges() {
		v := e.Either()
		w := e.Other(v)

		if uf.Find(v) == uf.Find(w) {
			log.Fatalln("Not a forest")
			return false
		}
		uf.Union(v, w)
	}

	// check that it is a spanning forest
	for _, e := range g.Edges() {
		v := e.Either()
		w := e.Other(v)
		if uf.Find(v) != uf.Find(w) {
			log.Fatalln("Not a forest")
			return false
		}
	}

	// check that it is a minimal spanning forest (cut optimality conditions)
	for _, e := range pm.Edges() {
		// all edges in MST except e
		uf := UF.NewUF(g.V)
		for _, f := range pm.Edges() {
			x := f.Either()
			y := f.Other(x)

			if f.CompareTo(e) != 0 {
				uf.Union(x, y)
			}
		}

		// check that e is min weight edge in crossing cut
		for _, f := range g.Edges() {
			x := f.Either()
			y := f.Other(x)

			if uf.Find(x) != uf.Find(y) {
				if f.Weight() < e.Weight() {
					log.Fatalln("Edge ", f, "violates cut optimality conditions")
					return false
				}
			}
		}
	}

	return true
}
