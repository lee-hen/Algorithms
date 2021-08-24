package graph

import (
	"fmt"
	"log"
	"strings"
)

// Definition. A path in a graph is a sequence of vertices connected by edges. A simple path is one with no repeated vertices.
// A cycle is a path with at least one edge whose first and last vertices are the same. A simple cycle is a cycle with no repeated edges or vertices (except the requisite repetition of the first and last vertices).
// The length of a path or a cycle is its number of edges.

// Definition. A graph is connected if there is a path from every vertex to every other vertex in the graph.
// A graph that is not connected consists of a set of connected components, which are maximal connected subgraphs.

// Definition. A tree is an acyclic connected graph. A disjoint set of trees is called a forest.
// A spanning tree of a connected graph is a subgraph that contains all of that graph’s vertices and is a single tree.
// A spanning forest of a graph is the union of spanning trees of its connected components.

const NEWLINE = "\n"

type Graph struct {
	V, E int
	adj [][]int
}

// NewGraph
// Initializes an empty graph with V vertices and 0 edges.
// param V the number of vertices
func NewGraph(v int) *Graph {
	if v < 0 {
		log.Fatalln("Number of vertices must be non-negative")
	}
	adj := make([][]int, v, v)
	return &Graph {
		V: v,
		adj: adj,
	}
}

// InitGraph
// Initializes a graph from the specified input stream.
// The format is the number of vertices V,
// followed by the number of edges E,
// followed by E pairs of vertices, with each entry separated by whitespace.
func InitGraph() *Graph {
	var v, e int
	var err error
	_, err = fmt.Scan(&v)
	if err != nil {
		log.Fatalln(err)
	}
	if v < 0 {
		log.Fatalln("number of vertices in a Graph must be non-negative")
	}

	_, err = fmt.Scan(&e)
	if err != nil {
		log.Fatalln(err)
	}

	if e < 0 {
		log.Fatalln("number of edges in a Graph must be non-negative")
	}

	graph := NewGraph(v)

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

// CloneGraph
// Initializes a new graph that is a deep copy of G
func CloneGraph(g *Graph) *Graph {
	graph := NewGraph(g.V)
	graph.E = g.E

	if graph.V < 0 {
		log.Fatalln("Number of vertices must be non-negative")
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
// Adds the undirected edge v-w to this graph.
func (graph *Graph) AddEdge(v, w int) {
	graph.validateVertex(v)
	graph.validateVertex(w)
	graph.E++
	graph.adj[v] = append(graph.adj[v], w)
	graph.adj[w] = append(graph.adj[w], v)
}

// Adj
// Returns the vertices adjacent to vertex v
func (graph *Graph) Adj(v int) []int {
	graph.validateVertex(v)
	return graph.adj[v]
}

// Degree
// Returns the degree of vertex v
func (graph *Graph) Degree(v int) int {
	graph.validateVertex(v)
	return len(graph.adj[v])
}

// String
// Returns a string representation of this graph.
func (graph *Graph) String() string {
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

func (graph *Graph) validateVertex(v int) {
	if v < 0 || v >= graph.V {
		panic(fmt.Sprintf("vertex %d is not between 0 and %d", v, graph.V-1))
	}
}
