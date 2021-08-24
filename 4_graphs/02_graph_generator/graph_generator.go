package graph_generator

import (
	minPQ "github.com/lee-hen/Algorithms/2_sorting/21_min_pq"
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	"github.com/lee-hen/Algorithms/util"

	"fmt"
	"log"
	"math/rand"
	"time"
)

// A bipartite graph is a graph whose vertices we can divide into two sets such that all edges connect a vertex in one set with a vertex in the other set.

type Edge struct {
	V, W int
}

func newEdge(v, w int) *Edge {
	if v > w {
		v, w = w, v
	}

	return &Edge{v ,w}
}

// Simple
// Returns a random simple graph containing v vertices and e edges.
func Simple(v, e int) *graph.Graph {
	if e > v * (v - 1) / 2 {
		log.Fatalln("Too many edges")
	}
	if e < 0 {
		log.Fatalln("Too few edges")
	}
	g := graph.NewGraph(v)
	set := make(map[string]struct{})

	rand.Seed(time.Now().UnixNano())
	for g.E < e {
		v, w := rand.Intn(v), rand.Intn(v)
		edge := newEdge(v, w)
		if _, ok := set[key(edge)]; !ok && v != w {
			set[key(edge)] = struct{}{}
			g.AddEdge(v, w)
		}
	}
	return g
}

// SimpleProb
// Returns a random simple graph on V vertices, with an
// edge between any two vertices with probability p. This is sometimes
// referred to as the Erdos-Renyi random graph model.
func SimpleProb(v int, p float64) *graph.Graph {
	if p < 0.0 || p > 1.0 {
		log.Fatalln("Probability must be between 0 and 1")
	}
	g := graph.NewGraph(v)
	for v := 0; v < g.V; v++ {
		for w := v+1; w < g.V; w++ {
			if util.Bernoulli(p) {
				g.AddEdge(v, w)
			}
		}
	}

	return g
}

// Complete
// Returns the complete graph on v vertices.
func Complete(v int) *graph.Graph {
	return SimpleProb(v, 1.0)
}

// CompleteBipartite
// Returns a complete bipartite graph on v1 and v2 vertices.
func CompleteBipartite(v1, v2 int) *graph.Graph {
	return Bipartite(v1, v2, v1*v2)
}

// Bipartite
// Returns a random simple bipartite graph on v1 and v2 vertices
// with e edges.
func Bipartite(v1, v2, e int) *graph.Graph {
	if e > v1 * v2 {
		log.Fatalln("Too many edges")
	}
	if e < 0 {
		log.Fatalln("Too few edges")
	}
	g := graph.NewGraph(v1 + v2)

	vertices := make([]int, v1 + v2, v1 + v2)
	for i := 0; i < v1 + v2; i++ {
		vertices[i] = i
	}
	util.ShuffleIntSlice(vertices)

	set := make(map[string]struct{})
	rand.Seed(time.Now().UnixNano())
	for g.E < e {
		i := rand.Intn(v1)
		j := v1 + rand.Intn(v2)
		edge := newEdge(vertices[i], vertices[j])
		if _, ok := set[key(edge)]; !ok {
			set[key(edge)] = struct{}{}
			g.AddEdge(vertices[i], vertices[j])
		}
	}
	return g
}

// BipartiteProb
// Returns a random simple bipartite graph on v1 and v2 vertices,
// containing each possible edge with probability p.
func BipartiteProb(v1, v2 int, p float64) *graph.Graph {
	if p < 0.0 || p > 1.0 {
		log.Fatalln("Probability must be between 0 and 1")
	}
	vertices := make([]int, v1 + v2, v1 + v2)
	for i := 0; i < v1 + v2; i++ {
		vertices[i] = i
	}
	util.ShuffleIntSlice(vertices)

	g := graph.NewGraph(v1 + v2)
	for i := 0; i < v1; i++ {
		for j := 0; j < v2; j++ {
			if util.Bernoulli(p) {
				g.AddEdge(vertices[i], vertices[v1 + j])
			}
		}
	}

	return g
}

// Path
// Returns a path graph on V vertices.
func Path(v int) *graph.Graph {
	g := graph.NewGraph(v)
	vertices := make([]int, v, v)
	for i := 0; i < v; i++ {
		vertices[i] = i
	}
	util.ShuffleIntSlice(vertices)

	for i := 0; i < v-1; i++ {
		g.AddEdge(vertices[i], vertices[i+1])
	}

	return g
}

// BinaryTree
// Returns a complete binary tree graph on v vertices.
func BinaryTree(v int) *graph.Graph {
	g := graph.NewGraph(v)
	vertices := make([]int, v, v)
	for i := 0; i < v; i++ {
		vertices[i] = i
	}
	util.ShuffleIntSlice(vertices)

	for i := 1; i < v; i++ {
		g.AddEdge(vertices[i], vertices[i-1]/2)
	}
	return g
}

// Cycle
// Returns a cycle graph on v vertices.
func Cycle(v int) *graph.Graph {
	g := graph.NewGraph(v)
	vertices := make([]int, v, v)
	for i := 0; i < v; i++ {
		vertices[i] = i
	}
	util.ShuffleIntSlice(vertices)

	for i := 0; i < v-1; i++ {
		g.AddEdge(vertices[i], vertices[i+1])
	}
	g.AddEdge(vertices[v-1], vertices[0])

	return g
}

// EulerianCycle
// Returns an Eulerian cycle graph on v vertices.
func EulerianCycle(v, e int) *graph.Graph {
	if e <= 0 {
		log.Fatalln("An Eulerian cycle must have at least one edge")
	}

	if v <= 0 {
		log.Fatalln("An Eulerian cycle must have at least one vertex")
	}

	g := graph.NewGraph(v)
	vertices := make([]int, e, e)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < e; i++ {
		vertices[i] = rand.Intn(v)
	}
	for i := 0; i < e-1; i++ {
		g.AddEdge(vertices[i], vertices[i+1])
	}
	g.AddEdge(vertices[e-1], vertices[0])

	return g
}

// EulerianPath
// Returns an Eulerian cycle graph on v vertices.
func EulerianPath(v, e int) *graph.Graph {
	if e < 0 {
		log.Fatalln("An Eulerian cycle must have at least one edge")
	}

	if v <= 0 {
		log.Fatalln("An Eulerian cycle must have at least one vertex")
	}

	g := graph.NewGraph(v)
	vertices := make([]int, e+1, e+1)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < e+1; i++ {
		vertices[i] = rand.Intn(v)
	}
	for i := 0; i < e; i++ {
		g.AddEdge(vertices[i], vertices[i+1])
	}

	return g
}

// Wheel
// returns  a wheel graph on V vertices.
func Wheel(v int) *graph.Graph {
	if v <= 1 {
		log.Fatalln("Number of vertices must be at least 2")
	}
	g := graph.NewGraph(v)
	vertices := make([]int, v, v)
	for i := 0; i < v; i++ {
		vertices[i] = i
	}
	util.ShuffleIntSlice(vertices)

	// simple cycle on V-1 vertices
	for i := 1; i < v-1; i++ {
		g.AddEdge(vertices[i], vertices[i+1])
	}
	g.AddEdge(vertices[v-1], vertices[1])

	// connect vertices[0] to every vertex on cycle
	for i := 1; i < v; i++ {
		g.AddEdge(vertices[0], vertices[i])
	}

	return g
}

// Star
// Returns a star graph on v vertices.
func Star(v int) *graph.Graph {
	if v <= 0 {
		log.Fatalln("Number of vertices must be at least 1")
	}
	g := graph.NewGraph(v)
	vertices := make([]int, v, v)
	for i := 0; i < v; i++ {
		vertices[i] = i
	}
	util.ShuffleIntSlice(vertices)

	// connect vertices[0] to every other vertex
	for i := 1; i < v; i++ {
		g.AddEdge(vertices[0], vertices[i])
	}
	return g
}

// Regular
// Returns a uniformly random k-regular graph on v vertices
// (not necessarily simple). The graph is simple with probability only about e^(-k^2/4),
// which is tiny when k = 14.
func Regular(v, k int) *graph.Graph {
	if v*k % 2 != 0 {
		log.Fatalln("Number of vertices * k must be even")
	}
	g := graph.NewGraph(v)

	// create k copies of each vertex
	vertices := make([]int, v*k, v*k)
	for v := 0; v < g.V; v++ {
		for j := 0; j < k; j++ {
			vertices[v + g.V*j] = v
		}
	}

	// pick a random perfect matching
	util.ShuffleIntSlice(vertices)
	for i := 0; i < g.V*k/2; i++ {
		g.AddEdge(vertices[2*i], vertices[2*i+1])
	}
	return g
}

// Tree
// http://www.proofwiki.org/wiki/Labeled_Tree_from_Prüfer_Sequence
// http://citeseerx.ist.psu.edu/viewdoc/download?doi=10.1.1.36.6484&rep=rep1&type=pdf
// Returns a uniformly random tree on V vertices.
// This algorithm uses a Prufer sequence and takes time proportional to V log V.
func Tree(v int) *graph.Graph {
	g := graph.NewGraph(v)

	// special case
	if v == 1 { return g }

	rand.Seed(time.Now().UnixNano())

	// Cayley's theorem: there are V^(V-2) labeled trees on V vertices
	// Prufer sequence: sequence of V-2 values between 0 and V-1
	// Prufer's proof of Cayley's theorem: Prufer sequences are in 1-1
	// with labeled trees on V vertices
	prufer := make([]int, v-2, v-2)
	for i := 0; i < v-2; i++ {
		prufer[i] = rand.Intn(v)
	}

	// degree of vertex v = 1 + number of times it appers in Prufer sequence
	degree := make([]int, v, v)
	for v := 0; v < g.V; v++ {
		degree[v] = 1
	}
	for i := 0; i < g.V-2; i++ {
		degree[prufer[i]]++
	}

	// pq contains all vertices of degree 1
	pq := minPQ.MinPQ{Heap: minPQ.NewHeap(minPQ.MinHeapFunc, v+1)}
	for v := 0; v < g.V; v++ {
		if degree[v] == 1 {
			pq.Insert(v)
		}
	}

	// repeatedly delMin() degree 1 vertex that has the minimum index
	for i := 0; i < g.V-2; i++ {
		v := pq.DelMin()
		g.AddEdge(v, prufer[i])
		degree[v]--
		degree[prufer[i]]--
		if degree[prufer[i]] == 1 {
			pq.Insert(prufer[i])
		}
	}
	g.AddEdge(pq.DelMin(), pq.DelMin())
	return g
}

func key(edge *Edge) string {
	return fmt.Sprintf("%v:%d", edge.V, edge.W)
}

