package depth_first_directed_paths

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	"github.com/lee-hen/Algorithms/util"
)

type Paths struct {
	marked map[int]bool  // marked[v] = is there an s-v path?
	edgeTo map[int]int   // edgeTo[v] = last edge on s-v path
	s int                // source vertex
}

func DepthFirstDirectedPaths(g *graph.Digraph, s int) *Paths {
	search := &Paths{s:s}
	search.edgeTo = make(map[int]int)
	search.marked = make(map[int]bool)
	search.dfs(g, s)
	return search
}

func (search *Paths) dfs(g *graph.Digraph, v int) {
	search.marked[v] = true

	for i := len(g.Adj(v))-1; i >= 0; i-- {
		w := g.Adj(v)[i]
		if !search.marked[w] {
			search.edgeTo[w] = v
			search.dfs(g, w)
		}
	}
}

// HasPathTo
// Is there a path between the source vertex s and vertex v?
func (search *Paths) HasPathTo(v int) bool {
	return search.marked[v]
}

// PathTo
// Returns a path between the source vertex s and vertex v, or
// nil if no such path.
func (search *Paths) PathTo(v int) []int {
	if !search.HasPathTo(v) {
		return nil
	}

	path := make(util.Stack, 0)
	for x := v; x != search.s; x = search.edgeTo[x]  {
		path.Push(x)
	}

	path.Push(search.s)
	return path
}

