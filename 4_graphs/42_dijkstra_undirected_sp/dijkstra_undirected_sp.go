package dijkstra_undirected_sp

import (
	minPQ "github.com/lee-hen/Algorithms/2_sorting/22_index_min_pq"
	edge "github.com/lee-hen/Algorithms/4_graphs/21_edge"
	graph "github.com/lee-hen/Algorithms/4_graphs/23_edge_weighted_graph"

	"log"
	"math"
)

type DijkstraSP struct {
	distTo []float64          // distTo[v] = distance  of shortest s->v path
	edgeTo map[int]*edge.Edge // edgeTo[v] = last edge on shortest s->v path

	pq *minPQ.IndexMinPQ
}

// New
// Computes a shortest-paths tree from the source vertex s to every other
// vertex in the edge-weighted digraph G.
func New(g *graph.EdgeWeightedGraph, s int) *DijkstraSP {
	for _, e := range g.Edges() {
		if e.Weight() < 0 {
			log.Fatalln("edge ", e, "has negative weight")
		}
	}

	sp := DijkstraSP{}

	sp.distTo = make([]float64, g.V)
	sp.edgeTo = make(map[int]*edge.Edge)

	for v := 0; v < g.V; v++ {
		sp.distTo[v] = math.MaxFloat64
	}
	sp.distTo[s] = 0.0

	// relax vertices in order of distance from s
	sp.pq = minPQ.NewIndexMinPQ(g.V)
	sp.pq.Insert(s, sp.distTo[s])

	for !sp.pq.IsEmpty() {
		v := sp.pq.DelMin()
		for _, e := range g.Adj(v) {
			sp.relax(e, v)
		}
	}

	// check optimality conditions
	if sp.check(g, s) {
		return &sp
	}
	return nil
}

// relax edge e and update pq if changed
func (sp *DijkstraSP) relax(e *edge.Edge, v int) {
	w := e.Other(v)
	if e.Weight()+sp.distTo[v] < sp.distTo[w] {
		sp.distTo[w] = e.Weight() + sp.distTo[v]
		sp.edgeTo[w] = e

		if sp.pq.Contains(w) {
			sp.pq.DecreasePriority(w, sp.distTo[w])
		} else {
			sp.pq.Insert(w, sp.distTo[w])
		}
	}
}

// DistTo
// Returns the weight of a shortest path from the source vertex s to vertex v.
func (sp *DijkstraSP) DistTo(v int) float64 {
	return sp.distTo[v]
}

// HasPathTo
// Returns true if there is a path from the source vertex s to vertex v.
func (sp *DijkstraSP) HasPathTo(v int) bool {
	return sp.distTo[v] < math.MaxFloat64
}

// PathTo
// Returns a shortest path from the source vertex s to vertex v.
func (sp *DijkstraSP) PathTo(v int) []*edge.Edge {
	if !sp.HasPathTo(v) {
		return nil
	}

	path := make([]*edge.Edge, 0)
	x := v
	for e := sp.edgeTo[v]; e != nil; e = sp.edgeTo[x] {
		path = append(path, e)
		x = e.Other(x)
	}

	return path
}

// check optimality conditions:
// (i) for all edges e:            distTo[e.to()] <= distTo[e.from()] + e.weight()
// (ii) for all edge e on the SPT: distTo[e.to()] == distTo[e.from()] + e.weight()
func (sp *DijkstraSP) check(g *graph.EdgeWeightedGraph, s int) bool {
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
			w := e.Other(v)
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
		if w != e.Either() && w != e.Other(e.Either()) {
			return false
		}

		v := e.Other(w)
		if sp.distTo[v]+e.Weight() != sp.distTo[w] {
			log.Fatalln("edge", e, "on shortest path not tight")
			return false
		}
	}

	return true
}
