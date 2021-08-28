package directed_dfs

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
)

type DirectedDFS struct {
	marked map[int]bool // marked[v] = true iff v is reachable from source(s)
	count  int          // number of vertices reachable from source(s)
}

// New
// Computes the vertices in digraph G that are
// reachable from the source vertex s.
func New(g *graph.Digraph, s int) *DirectedDFS {
	search := &DirectedDFS{}
	search.marked = make(map[int]bool)
	search.dfs(g, s)
	return search
}

// Multi
// Computes the vertices in digraph g that are
// connected to any of the source vertices sources.
func Multi(g *graph.Digraph, sources []int) *DirectedDFS {
	search := &DirectedDFS{}
	search.marked = make(map[int]bool)
	for _, v := range sources {
		if !search.marked[v] {
			search.dfs(g, v)
		}
	}
	return search
}

func (search *DirectedDFS) dfs(g *graph.Digraph, v int) {
	search.count++
	search.marked[v] = true

	for i := len(g.Adj(v)) - 1; i >= 0; i-- {
		w := g.Adj(v)[i]
		if !search.marked[w] {
			search.dfs(g, w)
		}
	}
}

func (search *DirectedDFS) Marked(v int) bool {
	return search.marked[v]
}

func (search *DirectedDFS) Count() int {
	return search.count
}
