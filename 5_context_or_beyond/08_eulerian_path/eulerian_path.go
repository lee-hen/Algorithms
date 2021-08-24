package eulerian_path

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	path "github.com/lee-hen/Algorithms/4_graphs/06_breadth_first_paths"
	"github.com/lee-hen/Algorithms/util"

	"log"
)

// The EulerianPath class represents a data type
// for finding an Eulerian path in a graph.
// An Eulerian path is a path (not necessarily simple) that
// uses every edge in the graph exactly once.

type EulerianPath struct {
	path util.Stack
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

// New
// Computes an Eulerian path in the specified graph, if one exists.
func New(g *graph.Graph) *EulerianPath {
	euler := &EulerianPath{}

	// find vertex from which to start potential Eulerian path:
	// a vertex v with odd degree(v) if it exits;
	// otherwise a vertex with degree(v) > 0
	oddDegreeVertices := 0
	s := euler.nonIsolatedVertex(g)
	for v := 0; v < g.V; v++ {
		if g.Degree(v) % 2 == 0 {
			oddDegreeVertices++
			s = v
		}
	}

	// graph can't have an Eulerian path
	// (this condition is needed for correctness)
	if oddDegreeVertices > 2 {
		return euler
	}

	// special case for graph with zero edges (has a degenerate Eulerian path)
	if s == -1 {
		s = 0
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
	stack := util.Stack{}
	stack.Push(s)

	// greedily search through edges in iterative DFS style
	euler.path = util.Stack{}
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
		euler.path.Push(v)
	}

	// check if all edges are used
	if euler.path.Size() != g.E + 1 {
		euler.path = nil
	}

	euler.certifySolution(g)
	return euler
}

// Path
// Returns the sequence of vertices on an Eulerian path.
func (euler *EulerianPath) Path() util.Stack {
	return euler.path
}

// HasEulerianPath
// Returns true if the graph has an Eulerian path.
func (euler *EulerianPath) HasEulerianPath() bool {
	return euler.path != nil
}

// returns any non-isolated vertex; -1 if no such vertex
func (euler *EulerianPath) nonIsolatedVertex(g *graph.Graph) int {
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

// Determines whether a graph has an Eulerian path using necessary
// and sufficient conditions (without computing the path itself):
//    - degree(v) is even for every vertex, except for possibly two
//    - the graph is connected (ignoring isolated vertices)
// This method is solely for unit testing.
func (euler *EulerianPath) satisfiesNecessaryAndSufficientConditions(g *graph.Graph) bool {
	if g.E == 0 {
		return true
	}

	// Condition 1: degree(v) is even except for possibly two
	oddDegreeVertices := 0
	for v := 0; v < g.V; v++ {
		if g.Degree(v) % 2 != 0 {
			oddDegreeVertices++
		}
	}
	if oddDegreeVertices > 2 {
		return false
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
func (euler *EulerianPath) certifySolution(g *graph.Graph) bool {

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
