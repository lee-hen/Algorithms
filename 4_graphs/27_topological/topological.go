package topological

import (
	D "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	directedCycle "github.com/lee-hen/Algorithms/4_graphs/19_directed_cycle"
	E "github.com/lee-hen/Algorithms/4_graphs/24_edge_weighted_digraph"
	edgeWeightedDirectedCycle "github.com/lee-hen/Algorithms/4_graphs/25_edge_weighted_directed_cycle"
	depthFirstOrder "github.com/lee-hen/Algorithms/4_graphs/26_depth_first_order"
)

// Proposition E. A digraph has a topological order if and only if it is a DAG.
// Proof: If the digraph has a directed cycle, it has no topological order.
// Conversely, the algorithm that we are about to examine computes a topological order for any given DAG.

// Proposition F. Reverse postorder in a DAG is a topological sort.
// Proof: Consider any edge v->w. One of the following three cases must hold when dfs(v) is called (see the diagram on page 583):
// • dfs(w) has already been called and has returned (w is marked).
// • dfs(w) has not yet been called (w is unmarked), so v->w will cause dfs(w) to be called (and return), either directly or indirectly, before dfs(v) returns.
// • dfs(w) has been called and has not yet returned when dfs(v) is called. The key to the proof is that this case is impossible in a DAG, because the recursive call chain implies a path from w to v and v->w would complete a directed cycle.
// In the two possible cases, dfs(w) is done before dfs(v), so w appears before v in postorder and after v in reverse postorder.
// Thus, each edge v->w points from a vertex earlier in the order to a vertex later in the order, as desired.

// Proposition G. With DFS, we can topologically sort a DAG in time proportional to V+E.
// Proof: Immediate from the code.
// It uses one depth-first search to ensure that the graph has no directed cycles, and another to do the reverse postorder ordering.
// Both involve examining all the edges and all the vertices, and thus take time proportional to V+E.

type Topological struct {
	order []int // topological order
	rank map[int]int // rank[v] = rank of vertex v in order
}

// NewDigraph
// Determines whether the digraph g has a topological order and, if so,
// finds such a topological order.
func NewDigraph(g *D.Digraph) *Topological {
	topological := &Topological{}
	finder := directedCycle.New(g)
	if !finder.HasCycle() {
		dfs := depthFirstOrder.NewDigraph(g)
		topological.order = dfs.ReversePost()
		topological.rank = make(map[int]int)
		var i int
		for _,v := range topological.order {
			topological.rank[v] = i
			i++
		}
	}

	return topological
}

func NewEdgeWeightedDigraph(g *E.EdgeWeightedDigraph) *Topological {
	topological := &Topological{}
	finder := edgeWeightedDirectedCycle.New(g)
	if !finder.HasCycle() {
		dfs := depthFirstOrder.NewEdgeWeightedDigraph(g)
		topological.order = dfs.ReversePost()
	}
	return topological
}

// Order
// Returns a topological order if the digraph has a topologial order,
// nil otherwise.
func (topological *Topological) Order() []int {
	return topological.order
}

// HasOrder
// Does the digraph have a topological order?
// return true if the digraph has a topological order (or equivalently,
// if the digraph is a DAG), and false otherwise
func (topological *Topological) HasOrder() bool {
	return topological.order != nil
}

// Rank
// The the rank of vertex v in the topological order;
// -1 if the digraph is not a DAG
func (topological *Topological) Rank(v int) int {
	if topological.HasOrder() {
		return topological.rank[v]
	} else {
		return -1
	}
}
