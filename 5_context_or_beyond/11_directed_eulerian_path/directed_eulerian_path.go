package directed_eulerian_path

import (
	G "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	path "github.com/lee-hen/Algorithms/4_graphs/06_breadth_first_paths"
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	"github.com/lee-hen/Algorithms/util"
)

type DirectedEulerianPath struct {
	path util.Stack
}

// New
// Computes an Eulerian path in the specified digraph, if one exists.
func New(g *graph.Digraph) *DirectedEulerianPath {
	euler := &DirectedEulerianPath{}

	// find vertex from which to start potential Eulerian path:
	// a vertex v with outdegree(v) > indegree(v) if it exits;
	// otherwise a vertex with outdegree(v) > 0
	deficit := 0
	s := euler.nonIsolatedVertex(g)
	for v := 0; v < g.V; v++ {
		if g.OutDegree(v) > g.InDegree(v) {
			deficit += g.OutDegree(v) - g.InDegree(v)
			s = v
		}
	}

	// digraph can't have an Eulerian path
	// (this condition is needed)
	if deficit > 1 {
		return euler
	}

	// special case for digraph with zero edges (has a degenerate Eulerian path)
	if s == -1 {
		s = 0
	}

	// create local view of adjacency lists, to iterate one vertex at a time
	adj := make([][]int, g.V, g.V)

	// initialize stack with any non-isolated vertex
	stack := util.Stack{}
	stack.Push(s)

	// greedily add to cycle, depth-first search style
	euler.path = util.Stack{}
	for !stack.IsEmpty() {
		v := stack.Pop()
		var next int
		for len(adj[v]) > 0 {
			stack.Push(v)
			v = adj[v][next]
			next++
		}
		// push vertex with no more leaving edges to cycle
		euler.path.Push(v)
	}

	// check if all edges are used
	if euler.path.Size() != g.E + 1 {
		euler.path = nil
	}
	euler.check(g)
	return euler
}

// Path
// Returns the sequence of vertices on an Eulerian path.
func (euler *DirectedEulerianPath) Path() util.Stack {
	return euler.path
}

// HasEulerianPath
// Returns true if the graph has an Eulerian path.
func (euler *DirectedEulerianPath) HasEulerianPath() bool {
	return euler.path != nil
}

// returns any non-isolated vertex; -1 if no such vertex
func (euler *DirectedEulerianPath) nonIsolatedVertex(g *graph.Digraph) int {
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

// Determines whether a digraph has an Eulerian path using necessary
// and sufficient conditions (without computing the path itself):
//    - indegree(v) = outdegree(v) for every vertex,
//      except one vertex v may have outdegree(v) = indegree(v) + 1
//      (and one vertex v may have indegree(v) = outdegree(v) + 1)
//    - the graph is connected, when viewed as an undirected graph
//      (ignoring isolated vertices)
func (euler *DirectedEulerianPath) satisfiesNecessaryAndSufficientConditions(g *graph.Digraph) bool {
	if g.E == 0 {
		return true
	}

	// Condition 1: indegree(v) == outdegree(v) for every vertex,
	// except one vertex may have outdegree(v) = indegree(v) + 1
	deficit := 0
	for v := 0; v < g.V; v++ {
		if g.OutDegree(v) > g.InDegree(v) {
			deficit += g.OutDegree(v) - g.InDegree(v)
		}
	}
	if deficit > 1 {
		return false
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
func (euler *DirectedEulerianPath) check(g *graph.Digraph) bool {

	// internal consistency check
	if euler.HasEulerianPath() == (euler.Path() == nil) {
		return false
	}

	// hashEulerianPath() returns correct value
	if euler.HasEulerianPath() != euler.satisfiesNecessaryAndSufficientConditions(g) {
		return false
	}

	// nothing else to check if no Eulerian path
	if euler.Path() == nil {
		return true
	}

	// check that path() uses correct number of edges
	if euler.path.Size() != g.E + 1 {
		return false
	}

	// check that path() is a path in G
	// TODO

	return true
}
