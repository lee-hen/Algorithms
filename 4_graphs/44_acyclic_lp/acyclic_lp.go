package acyclic_lp

import (
	directedEdge "github.com/lee-hen/Algorithms/4_graphs/22_directed_edge"
	graph "github.com/lee-hen/Algorithms/4_graphs/24_edge_weighted_digraph"
	topological "github.com/lee-hen/Algorithms/4_graphs/27_topological"
	"log"

	"github.com/lee-hen/Algorithms/util"
	"math"
)

// Proposition T. We can solve the longest-paths problem in edge-weighted DAGs in time proportional to E + V.
// Proof: Given a longest-paths problem, create a copy of the given edge-weighted DAG that is identical to the original,
// except that all edge weights are negated. Then the shortest path in this copy is the longest path in the original.
// To transform the solution of the shortest-paths problem to a solution of the longest-paths problem, negate the weights in the solution.
// The running time follows immediately from PROPOSITION S.


type AcyclicLP struct {
	distTo []float64
	edgeTo map[int]*directedEdge.Edge
}

// New
// Computes a longest paths tree from {@code s} to every other vertex in
// the directed acyclic graph g.
func New(g *graph.EdgeWeightedDigraph, s int) *AcyclicLP {
	lp := AcyclicLP{}

	lp.distTo = make([]float64, g.V, g.V)
	lp.edgeTo = make(map[int]*directedEdge.Edge)

	for v := 0; v < g.V; v++ {
		lp.distTo[v] = -math.MaxFloat64
	}
	lp.distTo[s] = 0.0

	// relax vertices in order of distance from s
	topologi :=  topological.NewEdgeWeightedDigraph(g)
	if !topologi.HasOrder() {
		log.Fatalln("Digraph is not acyclic.")
	}

	for _, v := range topologi.Order() {
		for _, e := range g.Adj(v) {
			lp.relax(e)
		}
	}

	return &lp
}

// relax edge e
func (lp *AcyclicLP) relax(e *directedEdge.Edge) {
	v, w := e.From(), e.To()
	if lp.distTo[w] < e.Weight() + lp.distTo[v] {
		lp.distTo[w] = e.Weight() + lp.distTo[v]
		lp.edgeTo[w] = e
	}
}

// DistTo
// Returns the weight of a longest path from the source vertex s to vertex v.
func (lp *AcyclicLP) DistTo(v int) float64 {
	return lp.distTo[v]
}

// HasPathTo
// Returns true if there is a path from the source vertex s to vertex v.
func (lp *AcyclicLP) HasPathTo(v int) bool {
	return lp.distTo[v] > -math.MaxFloat64
}

// PathTo
// Returns a longest path from the source vertex s to vertex v.
func (lp *AcyclicLP) PathTo(v int) util.DirectedEdgeStack {
	if !lp.HasPathTo(v) {
		return nil
	}

	path := make(util.DirectedEdgeStack, 0)
	for e := lp.edgeTo[v]; e != nil; e = lp.edgeTo[e.From()] {
		path.Push(e)
	}

	return path
}
