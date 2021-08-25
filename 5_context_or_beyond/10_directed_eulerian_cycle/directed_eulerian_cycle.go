package directed_eulerian_cycle

import (
	G "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	path "github.com/lee-hen/Algorithms/4_graphs/06_breadth_first_paths"
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	"github.com/lee-hen/Algorithms/util"
)

// The DirectedEulerianCycle class represents a data type
// for finding an Eulerian cycle or path in a graph.
// An Eulerian cycle is a cycle (not necessarily simple) that
// uses every edge in the graph exactly once.

// DirectedEulerianCycle
// Eulerian cycle; null if no such cycle
type DirectedEulerianCycle struct {
	cycle util.Stack
}

func New(g *graph.Digraph) *DirectedEulerianCycle {
	euler := &DirectedEulerianCycle{}
	if g.E == 0 {
		return euler
	}

	// necessary condition: indegree(v) = outdegree(v) for each vertex v
	// (without this check, DFS might return a path instead of a cycle)
	for v := 0; v < g.V; v++ {
		if g.OutDegree(v) != g.InDegree(v) {
			return euler
		}
	}

	// create local view of adjacency lists, to iterate one vertex at a time
	adj := make([][]int, g.V, g.V)

	// initialize stack with any non-isolated vertex
	s := euler.nonIsolatedVertex(g)
	stack := util.Stack{}
	stack.Push(s)

	// greedily search through edges in iterative DFS style
	euler.cycle = util.Stack{}
	for !stack.IsEmpty() {
		v := stack.Pop()
		var next int
		for len(adj[v]) > 0 {
			stack.Push(v)
			v = adj[v][next]
			next++
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
func (euler *DirectedEulerianCycle) Cycle() util.Stack {
	return euler.cycle
}

// HasEulerianCycle
// Returns true if the digraph has an Eulerian cycle
func (euler *DirectedEulerianCycle) HasEulerianCycle() bool {
	return euler.cycle != nil
}

// returns any non-isolated vertex; -1 if no such vertex
func (euler *DirectedEulerianCycle) nonIsolatedVertex(g *graph.Digraph) int {
	for v := 0; v < g.V; v++ {
		if g.OutDegree(v) > 0 {
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
func (euler *DirectedEulerianCycle) satisfiesNecessaryAndSufficientConditions(g *graph.Digraph) bool {
	// Condition 0: at least 1 edge
	if g.E == 0 {
		return false
	}

	// Condition 1: degree(v) is even for every vertex
	for v := 0; v < g.V; v++ {
		if g.OutDegree(v) != g.InDegree(v) {
			return false
		}
	}

	h := G.NewGraph(g.V)
	for v := 0; v < g.V; v++ {
		for _, w := range g.Adj(v) {
			h.AddEdge(v, w)
		}
	}

	// check that all non-isolated vertices are conneted
	s := euler.nonIsolatedVertex(g)
	bfs := path.BreadthFirstPaths(h, s)
	for v := 0; v < g.V; v++ {
		if h.Degree(v) > 0 && !bfs.HasPathTo(v) {
			return false
		}
	}

	return true
}

// check that solution is correct
func (euler *DirectedEulerianCycle) certifySolution(g *graph.Digraph) bool {

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

	return true
}
