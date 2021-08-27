package transitive_closure

import (
	D "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	dfs "github.com/lee-hen/Algorithms/4_graphs/15_directed_dfs"
)

// Definition. The transitive closure of a digraph G is another digraph with the same set of vertices,
// but with an edge from v to w in the transitive closure if and only if w is reachable from v in G.

var tc []*dfs.DirectedDFS

// TransitiveClosure
// Computes the transitive closure of the digraph G.
func TransitiveClosure(g *D.Digraph) {
	tc = make([]*dfs.DirectedDFS, g.V, g.V)
	for v := 0; v < g.V; v++ {
		tc[v] = dfs.New(g, v)
	}
}

// Reachable
// Is there a directed path from vertex v to vertex w in the digraph?
func Reachable(v, w int) bool {
	return tc[v].Marked(w)
}
