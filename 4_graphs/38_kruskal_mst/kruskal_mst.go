package kruskal_mst

import (
	UF "github.com/lee-hen/Algorithms/1_fundamentals/12_uf"
	edge "github.com/lee-hen/Algorithms/4_graphs/21_edge"
	graph "github.com/lee-hen/Algorithms/4_graphs/23_edge_weighted_graph"

	"log"
	"math"
	"sort"
)

// Proposition O. Kruskal’s algorithm computes the MST of any connected edge-weighted graph.
// Proof: Immediate from PROPOSITION K. If the next edge to be considered does not form a cycle with black edges,
// it crosses a cut defined by the set of vertices connected to one of the edge’s vertices by tree edges (and its complement).
// Since the edge does not create a cycle, it is the only crossing edge seen so far, and since we consider the edges in sorted order,
// it is a crossing edge of minimum weight. Thus, the algorithm is successively taking a minimal-weight crossing edge, in accordance with the greedy algorithm.

// Proposition O. Kruskal’s algorithm computes the MST of any connected edge-weighted graph.
// Proof: Immediate from PROPOSITION K. If the next edge to be considered does not form a cycle with black edges,
// it crosses a cut defined by the set of vertices connected to one of the edge’s vertices by tree edges (and its complement).
// Since the edge does not create a cycle, it is the only crossing edge seen so far, and since we consider the edges in sorted order,
// it is a crossing edge of minimum weight. Thus, the algorithm is successively taking a minimal-weight crossing edge,
// in accordance with the greedy algorithm.

// Proposition N (continued). Kruskal’s algorithm uses space proportional to E and time proportional to E log E (in the worst case)
// to compute the MST of an edge-weighted connected graph with E edges and V vertices.
// Proof: The implementation uses the priority-queue constructor that initializes the priority queue with all the edges,
// at a cost of at most E compares (see SECTION 2.4). After the priority queue is built, the argument is the same as for Prim’s algorithm.
// The number of edges on the priority queue is at most E, which gives the space bound, and the cost per operation is at most 2 lg E compares, which gives the time bound.
// Kruskal’s algorithm also performs up to E find() and V union() operations, but that cost does not contribute to the E log E order of growth of the total running time (see SECTION 1.5).

const floatingPointEpsilon = 1E-12

type KruskalMST struct {
	weight float64
	mst []*edge.Edge
}


// New
// Compute a minimum spanning tree (or forest) of an edge-weighted graph.
//G the edge-weighted graph
func New(g *graph.EdgeWeightedGraph) *KruskalMST {
	km := KruskalMST{}

	// create array of edges, sorted by weight
	edges := make([]*edge.Edge, g.E, g.E)
	for i, e := range g.Edges() {
		edges[i] = e
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].CompareTo(edges[j]) < 0
	})

	// run greedy algorithm
	km.mst = make([]*edge.Edge, 0)
	uf := UF.NewUF(g.V)
	for i := 0; i < g.E && len(km.mst) < g.V - 1; i++ {
		e := edges[i]
		v := e.Either()
		w := e.Other(v)

		// v-w does not create a cycle
		if !uf.Connected(v, w) {
			uf.Union(v, w)  // merge v and w components
			km.mst = append(km.mst, e) // add edge e to mst
			km.weight += e.Weight()
		}
	}

	km.check(g)

	return &km
}

// Edges
// Returns the edges in a minimum spanning tree (or forest).
func (km *KruskalMST) Edges() []*edge.Edge {
	return km.mst
}

// Weight
// Returns the sum of the edge weights in a minimum spanning tree (or forest).
func (km *KruskalMST) Weight() float64 {
	return km.weight
}

// check optimality conditions (takes time proportional to E V lg* V)
func (km *KruskalMST) check(g *graph.EdgeWeightedGraph) bool {
	// check weight
	totalWeight := 0.0
	for _, e := range km.Edges() {
		totalWeight += e.Weight()
	}

	if math.Abs(totalWeight - km.Weight()) > floatingPointEpsilon {
		log.Fatalf("Weight of edges does not equal weight(): %f vs. %f\n", totalWeight, km.Weight())
		return false
	}

	// check that it is acyclic
	uf := UF.NewUF(g.V)
	for _, e := range km.Edges() {
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
	for _, e := range km.Edges() {
		// all edges in MST except e
		uf := UF.NewUF(g.V)
		for _, f := range km.mst {
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
