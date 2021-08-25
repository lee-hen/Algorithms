package topological

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

