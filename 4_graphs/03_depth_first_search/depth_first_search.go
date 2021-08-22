package depth_first_search

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"

	"fmt"
)


// Proposition A. DFS marks all the vertices connected to a given source in time proportional to the sum of their degrees.
// Proof: First, we prove that the algorithm marks all the vertices connected to the source s (and no others).
// Every marked vertex is connected to s, since the algorithm finds vertices only by following edges. Now, suppose that some unmarked vertex w is connected to s.
// Since s itself is marked, any path from s to w must have at least one edge from the set of marked vertices to the set of unmarked vertices, say v-x.
// But the algorithm would have discovered x after marking v, so no such edge can exist, a contradiction.
// The time bound follows because marking ensures that each vertex is visited once (taking time proportional to its degree to check marks).

type DepthFirstSearch struct {
	marked map[int]bool // marked[v] = is there an s-v path?
	count int   // number of vertices connected to s
}

func New(g *graph.Graph, s int) *DepthFirstSearch {
	search := &DepthFirstSearch{}
	search.marked = make(map[int]bool)
	search.dfs(g, s)
	return search
}

func (search *DepthFirstSearch) dfs(g *graph.Graph, v int) {
	search.count++
	search.marked[v] = true

	for i := len(g.Adj(v))-1; i >= 0; i-- {
		w := g.Adj(v)[i]
		if !search.marked[w] {
			search.dfs(g, w)
		}
	}
}

func (search *DepthFirstSearch) Marked(v int) bool {
	search.validateVertex(v)
	return search.marked[v]
}

func (search *DepthFirstSearch) Count() int {
	return search.count
}

func (search *DepthFirstSearch) validateVertex(v int) {
	if v < 0 || v >= len(search.marked) {
		panic(fmt.Sprintf("vertex %d is not between 0 and %d", v, v-1))
	}
}
