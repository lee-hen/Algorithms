package bellman_ford_sp

import (
	directedEdge "github.com/lee-hen/Algorithms/4_graphs/22_directed_edge"
	graph "github.com/lee-hen/Algorithms/4_graphs/24_edge_weighted_digraph"
	cycle "github.com/lee-hen/Algorithms/4_graphs/25_edge_weighted_directed_cycle"
	"github.com/lee-hen/Algorithms/util"

	"fmt"
	"log"
	"math"
)

// Definition. A negative cycle in an edge-weighted digraph is a directed cycle whose total weight (sum of the weights of its edges) is negative.
// Proposition W. There exists a shortest path from s to v in an edge-weighted digraph if and only if there exists at least one directed path from s to v and no vertex on any directed path from s to v is on a negative cycle. Proof: See discussion above and EXERCISE 4.4.29.
// Proposition X. (Bellman-Ford algorithm) The following method solves the single-source shortest-paths problem from a given source s for any edge-weighted digraph with V vertices and no negative cycles reachable from s: Initialize distTo[s] to 0 and all other distTo[] values to infinity.
// Then, considering the digraph’s edges in any order, relax all edges. Make V such passes. Proof: For any vertex t that is reachable from s consider a specific shortest path from s to t: v0->v1->...->vk, where v0 is s and vk is t. Since there are no negative cycles, such a path exists that is simple, with k no larger than V−1.
// We show by induction on i that after pass i the algorithm computes a shortest path from s to vi. The base case (i = 0) is trivial. Assuming the claim to be true for i, v0->v1->...->vi is a shortest path from s to vi, and distTo[vi] is its length. Now, after pass (i+1), distTo[vi+1] must be
// equal to distTo[vi] plus the weight of vi->vi+1->vi+1 (and the distTo[] values can only decrease); it cannot be less because that is the length of v0->v1->...->vi+1, a shortest path. Thus the algorithm computes a shortest path from s to vi+1 after pass (i+1).

// Proposition W (continued). The Bellman-Ford algorithm takes time proportional to EV and extra space proportional to V. Proof: Each of the V passes relaxes E edges.
// Proposition Y. The queue-based implementation of the Bellman-Ford algorithm solves the single-source shortest-paths problem from a given source s (or finds a negative cycle reachable from s) for any edge-weighted digraph with E edges and V vertices, in time proportional to EV and extra space proportional to V, in the worst case.
// Proof: If there is no negative cycle reachable from s, the algorithm terminates after relaxations corresponding to the (V–1)st pass of the generic algorithm described in PROPOSITION X (since all shortest paths have fewer than V–1 edges).
// If there does exist a negative cycle reachable from s, the queue never empties. After relaxations corresponding to the Vth pass of the generic algorithm described in PROPOSITION X the edgeTo[] array has a path with a cycle (connects some vertex w to itself) and that cycle must be negative, since the path from s to the second occurrence of w must be shorter than the path from s to the first occurrence of w for w to be included on the path the second time.
// In the worst case, the algorithm mimics the general algorithm and relaxes all E edges in each of V passes.


const EPSILON = 1E-14

type BellmanFordSP struct {
	// for floating-point precision issues
	distTo []float64                  // distTo[v] = distance  of shortest s->v path
	edgeTo []*directedEdge.Edge // edgeTo[v] = last edge on shortest s->v path
	onQueue map[int]bool              // onQueue[v] = is v currently on the queue?
	queue []int                       // queue of vertices to relax
	cost int                          // number of calls to relax()

	cycle []*directedEdge.Edge
}

// New
// Computes a shortest paths tree from s to every other vertex in
// the edge-weighted digraph g.
func New(g *graph.EdgeWeightedDigraph, s int) *BellmanFordSP {
	sp := BellmanFordSP{}
	sp.distTo = make([]float64, g.V, g.V)
	sp.edgeTo = make([]*directedEdge.Edge, g.V, g.V)
	sp.onQueue = make(map[int]bool)

	for v := 0; v < g.V; v++ {
		sp.distTo[v] = math.MaxFloat64
	}
	sp.distTo[s] = 0.0

	// Bellman-Ford algorithm
	sp.queue = make([]int, 0)
	sp.queue = append(sp.queue, s)
	sp.onQueue[s] = true

	for len(sp.queue) > 0 && !sp.HasNegativeCycle() {
		var v int
		v, sp.queue = sp.queue[0], sp.queue[1:]
		sp.onQueue[v] = false // used for find negative cycles
		sp.relax(g, v)
	}

	if sp.check(g, s) {
		return &sp
	}
	return nil
}

// relax vertex v and put other endpoints on queue if changed
func (sp *BellmanFordSP) relax(g *graph.EdgeWeightedDigraph, v int) {
	for _, e := range g.Adj(v) {
	//for i := len(g.Adj(v))-1; i >= 0; i-- {
	//	e := g.Adj(v)[i]

		w := e.To()
		if e.Weight() + sp.distTo[v] + EPSILON < sp.distTo[w] {
			sp.distTo[w] = e.Weight() + sp.distTo[v]
			sp.edgeTo[w] = e

			if !sp.onQueue[w] {
				sp.queue = append(sp.queue, w)
				sp.onQueue[w] = true
			}
		}

		sp.cost++
		if sp.cost % g.V == 0 {
			// There is a negative cycle reachable from the source if and only if the queue is nonempty after the Vth pass through all the edges.
			// Moreover, the subgraph of edges in our edgeTo[] array must contain a negative cycle.
			// Accordingly, to implement negativeCycle() BellmanFordSP.java builds an edge-weighted digraph from the edges in edgeTo[] and looks for a cycle in that digraph.
			// To find the cycle, it uses EdgeWeightedDirectedCycle.java, a version of DirectedCycle.java from Section 4.3, adapted to work for edge-weighted digraphs.
			// We amortize the cost of this check by performing this check only after every Vth call to relax().
			sp.findNegativeCycle()
			if sp.HasNegativeCycle() { // found a negative cycle
				return
			}
		}
	}
}

// HasNegativeCycle
// Is there a negative cycle reachable from the source vertex s?
// return true if there is a negative cycle reachable from the
// source vertex s, and false otherwise
func (sp *BellmanFordSP) HasNegativeCycle() bool {
	return sp.cycle != nil
}

// NegativeCycle
// Returns a negative cycle reachable from the source vertex {@code s}, or {@code null}
// if there is no such cycle.
func (sp *BellmanFordSP) NegativeCycle() []*directedEdge.Edge {
	return sp.cycle
}

// by finding a cycle in predecessor graph
func (sp *BellmanFordSP) findNegativeCycle()  {
	V := len(sp.edgeTo)
	spt := graph.NewEdgeWeightedDigraph(V)

	for v := 0; v < V; v++ {
		fmt.Println(v, V)
		if sp.edgeTo[v] != nil {
			spt.AddEdge(sp.edgeTo[v])
		}
	}

	finder := cycle.New(spt)
	sp.cycle = finder.Cycle()
}

// DistTo
// Returns the weight of a shortest path from the source vertex s to vertex v.
func (sp *BellmanFordSP) DistTo(v int) float64 {
	return sp.distTo[v]
}

// HasPathTo
// Returns true if there is a path from the source vertex s to vertex v.
func (sp *BellmanFordSP) HasPathTo(v int) bool {
	return sp.distTo[v] < math.MaxFloat64
}

// PathTo
// Returns a shortest path from the source vertex s to vertex v.
func (sp *BellmanFordSP) PathTo(v int) util.DirectedEdgeStack {
	if sp.HasNegativeCycle() {
		log.Fatalln("Negative cost cycle exists")
	}

	if !sp.HasPathTo(v) {
		return nil
	}

	path := make(util.DirectedEdgeStack, 0)
	for e := sp.edgeTo[v]; e != nil; e = sp.edgeTo[e.From()] {
		path.Push(e)
	}

	return path
}

// check optimality conditions: either
// (i) there exists a negative cycle reacheable from s
//     or
// (ii)  for all edges e = v->w:            distTo[w] <= distTo[v] + e.weight()
// (ii') for all edges e = v->w on the SPT: distTo[w] == distTo[v] + e.weight()
func (sp *BellmanFordSP) check(g *graph.EdgeWeightedDigraph, s int) bool {
	// has a negative cycle
	if sp.HasNegativeCycle() {
		weight := 0.0
		for _, e := range sp.NegativeCycle() {
			weight += e.Weight()
		}

		if weight >= 0.0 {
			log.Fatalln("error: weight of negative cycle =", weight)
			return false
		}
	} else { // no negative cycle reachable from source
		// check that distTo[v] and edgeTo[v] are consistent
		if sp.distTo[s] != 0.0 || sp.edgeTo[s] != nil {
			log.Fatalln("distanceTo[s] and edgeTo[s] inconsistent")
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
				log.Fatalln("edge ", e, " on shortest path not tight")
				return false
			}
		}
	}

	fmt.Println("Satisfies optimality conditions")
	fmt.Println()
	return true
}
