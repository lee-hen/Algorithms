package digraph

import (
	"fmt"
	"log"
	"strings"
)

// Definition. A directed graph (or digraph) is a set of vertices and a collection of directed edges.
// Each directed edge connects an ordered pair of vertices.

// Definition. A directed path in a digraph is a sequence of vertices in which there is a (directed) edge pointing from each vertex in the sequence to its successor in the sequence.
// A directed cycle is a directed path with at least one edge whose first and last vertices are the same. A simple cycle is a cycle with no repeated edges or vertices (except the requisite repetition of the first and last vertices).
// The length of a path or a cycle is its number of edges.

// Proposition D. DFS marks all the vertices in a digraph reachable from a given set of sources in time proportional to the sum of the outdegrees of the vertices marked.
// Proof: Same as PROPOSITION A on page 531.

const NEWLINE = "\n"

type Digraph struct {
	V, E int // number of vertices in this digraph // number of edges in this digraph
	adj [][]int // adj[v] = adjacency list for vertex v
	inDegree []int // inDegree[v] = inDegree of vertex v
}

// NewDigraph
// Initializes an empty graph Digraph V vertices and 0 edges.
func NewDigraph(v int) *Digraph {
	if v < 0 {
		log.Fatalln("Number of vertices in a Digraph must be non-negative")
	}
	adj := make([][]int, v, v)
	inDegree := make([]int, v, v)

	return &Digraph {
		V: v,
		adj: adj,
		inDegree: inDegree,
	}
}

// InitDigraph
// Initializes a digraph from the specified input stream.
// The format is the number of vertices V,
// followed by the number of edges E,
// followed by E pairs of vertices, with each entry separated by whitespace.
func InitDigraph() *Digraph {
	var v, e int
	var err error
	_, err = fmt.Scan(&v)
	if err != nil {
		log.Fatalln(err)
	}
	if v < 0 {
		log.Fatalln("number of vertices in a Digraph must be non-negative")
	}

	_, err = fmt.Scan(&e)
	if err != nil {
		log.Fatalln(err)
	}

	if e < 0 {
		log.Fatalln("number of edges in a Digraph must be non-negative")
	}

	graph := NewDigraph(v)

	for i := 0; i < e; i++ {
		var v, w int
		_, err = fmt.Scan(&v, &w)
		if err != nil {
			log.Fatalln(err)
		}
		graph.AddEdge(v, w)
	}

	return graph
}

// CloneDigraph
// Initializes a new digraph that is a deep copy of the specified digraph.
func CloneDigraph(g *Digraph) *Digraph {
	if g == nil {
		log.Fatalln("argument is null")
	}

	graph := NewDigraph(g.V)
	graph.E = g.E

	if graph.V < 0 {
		log.Fatalln("Number of vertices must be non-negative")
	}

	// update indegrees
	graph.inDegree = make([]int, g.V, g.V)
	for v := 0; v < g.V; v++ {
		graph.inDegree[v] = g.inDegree[v]
	}

	// update adjacency lists
	graph.adj = make([][]int, graph.V, graph.V)

	for v := 0; v < g.V; v++ {
		reverse := make([]int, 0)
		for _, w := range g.adj[v] {
			reverse = append(reverse, w)
		}
		for _, w := range reverse {
			graph.adj[v] = append(graph.adj[v], w)
		}
	}

	return graph
}

// AddEdge
// Adds the directed edge v→w to this digraph.
func (graph *Digraph) AddEdge(v, w int) {
	graph.validateVertex(v)
	graph.validateVertex(w)
	graph.adj[v] = append(graph.adj[v], w)
	graph.E++
	graph.inDegree[w]++

}

// Adj
// Returns the vertices adjacent from vertex v in this digraph.
func (graph *Digraph) Adj(v int) []int {
	graph.validateVertex(v)
	return graph.adj[v]
}

// OutDegree
// Returns the number of directed edges incident from vertex v.
// This is known as the outDegree of vertex v.
func (graph *Digraph) OutDegree(v int) int {
	graph.validateVertex(v)
	return len(graph.adj[v])
}


// InDegree
// Returns the number of directed edges incident to vertex v.
// This is known as the InDegree of vertex v.
func (graph *Digraph) InDegree(v int) int {
	graph.validateVertex(v)
	return graph.inDegree[v]
}

// Reverse
// Returns the Reverse of the Digraph.
func (graph *Digraph) Reverse() *Digraph {
	reverse := NewDigraph(graph.V)
	for v := 0; v < graph.V; v++ {
		for _, w := range graph.adj[v] {
			reverse.AddEdge(w, v)
		}
	}
	return reverse
}

// String
// Returns a string representation of this graph.
func (graph *Digraph) String() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("%d", graph.V) + " vertices, " + fmt.Sprintf("%d", graph.E) + " edges " + NEWLINE)

	for v := 0; v < graph.V; v++ {
		s.WriteString(fmt.Sprintf("%d", v) + ": ")
		for _, w := range graph.adj[v] {
			s.WriteString(fmt.Sprintf("%d", w) + " ")
		}
		s.WriteString(NEWLINE)
	}
	return s.String()
}

func (graph *Digraph) validateVertex(v int) {
	if v < 0 || v >= graph.V {
		panic(fmt.Sprintf("vertex %d is not between 0 and %d", v, graph.V-1))
	}
}
