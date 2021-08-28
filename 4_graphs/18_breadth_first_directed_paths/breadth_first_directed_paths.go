package breadth_first_directed_paths

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	"github.com/lee-hen/Algorithms/util"

	"math"
)

const INFINITY = math.MaxInt32

type Paths struct {
	marked map[int]bool
	edgeTo map[int]int
	distTo map[int]int
}

// BreadthFirstDirectedPaths
// Computes the shortest path between the source vertex w
// and every other vertex in the graph g.
func BreadthFirstDirectedPaths(g *graph.Digraph, s int) *Paths {
	search := &Paths{}
	search.marked = make(map[int]bool)
	search.edgeTo = make(map[int]int)
	search.distTo = make(map[int]int)
	for v := 0; v < g.V; v++ {
		search.distTo[v] = INFINITY
	}

	search.bfs(g, s)
	return search
}

// breadth-first search from a single source
func (search *Paths) bfs(g *graph.Digraph, s int) {
	queue := make([]int, 0)
	search.distTo[s] = 0
	search.marked[s] = true
	queue = append(queue, s)

	for len(queue) > 0 {
		var v int
		v, queue = queue[0], queue[1:]
		for i := len(g.Adj(v)) - 1; i >= 0; i-- {
			w := g.Adj(v)[i]
			if !search.marked[w] {
				search.edgeTo[w] = v
				search.distTo[w] = search.distTo[v] + 1
				search.marked[w] = true
				queue = append(queue, w)
			}
		}
	}
}

// BreadthFirstDirectedPathsMulti
// Computes the shortest path between any one of the source vertices in sources
// and every other vertex in graph g.
func BreadthFirstDirectedPathsMulti(g *graph.Digraph, sources []int) *Paths {
	search := &Paths{}
	search.marked = make(map[int]bool)
	search.edgeTo = make(map[int]int)
	search.distTo = make(map[int]int)
	for v := 0; v < g.V; v++ {
		search.distTo[v] = INFINITY
	}
	search.bfsMulti(g, sources)
	return search
}

// breadth-first search from multiple sources
func (search *Paths) bfsMulti(g *graph.Digraph, sources []int) {
	queue := make([]int, 0)

	for _, s := range sources {
		search.marked[s] = true
		search.distTo[s] = 0
		queue = append(queue, s)
	}

	for len(queue) > 0 {
		var v int
		v, queue = queue[0], queue[1:]
		for i := len(g.Adj(v)) - 1; i >= 0; i-- {
			w := g.Adj(v)[i]
			if !search.marked[w] {
				search.edgeTo[w] = v
				search.distTo[w] = search.distTo[v] + 1
				search.marked[w] = true
				queue = append(queue, w)
			}
		}
	}
}

// HasPathTo
// Is there a path between the source vertex s and vertex v?
func (search *Paths) HasPathTo(v int) bool {
	return search.marked[v]
}

// DistTo
// Returns the number of edges in a shortest path between the source vertex s
// (or sources) and vertex v?
func (search *Paths) DistTo(v int) int {
	return search.distTo[v]
}

// PathTo
// Returns a shortest path between the source vertex s (or sources)
// and v, or nil if no such path.
func (search *Paths) PathTo(v int) []int {
	if !search.HasPathTo(v) {
		return nil
	}

	path := make(util.Stack, 0)
	var x int
	for x = v; search.distTo[x] != 0; x = search.edgeTo[x] {
		path.Push(x)
	}

	path.Push(x)
	return path
}
