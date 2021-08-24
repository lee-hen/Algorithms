package eulerian_cycle

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	path "github.com/lee-hen/Algorithms/4_graphs/06_breadth_first_paths"

	"github.com/lee-hen/Algorithms/util"
	"log"
)

// The EulerianCycle class represents a data type
// for finding an Eulerian cycle or path in a graph.
// An Eulerian cycle is a cycle (not necessarily simple) that
// uses every edge in the graph exactly once.

// EulerianCycle
// Eulerian cycle; null if no such cycle
type EulerianCycle struct {
	cycle util.Stack
	*Edge
}

// Edge
// an undirected edge, with a field to indicate whether the edge has already been used
type Edge struct {
	V, W int

	isUsed bool
}

// Other
// returns the other vertex of the edge
func (e *Edge) Other(vertex int) int {
	if vertex != e.V && vertex != e.W {
		log.Fatalln("Illegal endpoint")
	}

	if vertex == e.V {
		return e.W
	}

	return e.V
}

func New(g *graph.Graph) *EulerianCycle {
	euler := &EulerianCycle{}
	if g.E == 0 {
		return euler
	}

	// necessary condition: all vertices have even degree
	// (this test is needed or it might find an Eulerian path instead of cycle)
	for v := 0; v < g.V; v++ {
		if g.Degree(v) % 2 != 0 {
			return euler
		}
	}

	// create local view of adjacency lists, to iterate one vertex at a time
	// the helper Edge data type is used to avoid exploring both copies of an edge v-w
	adj := make([][]*Edge, g.V, g.V)

	for v := 0; v < g.V; v++ {
		selfLoops := 0
		for _, w := range g.Adj(v) {
			// careful with self loops
			if v == w {
				if selfLoops % 2 == 0 {
					e := &Edge{V: v, W: w}
					adj[v] = append(adj[v], e)
					adj[w] = append(adj[w], e)
				}
				selfLoops++
			} else if v < w {
				e := &Edge{V: v, W: w}
				adj[v] = append(adj[v], e)
				adj[w] = append(adj[w], e)
			}
		}
	}

	// initialize stack with any non-isolated vertex
	s := euler.nonIsolatedVertex(g)
	stack := util.Stack{}
	stack.Push(s)

	// greedily search through edges in iterative DFS style
	euler.cycle = util.Stack{}
	for !stack.IsEmpty() {
		v := stack.Pop()
		for len(adj[v]) > 0 {
			var e *Edge
			e, adj[v] = adj[v][0], adj[v][1:]

			if e.isUsed {
				continue
			}

			e.isUsed = true
			stack.Push(v)
			v = e.Other(v)
		}
		// push vertex with no more leaving edges to cycle
		euler.cycle.Push(v)
	}

	// check if all edges are used
	if euler.cycle.Size() != g.E + 1 {
		euler.cycle = nil
	}

	euler.certifySolution(g)
	return euler
}

// Cycle
// Returns the sequence of vertices on an Eulerian cycle.
func (euler *EulerianCycle) Cycle() util.Stack {
	return euler.cycle
}

// HasEulerianCycle
// Returns true if the graph g has a cycle.
func (euler *EulerianCycle) HasEulerianCycle() bool {
	return euler.cycle != nil
}

// returns any non-isolated vertex; -1 if no such vertex
func (euler *EulerianCycle) nonIsolatedVertex(g *graph.Graph) int {
	for v := 0; v < g.V; v++ {
		if g.Degree(v) > 0 {
			return v
		}
	}
	return -1
}

/**************************************************************************
 *
 *  The code below is solely for testing correctness of the data type.
 *
 **************************************************************************/

// Determines whether a graph has an Eulerian cycle using necessary
// and sufficient conditions (without computing the cycle itself):
//    - at least one edge
//    - degree(v) is even for every vertex v
//    - the graph is connected (ignoring isolated vertices)
func (euler *EulerianCycle) satisfiesNecessaryAndSufficientConditions(g *graph.Graph) bool {
	// Condition 0: at least 1 edge
	if g.E == 0 {
		return false
	}

	// Condition 1: degree(v) is even for every vertex
	for v := 0; v < g.V; v++ {
		if g.Degree(v) % 2 != 0 {
			return false
		}
	}

	// Condition 2: graph is connected, ignoring isolated vertices
	s := euler.nonIsolatedVertex(g)
	bfs := path.BreadthFirstPaths(g, s)
	for v := 0; v < g.V; v++ {
		if g.Degree(v) > 0 && !bfs.HasPathTo(v) {
			return false
		}
	}

	return true
}

// check that solution is correct
func (euler *EulerianCycle) certifySolution(g *graph.Graph) bool {

	// internal consistency check
	if euler.HasEulerianCycle() == (euler.Cycle() == nil) {
		return false
	}

	// hashEulerianCycle() returns correct value
	if euler.HasEulerianCycle() != euler.satisfiesNecessaryAndSufficientConditions(g) {
		return false
	}

	// nothing else to check if no Eulerian cycle
	if euler.cycle == nil {
		return true
	}

	// check that cycle() uses correct number of edges
	if euler.cycle.Size() != g.E + 1 {
		return false
	}

	// check that cycle() is a cycle of G
	// TODO

	// check that first and last vertices in cycle() are the same
	first, last := -1, -1
	for _, v := range euler.Cycle() {
		if first == -1 {
			first = v
		}
		last = v
	}

	if first != last {
		return false
	}

	return true
}
