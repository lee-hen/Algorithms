package edge_weighted_graph

import (
	edge "github.com/lee-hen/Algorithms/4_graphs/21_edge"

	"fmt"
	"log"
	"math"
	"math/rand"
	"strings"
	"time"
)

const NEWLINE = "\n"

type EdgeWeightedGraph struct {
	V, E int
	adj [][]*edge.Edge
}

// NewEdgeWeightedGraph
// Initializes an empty edge-weighted Digraph V vertices and 0 edges.
func NewEdgeWeightedGraph(v int) *EdgeWeightedGraph {
	if v < 0 {
		log.Fatalln("Number of vertices in a Digraph must be non-negative")
	}
	adj := make([][]*edge.Edge, v, v)

	return &EdgeWeightedGraph {
		V: v,
		adj: adj,
	}
}

// NewRandomEdgeWeightedGraph
// Initializes a random edge-weighted graph with v vertices and e edges.
func NewRandomEdgeWeightedGraph(v, e int) *EdgeWeightedGraph {
	g := NewEdgeWeightedGraph(v)
	if e < 0 {
		log.Fatalln("Number of edges must be non-negative")
	}

	rand.Seed(time.Now().UnixNano())
	V := g.V
	for i := 0; i < e; i++ {
		v := rand.Intn(V)
		w := rand.Intn(V)
		weight :=  math.Round(rand.Float64() * 100) / 100.0

		e := edge.NewEdge(v, w, weight)
		g.AddEdge(e)
	}

	return g
}

// InitEdgeWeightedGraph
// Initializes a edge-weighted graph from the specified input stream.
// The format is the number of vertices V,
// followed by the number of edges E,
// followed by E pairs of vertices, with each entry separated by whitespace.
func InitEdgeWeightedGraph() *EdgeWeightedGraph {
	var v, e int
	var err error
	_, err = fmt.Scan(&v)
	if err != nil {
		log.Fatalln(err)
	}
	if v < 0 {
		log.Fatalln("number of vertices in a edge-weighted must be non-negative")
	}

	_, err = fmt.Scan(&e)
	if err != nil {
		log.Fatalln(err)
	}

	if e < 0 {
		log.Fatalln("Number of edges must be non-negative")
	}

	g := NewEdgeWeightedGraph(v)

	for i := 0; i < e; i++ {
		var v, w int
		var weight float64
		_, err = fmt.Scan(&v, &w, &weight)
		if err != nil {
			log.Fatalln(err)
		}
		e := edge.NewEdge(v, w, weight)
		g.AddEdge(e)
	}

	return g
}

// CloneEdgeWeightedGraph
// Initializes a new edge-weighted graph that is a deep copy of g
func CloneEdgeWeightedGraph(g *EdgeWeightedGraph) *EdgeWeightedGraph {
	if g == nil {
		log.Fatalln("argument is null")
	}

	graph := NewEdgeWeightedGraph(g.V)
	graph.E = g.E

	for v := 0; v < g.V; v++ {
		reverse := make([]*edge.Edge, 0)
		for _, e := range g.adj[v] {
			reverse = append(reverse, e)
		}
		for _, e := range reverse {
			graph.adj[v] = append(graph.adj[v], e)
		}
	}

	return graph
}

// AddEdge
// Adds the directed edge v→w to this edge-weighted graph .
func (graph *EdgeWeightedGraph) AddEdge(e *edge.Edge) {
	v := e.Either()
	w := e.Other(v)

	graph.validateVertex(v)
	graph.validateVertex(w)
	graph.adj[v] = append(graph.adj[v], e)
	graph.adj[w] = append(graph.adj[w], e)
	graph.E++
}

// Adj
// Returns the vertices adjacent from vertex v in this edge-weighted graph.
func (graph *EdgeWeightedGraph) Adj(v int) []*edge.Edge {
	graph.validateVertex(v)
	return graph.adj[v]
}

// Degree
// Returns the number of edges incident from vertex v.
func (graph *EdgeWeightedGraph) Degree(v int) int {
	graph.validateVertex(v)
	return len(graph.adj[v])
}


// Edges
// Returns the Reverse of the edge-weighted graph.
func (graph *EdgeWeightedGraph) Edges() []*edge.Edge {
	list := make([]*edge.Edge, 0)
	for v := 0; v < graph.V; v++ {
		selfLoops := 0
		for _, e := range graph.Adj(v) {
			if e.Other(v) > graph.V {
				list = append(list, e)
			} else if e.Other(v) == v {
				if selfLoops % 2 == 0 {
					list = append(list, e)
				}
				selfLoops++
			}

		}
	}

	return list
}

// String
// Returns a string representation of this edge-weighted graph.
func (graph *EdgeWeightedGraph) String() string {
	s := strings.Builder{}
	s.WriteString(fmt.Sprintf("%d", graph.V) + " vertices, " + fmt.Sprintf("%d", graph.E) + " edges " + NEWLINE)

	for v := 0; v < graph.V; v++ {
		s.WriteString(fmt.Sprintf("%d", v) + ": ")
		for _, e := range graph.adj[v] {
			s.WriteString(fmt.Sprintf("%v", e) + " ")
		}
		s.WriteString(NEWLINE)
	}
	return s.String()
}

func (graph *EdgeWeightedGraph) validateVertex(v int) {
	if v < 0 || v >= graph.V {
		panic(fmt.Sprintf("vertex %d is not between 0 and %d", v, graph.V-1))
	}
}
