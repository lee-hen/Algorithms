package depth_first_paths

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	"github.com/lee-hen/Algorithms/util"

	"fmt"
)

// Proposition A (continued). DFS allows us to provide clients with a path from a given source to any marked vertex in time proportional its length.
// Proof: By induction on the number of vertices visited, it follows that the edgeTo[] array in DepthFirstPaths represents a tree rooted at the source.
// The pathTo() method builds the path in time proportional to its length.

type Paths struct {
	marked map[int]bool  // marked[v] = is there an s-v path?
	edgeTo map[int]int   // edgeTo[v] = last edge on s-v path
	s int                // source vertex
}

func DepthFirstPaths(g *graph.Graph, s int) *Paths {
	search := &Paths{s:s}
	search.edgeTo = make(map[int]int)
	search.marked = make(map[int]bool)
	search.dfs(g, s)
	return search
}

func (search *Paths) dfs(g *graph.Graph, v int) {
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
	search.validateVertex(v)
	return search.marked[v]
}

// PathTo
// Returns a path between the source vertex s and vertex v, or
// nil if no such path.
func (search *Paths) PathTo(v int) []int {
	search.validateVertex(v)

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

func (search *Paths) validateVertex(v int) {
	if v < 0 || v >= len(search.marked) {
		panic(fmt.Sprintf("vertex %d is not between 0 and %d", v, v-1))
	}
}
