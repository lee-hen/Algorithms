package edge_eeighted_digraph

import (
	directedEdge "github.com/lee-hen/Algorithms/4_graphs/22_directed_edge"

	"fmt"
	"log"
	"math/rand"
	"strings"
	"time"
)

const NEWLINE = "\n"

type EdgeWeightedDigraph struct {
	V, E int // number of vertices in this digraph // number of edges in this digraph
	adj [][]*directedEdge.Edge // adj[v] = adjacency list for vertex v
	inDegree []int // inDegree[v] = inDegree of vertex v
}

// NewEdgeWeightedDigraph
// Initializes an empty edge-weighted digraph V vertices and 0 edges.
func NewEdgeWeightedDigraph(v int) *EdgeWeightedDigraph {
	if v < 0 {
		log.Fatalln("Number of vertices in a Digraph must be non-negative")
	}
	adj := make([][]*directedEdge.Edge, v, v)
	inDegree := make([]int, v, v)

	return &EdgeWeightedDigraph {
		V: v,
		adj: adj,
		inDegree: inDegree,
	}
}

// NewRandomEdgeWeightedDigraph
// Initializes a random edge-weighted digraph with v vertices and e edges.
func NewRandomEdgeWeightedDigraph(v, e int) *EdgeWeightedDigraph {
	g := NewEdgeWeightedDigraph(v)
	if e < 0 {
		log.Fatalln("Number of edges in a Digraph must be non-negative")
	}

	rand.Seed(time.Now().UnixNano())
	V := g.V
	for i := 0; i < e; i++ {
		v := rand.Intn(V)
		w := rand.Intn(V)
		weight :=  0.01 * float64(rand.Intn(100))
		g.AddEdge(directedEdge.NewEdge(v, w, weight))
	}

	return g
}

// InitEdgeWeightedDigraph
// Initializes a edge-weighted digraph from the specified input stream.
// The format is the number of vertices V,
// followed by the number of edges E,
// followed by E pairs of vertices, with each entry separated by whitespace.
func InitEdgeWeightedDigraph() *EdgeWeightedDigraph {
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

	graph := NewEdgeWeightedDigraph(v)

	for i := 0; i < e; i++ {
		var v, w int
		var weight float64
		_, err = fmt.Scan(&v, &w, &weight)
		if err != nil {
			log.Fatalln(err)
		}
		graph.AddEdge(directedEdge.NewEdge(v, w, weight))
	}

	return graph
}

// CloneEdgeWeightedDigraph
// Initializes a new digraph that is a deep copy of the specified digraph.
func CloneEdgeWeightedDigraph(g *EdgeWeightedDigraph) *EdgeWeightedDigraph {
	if g == nil {
		log.Fatalln("argument is null")
	}

	graph := NewEdgeWeightedDigraph(g.V)
	graph.E = g.E

	for v := 0; v < g.V; v++ {
		graph.inDegree[v] = g.inDegree[v]
	}

	for v := 0; v < g.V; v++ {
		reverse := make([]*directedEdge.Edge, 0)
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
// Adds the directed edge v→w to this edge-weighted digraph.
func (graph *EdgeWeightedDigraph) AddEdge(e *directedEdge.Edge) {
	v := e.From()
	w := e.To()

	graph.validateVertex(v)
	graph.validateVertex(w)
	graph.adj[v] = append(graph.adj[v], e)
	graph.inDegree[w]++
	graph.E++
}


// Adj
// Returns the directed edges incident from vertex v.
func (graph *EdgeWeightedDigraph) Adj(v int) []*directedEdge.Edge {
	graph.validateVertex(v)
	return graph.adj[v]
}

// OutDegree
// Returns the number of directed edges incident from vertex v.
// This is known as the outDegree of vertex v.
func (graph *EdgeWeightedDigraph) OutDegree(v int) int {
	graph.validateVertex(v)
	return len(graph.adj[v])
}


// InDegree
// Returns the number of directed edges incident to vertex v.
// This is known as the InDegree of vertex v.
func (graph *EdgeWeightedDigraph) InDegree(v int) int {
	graph.validateVertex(v)
	return graph.inDegree[v]
}

// Edges
//  Returns all directed edges in this edge-weighted digraph.
func (graph *EdgeWeightedDigraph) Edges() []*directedEdge.Edge {
	list := make([]*directedEdge.Edge, 0)
	for v := 0; v < graph.V; v++ {
		for _, e := range graph.Adj(v) {
			list = append(list, e)
		}
	}
	return list
}

// String
// Returns a string representation of this edge-weighted digraph.
func (graph *EdgeWeightedDigraph) String() string {
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

func (graph *EdgeWeightedDigraph) validateVertex(v int) {
	if v < 0 || v >= graph.V {
		panic(fmt.Sprintf("vertex %d is not between 0 and %d", v, graph.V-1))
	}
}
