package acyclic_sp

import (
	directedEdge "github.com/lee-hen/Algorithms/4_graphs/22_directed_edge"
	graph "github.com/lee-hen/Algorithms/4_graphs/24_edge_weighted_digraph"
	topological "github.com/lee-hen/Algorithms/4_graphs/27_topological"
	"log"

	"github.com/lee-hen/Algorithms/util"
	"math"
)

// Proposition S. By relaxing vertices in topological order, we can solve the single-source shortest-paths problem for edge-weighted DAGs in time proportional to E + V.
// Proof: Every edge v->w is relaxed exactly once, when v is relaxed, leaving distTo[w] <= distTo[v] + e.weight().
// This inequality holds until the algorithm completes,
// since distTo[v] never changes (because of the topological order, no edge pointing to v will be processed after v is relaxed)
// and distTo[w] can only decrease (any relaxation can only decrease a distTo[] value).
// Thus, after all vertices reachable from s have been added to the tree, the shortest-paths optimality conditions hold,
// and PROPOSITION Q applies. The time bound is immediate: PROPOSITION G on page 583 tells us that the topological sort takes time proportional to E + V,
// and the second relaxation pass completes the job by relaxing each edge once, again in time proportional to E + V.


type AcyclicSP struct {
	distTo []float64
	edgeTo map[int]*directedEdge.Edge
}

// New
// Computes a shortest-paths tree from the source vertex s to every other
// the directed acyclic graph g.
func New(g *graph.EdgeWeightedDigraph, s int) *AcyclicSP {
	sp := AcyclicSP{}

	sp.distTo = make([]float64, g.V, g.V)
	sp.edgeTo = make(map[int]*directedEdge.Edge)

	for v := 0; v < g.V; v++ {
		sp.distTo[v] = math.MaxFloat64
	}
	sp.distTo[s] = 0.0

	// relax vertices in order of distance from s
	topologi :=  topological.NewEdgeWeightedDigraph(g)
	if !topologi.HasOrder() {
		log.Fatalln("Digraph is not acyclic.")
	}

	for _, v := range topologi.Order() {
		for _, e := range g.Adj(v) {
			sp.relax(e)
		}
	}

	return &sp
}

// relax edge e
func (sp *AcyclicSP) relax(e *directedEdge.Edge) {
	v, w := e.From(), e.To()
	if e.Weight()+sp.distTo[v] < sp.distTo[w] {
		sp.distTo[w] = e.Weight() + sp.distTo[v]
		sp.edgeTo[w] = e
	}
}

// DistTo
// Returns the weight of a shortest path from the source vertex s to vertex v.
func (sp *AcyclicSP) DistTo(v int) float64 {
	return sp.distTo[v]
}

// HasPathTo
// Returns true if there is a path from the source vertex s to vertex v.
func (sp *AcyclicSP) HasPathTo(v int) bool {
	return sp.distTo[v] < math.MaxFloat64
}

// PathTo
// Returns a shortest path from the source vertex s to vertex v.
func (sp *AcyclicSP) PathTo(v int) util.DirectedEdgeStack {
	if !sp.HasPathTo(v) {
		return nil
	}

	path := make(util.DirectedEdgeStack, 0)
	for e := sp.edgeTo[v]; e != nil; e = sp.edgeTo[e.From()] {
		path.Push(e)
	}

	return path
}
