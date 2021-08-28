package breadth_first_paths

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	"github.com/lee-hen/Algorithms/util"

	"fmt"
	"math"
)

// Proposition B. For any vertex v reachable from s, BFS computes a shortest path from s to v (no path from s to v has fewer edges).
// Proof: It is easy to prove by induction that the queue always consists of zero or more vertices of distance k from the source,
// followed by zero or more vertices of distance k+1 from the source, for some integer k, starting with k equal to 0. This property implies,
// in particular, that vertices enter and leave the queue in order of their distance from s. When a vertex v enters the queue, no shorter path to v will be found
// before it comes off the queue, and no path to v that is discovered after it comes off the queue can be shorter than v’s tree path length.

// Proposition B (continued). BFS takes time proportional to V+E in the worst case.
// Proof: As for PROPOSITION A (page 531), BFS marks all the vertices connected to s in time proportional to the sum of their degrees.
// If the graph is connected, this sum equals the sum of the degrees of all the vertices, or 2E. Initialzing the marked[] and edgeTo[] arrays takes time proportional to V.

const INFINITY = math.MaxInt32

type Paths struct {
	marked map[int]bool
	edgeTo map[int]int
	distTo map[int]int
}

// BreadthFirstPaths
// Computes the shortest path between the source vertex w
// and every other vertex in the graph g.
func BreadthFirstPaths(g *graph.Graph, s int) *Paths {
	search := &Paths{}
	search.marked = make(map[int]bool)
	search.edgeTo = make(map[int]int)
	search.distTo = make(map[int]int)

	search.bfs(g, s)
	if !search.check(g, s) {
		panic("not satisfied bfs conditions...")
	}
	return search
}

// breadth-first search from a single source
func (search *Paths) bfs(g *graph.Graph, s int) {
	queue := make([]int, 0)
	for v := 0; v < g.V; v++ {
		search.distTo[v] = INFINITY
	}

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

// BreadthFirstPathsMulti
// Computes the shortest path between any one of the source vertices in sources
// and every other vertex in graph g.
func BreadthFirstPathsMulti(g *graph.Graph, sources []int) *Paths {
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
func (search *Paths) bfsMulti(g *graph.Graph, sources []int) {
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

func (search *Paths) check(g *graph.Graph, s int) bool {
	// check that the distance of s = 0
	// check that the distance of s = 0
	if search.distTo[s] != 0 {
		fmt.Println("distance of source", s, " to itself =", search.distTo[s])
		return false
	}

	// check that for each edge v-w dist[w] <= dist[v] + 1
	// provided v is reachable from s
	for v := 0; v < g.V; v++ {
		for _, w := range g.Adj(v) {
			if search.HasPathTo(v) != search.HasPathTo(w) {
				fmt.Println("edge", v, "-", w)
				fmt.Println("hasPathTo(", v, ") = ", search.HasPathTo(v))
				fmt.Println("hasPathTo(", w, ") = ", search.HasPathTo(w))
				return false
			}
			if search.HasPathTo(v) && (search.distTo[w] > search.distTo[v]+1) {
				fmt.Println("edge", v, "-", w)
				fmt.Println("distTo[", v, "] = ", search.distTo[v])
				fmt.Println("distTo[", w, "] = ", search.distTo[w])
				return false
			}
		}
	}

	// check that v = edgeTo[w] satisfies distTo[w] = distTo[v] + 1
	// provided v is reachable from s
	for w := 0; w < g.V; w++ {
		if !search.HasPathTo(w) || w == s {
			continue
		}
		v := search.edgeTo[w]
		if search.distTo[w] != search.distTo[v]+1 {
			fmt.Println("shortest path edge", v, "-", w)
			fmt.Println("distTo[", v, "] = ", search.distTo[v])
			fmt.Println("distTo[", w, "] = ", search.distTo[w])
			return false
		}
	}

	return true
}
