package lazy_prim_mst

import (
	UF "github.com/lee-hen/Algorithms/1_fundamentals/12_uf"
	edge "github.com/lee-hen/Algorithms/4_graphs/21_edge"
	graph "github.com/lee-hen/Algorithms/4_graphs/23_edge_weighted_graph"
	"github.com/lee-hen/Algorithms/util"

	"container/heap"
	"log"
	"math"
)


// Definition. Recall that a spanning tree of a graph is a connected subgraph with no cycles that includes all the vertices.
// A minimum spanning tree (MST) of an edge-weighted graph is a spanning tree whose weight (the sum of the weights of its edges) is no larger than the weight of any other spanning tree.

// Definition. A cut of a graph is a partition of its vertices into two nonempty disjoint sets.
// A crossing edge of a cut is an edge that connects a vertex in one set with a vertex in the other.

// Proposition J. (Cut property) Given any cut in an edge-weighted graph, the crossing edge of minimum weight is in the MST of the graph.
// Proof: Let e be the crossing edge of minimum weight and let T be the MST. The proof is by contradiction: Suppose that T does not contain e.
// Now consider the graph formed by adding e to T. This graph has a cycle that contains e,
// and that cycle must contain at least one other crossing edge—say, f, which has higher weight than e (since e is minimal and all edge weights are different).
// We can get a spanning tree of strictly lower weight by deleting f and adding e, contradicting the assumed minimality of T.

// Proposition K. (Greedy MST algorithm) The following method colors black all edges in the the MST of any connected edge-weighted graph with V vertices: starting with all edges colored gray,
// find a cut with no black edges, color its minimum-weight edge black, and continue until V−1 edges have been colored black.
// Proof: For simplicity, we assume in the discussion that the edge weights are all different, though the proposition is still true when that is not the case (see EXERCISE 4.3.5).
// By the cut property, any edge that is colored black is in the MST. If fewer than V−1 edges are black, a cut with no black edges exists (recall that we assume the graph to be connected).
// Once V−1 edges are black, the black edges form a spanning tree.

// Proposition L. Prim’s algorithm computes the MST of any connected edge-weighted graph. Proof: Immediate from PROPOSITION K.
// The growing tree defines a cut with no black edges; the algorithm takes the crossing edge of minimal weight, so it is successively coloring edges black in accordance with the greedy algorithm.

// Proposition M. The lazy version of Prim’s algorithm uses space proportional to E and time proportional to E log E (in the worst case) to compute the MST of a connected edge-weighted graph with E edges and V vertices.
// Proof: The bottleneck in the algorithm is the number of edge-weight comparisons in the priority-queue methods insert() and delMin().
// The number of edges on the priority queue is at most E, which gives the space bound. In the worst case, the cost of an insertion is ~lg E and the cost to delete the minimum is ~2 lg E (see PROPOSITION O in CHAPTER 2).
// Since at most E edges are inserted and at most E are deleted, the time bound follows.


const floatingPointEpsilon = 1E-12

type LazyPrimMST struct {
	weight float64
	mst []*edge.Edge

	marked map[int]bool
	pq util.EdgeHeap
}

// New
// Compute a minimum spanning tree (or forest) of an edge-weighted graph.
// G the edge-weighted graph
func New(g *graph.EdgeWeightedGraph) *LazyPrimMST {
	lpm := LazyPrimMST{}

	lpm.mst = make([]*edge.Edge, 0)
	lpm.pq = make(util.EdgeHeap, 0)
	lpm.marked = make(map[int]bool)

	for v := 0; v < g.V; v++ {
		if !lpm.marked[v] {
			lpm.prim(g, v)
		}
	}

	lpm.check(g)

	return &lpm
}

// run Prim's algorithm
func (lpm *LazyPrimMST) prim(g *graph.EdgeWeightedGraph, s int) {
	lpm.scan(g, s)

	for !lpm.pq.IsEmpty() { // better to stop when mst has V-1 edges
		e := heap.Pop(&lpm.pq) // smallest edge on pq

		// two endpoints
		v := e.(*edge.Edge).Either()
		w := e.(*edge.Edge).Other(v)

		if lpm.marked[v] && lpm.marked[w] { // lazy, both v and w already scanned
			continue
		}

		lpm.mst = append(lpm.mst, e.(*edge.Edge)) // add e to MST
		lpm.weight += e.(*edge.Edge).Weight()

		if !lpm.marked[v] {   // v becomes part of tree
			lpm.scan(g, v)
		}

		if !lpm.marked[w] {  // w becomes part of tree
			lpm.scan(g, w)
		}
	}
}

// add all edges e incident to v onto pq if the other endpoint has not yet been scanned
func (lpm *LazyPrimMST) scan(g *graph.EdgeWeightedGraph, v int) {
	lpm.marked[v] = true
	for _, e := range g.Adj(v) {
		if !lpm.marked[e.Other(v)] {
			heap.Push(&lpm.pq, e)
		}
	}
}

// Edges
// Returns the edges in a minimum spanning tree (or forest).
func (lpm *LazyPrimMST) Edges() []*edge.Edge {
	return lpm.mst
}

// Weight
// Returns the sum of the edge weights in a minimum spanning tree (or forest).
func (lpm *LazyPrimMST) Weight() float64 {
	return lpm.weight
}

// check optimality conditions (takes time proportional to E V lg* V)
func (lpm *LazyPrimMST) check(g *graph.EdgeWeightedGraph) bool {
	// check weight
	totalWeight := 0.0
	for _, e := range lpm.Edges() {
		totalWeight += e.Weight()
	}

	if math.Abs(totalWeight - lpm.Weight()) > floatingPointEpsilon {
		log.Fatalf("Weight of edges does not equal weight(): %f vs. %f\n", totalWeight, lpm.Weight())
		return false
	}

	// check that it is acyclic
	uf := UF.NewUF(g.V)
	for _, e := range lpm.Edges() {
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
	for _, e := range lpm.Edges() {
		// all edges in MST except e
		uf := UF.NewUF(g.V)
		for _, f := range lpm.mst {
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
