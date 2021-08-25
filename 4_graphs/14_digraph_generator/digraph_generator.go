package digraph_generator

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	"github.com/lee-hen/Algorithms/util"

	"fmt"
	"log"
	"math/rand"
	"time"
)

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
// Returns a random simple digraph containing V vertices and E edges.
func Simple(v, e int) *graph.Digraph {
	if e > v * (v - 1) {
		log.Fatalln("Too many edges")
	}
	if e < 0 {
		log.Fatalln("Too few edges")
	}
	g := graph.NewDigraph(v)
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
// Returns a random simple digraph on V vertices, with an
// edge between any two vertices with probability p. This is sometimes
// referred to as the Erdos-Renyi random digraph model.
// This implementations takes time propotional to V^2 (even if p is small).
func SimpleProb(v int, p float64) *graph.Digraph {
	if p < 0.0 || p > 1.0 {
		log.Fatalln("Probability must be between 0 and 1")
	}
	g := graph.NewDigraph(v)
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
// Returns the complete digraph on v vertices.
// In a complete digraph, every pair of distinct vertices is connected
// by two antiparallel edges. There are V*(V-1) edges.
func Complete(v int) *graph.Digraph {
	g := graph.NewDigraph(v)
	for v := 0; v < g.V; v++ {
		for w := 0; w < v; w++ {
			if v != w {
				g.AddEdge(v, w)
			}
		}
	}

	return g
}

// Dag
// Returns a random simple DAG containing v vertices and e edges.
// Note: it is not uniformly selected at random among all such DAGs.
func Dag(v, e int) *graph.Digraph {
	if e > v * (v - 1)/2 {
		log.Fatalln("Too many edges")
	}
	if e < 0 {
		log.Fatalln("Too few edges")
	}

	g := graph.NewDigraph(v)
	set := make(map[string]struct{})
	vertices := make([]int, v)
	for i := 0; i < v; i++ {
		vertices[i] = i
	}
	util.ShuffleIntSlice(vertices)
	rand.Seed(time.Now().UnixNano())
	for g.E < e {
		v := rand.Intn(g.V)
		w := rand.Intn(g.V)
		edge := newEdge(v, w)
		if _, ok := set[key(edge)]; !ok && v < w {
			set[key(edge)] = struct{}{}
			g.AddEdge(vertices[v], vertices[w])
		}
	}
	return g
}


// Tournament
// Returns a random tournament digraph on V vertices. A tournament digraph
// is a digraph in which, for every pair of vertices, there is one and only one
// directed edge connecting them. A tournament is an oriented complete graph.
func Tournament(v int) *graph.Digraph {
	g := graph.NewDigraph(v)
	rand.Seed(time.Now().UnixNano())
	for v := 0; v < g.V; v++ {
		for w := v+1; w < g.V; w++ {
			if util.Bernoulli(0.5) {
				g.AddEdge(v, w)
			} else {
				g.AddEdge(w, v)
			}
		}
	}

	return g
}

// CompleteRootedInDAG
// Returns a complete rooted-in DAG on V vertices.
// A rooted in-tree is a DAG in which there is a single vertex
// reachable from every other vertex. A complete rooted in-DAG
// has V*(V-1)/2 edges.
func CompleteRootedInDAG(v int) *graph.Digraph {
	g := graph.NewDigraph(v)
	vertices := make([]int, v)
	for i := 0; i < v; i++ {
		vertices[i] = i
	}
	util.ShuffleIntSlice(vertices)
	for i := 0; i < g.V; i++ {
		for j := i+1; j < g.V; j++ {
			g.AddEdge(vertices[i], vertices[j])
		}
	}
	return g
}


// RootedInDAG
// Returns a random rooted-in DAG on V vertices and E edges.
// A rooted in-tree is a DAG in which there is a single vertex
// reachable from every other vertex.
// The DAG returned is not chosen uniformly at random among all such DAGs.
func RootedInDAG(v, e int) *graph.Digraph {
	if e > v * (v - 1)/2 {
		log.Fatalln("Too many edges")
	}
	if e < v-1 {
		log.Fatalln("Too few edges")
	}

	g := graph.NewDigraph(v)
	set := make(map[string]struct{})

	// fix a topological order
	vertices := make([]int, v)
	for i := 0; i < v; i++ {
		vertices[i] = i
	}
	util.ShuffleIntSlice(vertices)
	rand.Seed(time.Now().UnixNano())

	// one edge pointing from each vertex, other than the root = vertices[V-1]
	for v := 0; v < g.V-1; v++ {
		min := v+1
		max := g.V
		w := rand.Intn(max - min) + min
		edge := newEdge(v, w)
		set[key(edge)] = struct{}{}
		g.AddEdge(vertices[v], vertices[w])
	}


	for g.E < e {
		v := rand.Intn(g.V)
		w := rand.Intn(g.V)
		edge := newEdge(v, w)
		if _, ok := set[key(edge)]; !ok && v < w {
			set[key(edge)] = struct{}{}
			g.AddEdge(vertices[v], vertices[w])
		}
	}
	return g
}

// CompleteRootedOutDAG
// Returns a complete rooted-out DAG on V vertices.
// A rooted out-tree is a DAG in which every vertex is reachable
// from a single vertex. A complete rooted in-DAG has V*(V-1)/2 edges.
func CompleteRootedOutDAG(v, e int) *graph.Digraph {
	g := graph.NewDigraph(v)

	vertices := make([]int, v)
	for i := 0; i < v; i++ {
		vertices[i] = i
	}
	util.ShuffleIntSlice(vertices)
	for i := 0; i < g.V; i++ {
		for j := i+1; j < g.V; j++ {
			g.AddEdge(vertices[j], vertices[i])
		}
	}

	return g
}


// RootedOutDAG
// Returns a random rooted-out DAG on V vertices and E edges.
// A rooted out-tree is a DAG in which every vertex is reachable from a
// single vertex.
// The DAG returned is not chosen uniformly at random among all such DAGs.
func RootedOutDAG(v, e int) *graph.Digraph {
	if e > v * (v - 1)/2 {
		log.Fatalln("Too many edges")
	}
	if e < v-1 {
		log.Fatalln("Too few edges")
	}

	g := graph.NewDigraph(v)
	set := make(map[string]struct{})

	// fix a topological order
	vertices := make([]int, v)
	for i := 0; i < v; i++ {
		vertices[i] = i
	}
	util.ShuffleIntSlice(vertices)
	rand.Seed(time.Now().UnixNano())

	// one edge pointing from each vertex, other than the root = vertices[V-1]
	for v := 0; v < g.V-1; v++ {
		min := v+1
		max := g.V
		w := rand.Intn(max - min) + min
		edge := newEdge(w, v)
		set[key(edge)] = struct{}{}
		g.AddEdge(vertices[w], vertices[v])
	}


	for g.E < e {
		v := rand.Intn(g.V)
		w := rand.Intn(g.V)
		edge := newEdge(w, v)
		if _, ok := set[key(edge)]; !ok && v < w {
			set[key(edge)] = struct{}{}
			g.AddEdge(vertices[w], vertices[v])
		}
	}
	return g
}

// RootedInTree
// Returns a random rooted-in tree on V vertices.
// A rooted in-tree is an oriented tree in which there is a single vertex
// reachable from every other vertex.
// The tree returned is not chosen uniformly at random among all such trees.
func RootedInTree(v int) *graph.Digraph {
	return RootedInDAG(v, v-1)
}



// RootedOutTree
// Returns a random rooted-out tree on V vertices. A rooted out-tree
// is an oriented tree in which each vertex is reachable from a single vertex.
// It is also known as a arborescence or branching
// The tree returned is not chosen uniformly at random among all such trees.
func RootedOutTree(v int) *graph.Digraph {
	return RootedOutDAG(v, v-1)
}

// Path
// Returns a path digraph on V vertices.
func Path(v int) *graph.Digraph {
	g := graph.NewDigraph(v)
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
// Returns a complete binary tree digraph on v vertices.
func BinaryTree(v int) *graph.Digraph {
	g := graph.NewDigraph(v)
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
// Returns a cycle digraph on v vertices.
func Cycle(v int) *graph.Digraph {
	g := graph.NewDigraph(v)
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
// Returns an Eulerian cycle digraph on v vertices.
func EulerianCycle(v, e int) *graph.Digraph {
	if e <= 0 {
		log.Fatalln("An Eulerian cycle must have at least one edge")
	}

	if v <= 0 {
		log.Fatalln("An Eulerian cycle must have at least one vertex")
	}

	g := graph.NewDigraph(v)
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
// Returns an Eulerian cycle digraph on v vertices.
func EulerianPath(v, e int) *graph.Digraph {
	if e < 0 {
		log.Fatalln("An Eulerian cycle must have at least one edge")
	}

	if v <= 0 {
		log.Fatalln("An Eulerian cycle must have at least one vertex")
	}

	g := graph.NewDigraph(v)
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


// Strong
// Returns a random simple digraph on V vertices, E
// edges and (at least) c strong components. The vertices are randomly
// assigned integer labels between 0 and c-1 (corresponding to
// strong components). Then, a strong component is creates among the vertices
// with the same label. Next, random edges (either between two vertices with
// the same labels or from a vetex with a smaller label to a vertex with a
// larger label). The number of components will be equal to the number of
// distinct labels that are assigned to vertices.
func Strong(v, e, c int) *graph.Digraph {
	if c >= v || c <= 0 {
		log.Fatalln("Number of components must be between 1 and V")
	}

	if e <= 2*(v-c) {
		log.Fatalln("Number of edges must be at least 2(V-c)")
	}

	if e >  v*(v-1) / 2 {
		log.Fatalln("Too many edges")
	}

	// the digraph
	g := graph.NewDigraph(v)

	// edges added to G (to avoid duplicate edges)
	set := make(map[string]struct{})
	rand.Seed(time.Now().UnixNano())
	label := make([]int, v)
	for v := 0; v < g.V; v++ {
		label[v] = rand.Intn(c)
	}

	// make all vertices with label c a strong component by
	// combining a rooted in-tree and a rooted out-tree
	for i := 0; i < c; i++ {
		// how many vertices in component c
		count := 0
		for v := 0; v < g.V; v++ {
			if label[v] == i {
				count++
			}
		}

		// if (count == 0) System.err.println("less than desired number of strong components");
		vertices := make([]int, count)
		j := 0
		for v := 0; v < g.V; v++ {
			if label[v] == i {
				vertices[j] = v
				j++
			}
		}
		util.ShuffleIntSlice(vertices)

		// rooted-in tree with root = vertices[count-1]
		for v := 0; v < count-1; v++ {
			min := v+1
			max := count
			w := rand.Intn(max - min + 1) + min
			edge := newEdge(w, v)
			set[key(edge)] = struct{}{}
			g.AddEdge(vertices[w], vertices[v])
		}

		// rooted-out tree with root = vertices[count-1]
		for v := 0; v < count-1; v++ {
			min := v+1
			max := count
			w := rand.Intn(max - min + 1) + min
			edge := newEdge(v, w)
			set[key(edge)] = struct{}{}
			g.AddEdge(vertices[v], vertices[w])
		}
	}


	for g.E < e {
		v := rand.Intn(g.V)
		w := rand.Intn(g.V)
		edge := newEdge(v, w)
		if _, ok := set[key(edge)]; !ok && v != w && label[v] <= label[w] {
			set[key(edge)] = struct{}{}
			g.AddEdge(v, w)
		}
	}

	return g
}

func key(edge *Edge) string {
	return fmt.Sprintf("%v:%d", edge.V, edge.W)
}

