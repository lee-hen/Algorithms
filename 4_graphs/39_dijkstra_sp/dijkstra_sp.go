package dijkstra_sp

import (
	minPQ "github.com/lee-hen/Algorithms/2_sorting/22_index_min_pq"
	directedEdge "github.com/lee-hen/Algorithms/4_graphs/22_directed_edge"
	graph "github.com/lee-hen/Algorithms/4_graphs/24_edge_weighted_digraph"
	"github.com/lee-hen/Algorithms/util"

	"log"
	"math"
)


// Definition. A shortest path from vertex s to vertex t in an edge-weighted digraph is a directed path from s to t with the property that no other such path has a lower weight.
// Definition. Given an edge-weighted digraph and a designated vertex s, a shortest-paths tree for a source s is a subgraph containing s and all the vertices reachable from s that forms a directed tree rooted at s such that every tree path is a shortest path in the digraph.

// Proposition P. (Shortest-paths optimality conditions) Let G be an edge-weighted digraph, with s a source vertex in G and
// distTo[] a vertex-indexed array of path lengths in G such that, for all v reachable from s, the value of distTo[v] is the
// length of some path from s to v with distTo[v] equal to infinity for all v not reachable from s.
// These values are the lengths of shortest paths if and only if they satisfy distTo[w] <= distTo[v] + e.weight() for each
// edge e from v to w (or, in other words, no edge is eligible). Proof: Suppose that distTo[w] is the length of a shortest path
// from s to w. If distTo[w] > distTo[v] + e.weight() for some edge e from v to w, then e would give a path from s to w (through v)
// of length less than distTo[w], a contradiction. Thus the optimality conditions are necessary. To prove that the optimality
// conditions are sufficient, suppose that w is reachable from s and that s = v0->v1->v2...->vk = w is a shortest path from s to w, of weight OPTsw. For i from 1 to k,
// denote the edge from vi-1 to vi by ei. By the optimality conditions, we have the following sequence of inequalities: Click here to view code image distTo[w] = distTo[vk]   <= distTo[vk-1] + ek.weight()
// distTo[vk-1] <= distTo[vk-2] + ek-1.weight()
// ...
// distTo[v2]   <= distTo[v1]  + e2.weight()
// distTo[v1]   <= distTo[s]   + e1.weight() Collapsing these inequalities and eliminating distTo[s] = 0.0, we have distTo[w] <= e1.weight() + ... + ek.weight() = OPTsw. Now, distTo[w] is the length of some path from s to w, so it cannot be smaller than the length of a shortest path. Thus, we have shown that OPTsw <= distTo[w] <= OPTsw and equality must hold.


// Proposition Q. (Generic shortest-paths algorithm) Initialize distTo[s] to 0 and all other distTo[] values to infinity,
// and proceed as follows: Relax any edge in G, continuing until no edge is eligible. For all vertices w reachable from s,
// the value of distTo[w] after this computation is the length of a shortest path from s to w (and the value of edgeTo[]
// is the last edge on that path). Proof: Relaxing an edge v->w always sets distTo[w] to the length of some path from s
// (and edgeTo[w] to the last edge on that path). For any vertex w reachable from s, some edge on the shortest path to w
// is eligible as long as distTo[w] remains infinite, so the algorithm continues until the distTo[] value of each vertex
// reachable from s is the length of some path to that vertex. For any vertex v for which the shortest path is well-defined,
// throughout the algorithm distTo[v] is the length of some (simple) path from s to v and is strictly monotonically decreasing.
// Thus, it can decrease at most a finite number of times (once for each simple path from s to v). When no edge is eligible, PROPOSITION P applies.

// Proposition R. Dijkstra’s algorithm solves the single-source shortest-paths problem in edge-weighted digraphs with
// nonnegative weights. Proof: If v is reachable from the source, every edge v->w is relaxed exactly once, when v is
// relaxed, leaving distTo[w] <= distTo[v] + e.weight(). This inequality holds until the algorithm completes,
// since distTo[w] can only decrease (any relaxation can only decrease a distTo[] value) and distTo[v] never changes
// (because edge weights are nonnegative and we choose the lowest distTo[] value at each step, no subsequent relaxation
// can set any distTo[] entry to a lower value than distTo[v]). Thus, after all vertices reachable from s have been added to the tree,
// the shortest-paths optimality conditions hold, and PROPOSITION P applies.

// Proposition R (continued). Dijkstra’s algorithm uses extra space proportional to V and time proportional to E log V (in the worst case) to
// solve the single-source shortest paths problem in an edge-weighted digraph with E edges and V vertices.
// Proof: Same as for Prim’s algorithm (see PROPOSITION N).


type DijkstraSP struct {
	distTo []float64  // distTo[v] = distance  of shortest s->v path
	edgeTo map[int]*directedEdge.Edge // edgeTo[v] = last edge on shortest s->v path

	pq *minPQ.IndexMinPQ
}

// New
// Computes a shortest-paths tree from the source vertex {@code s} to every other
// vertex in the edge-weighted digraph G.
func New(g *graph.EdgeWeightedDigraph, s int) *DijkstraSP {
	for _, e := range g.Edges() {
		if e.Weight() < 0 {
			log.Fatalln("edge ", e, "has negative weight")
		}
	}

	sp := DijkstraSP{}

	sp.distTo = make([]float64, g.V, g.V)
	sp.edgeTo = make(map[int]*directedEdge.Edge)

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
			sp.relax(e)
		}
	}

	if sp.check(g, s) {
		return &sp
	}
	return nil
}

// relax edge e and update pq if changed
func (sp *DijkstraSP) relax(e *directedEdge.Edge) {
	v, w := e.From(), e.To()
	if e.Weight() + sp.distTo[v] < sp.distTo[w] {
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
// Returns the weight of a shortest path from the source vertex {@code s} to vertex v.
func (sp *DijkstraSP) DistTo(v int) float64 {
	return sp.distTo[v]
}

// HasPathTo
// Returns true if there is a path from the source vertex {@code s} to vertex v.
func (sp *DijkstraSP) HasPathTo(v int) bool {
	return sp.distTo[v] < math.MaxFloat64
}

// PathTo
// Returns a shortest path from the source vertex s to vertex v.
func (sp *DijkstraSP) PathTo(v int) util.DirectedEdgeStack {
	if !sp.HasPathTo(v) {
		return nil
	}

	path := make(util.DirectedEdgeStack, 0)
	for e := sp.edgeTo[v]; e != nil; e = sp.edgeTo[e.From()]  {
		path.Push(e)
	}

	return path
}

// check optimality conditions:
// (i) for all edges e:            distTo[e.to()] <= distTo[e.from()] + e.weight()
// (ii) for all edge e on the SPT: distTo[e.to()] == distTo[e.from()] + e.weight()
func (sp *DijkstraSP) check(g *graph.EdgeWeightedDigraph, s int) bool {
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
			if sp.distTo[v] + e.Weight() < sp.distTo[w] {
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

		if sp.distTo[v] + e.Weight() != sp.distTo[w] {
			log.Fatalln("edge", e, "on shortest path not tight")
			return false
		}
	}

	return true
}
