package topological_x

import (
	D "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	E "github.com/lee-hen/Algorithms/4_graphs/24_edge_eeighted_digraph"
	"log"
)

type TopologicalX struct {
	order []int // topological order
	ranks map[int]int // rank[v] = rank of vertex v in order
}

// NewDigraph
// Determines whether the digraph g has a topological order and, if so,
// finds such a topological order.
func NewDigraph(g *D.Digraph) *TopologicalX {
	topological := &TopologicalX{}

	// indegrees of remaining vertices
	inDegree := make([]int, g.V, g.V)
	for v := 0; v < g.V; v++ {
		inDegree[v] = g.InDegree(v)
	}

	// initialize
	topological.ranks = make(map[int]int)
	topological.order = make([]int, 0)
	count := 0

	// initialize queue to contain all vertices with indegree = 0
	queue := make([]int, 0)
	for v := 0; v < g.V; v++ {
		if inDegree[v] == 0 {
			queue = append(queue, v)
		}
	}

	for len(queue) > 0 {
		var v int
		v, queue = queue[0], queue[1:]
		topological.order = append(topological.order, v)
		topological.ranks[v] = count
		count++

		for i := len(g.Adj(v))-1; i >= 0; i-- {
			w := g.Adj(v)[i]
			inDegree[w]--
			if inDegree[w] == 0 {
				queue = append(queue, w)
			}
		}
	}

	// there is a directed cycle in subgraph of vertices with indegree >= 1.
	if count != g.V {
		topological.order = nil
	}

	topological.check(g)
	return topological
}

// NewEdgeWeightedDigraph
// Determines whether the digraph g has a topological order and, if so,
// finds such a topological order.
func NewEdgeWeightedDigraph(g *E.EdgeWeightedDigraph) *TopologicalX {
	topological := &TopologicalX{}

	// indegrees of remaining vertices
	inDegree := make([]int, g.V, g.V)
	for v := 0; v < g.V; v++ {
		inDegree[v] = g.InDegree(v)
	}

	// initialize
	topological.ranks = make(map[int]int)
	topological.order = make([]int, 0)
	count := 0

	// initialize queue to contain all vertices with indegree = 0
	queue := make([]int, 0)
	for v := 0; v < g.V; v++ {
		if inDegree[v] == 0 {
			queue = append(queue, v)
		}
	}

	for len(queue) > 0 {
		var v int
		v, queue = queue[0], queue[1:]
		topological.order = append(topological.order, v)
		topological.ranks[v] = count
		count++

		for i := len(g.Adj(v))-1; i >= 0; i-- {
			e := g.Adj(v)[i]
			w := e.To()
			inDegree[w]--
			if inDegree[w] == 0 {
				queue = append(queue, w)
			}
		}
	}

	// there is a directed cycle in subgraph of vertices with indegree >= 1.
	if count != g.V {
		topological.order = nil
	}

	return topological
}

// Order
// Returns a topological order if the digraph has a topologial order,
// nil otherwise.
func (topological *TopologicalX) Order() []int {
	return topological.order
}

// HasOrder
// Does the digraph have a topological order?
// return true if the digraph has a topological order (or equivalently,
// if the digraph is a DAG), and false otherwise
func (topological *TopologicalX) HasOrder() bool {
	return topological.order != nil
}

// Rank
// The the rank of vertex v in the topological order;
// -1 if the digraph is not a DAG
func (topological *TopologicalX) Rank(v int) int {
	if topological.HasOrder() {
		return topological.ranks[v]
	} else {
		return -1
	}
}

// certify that digraph has a directed cycle if it reports one
func (topological *TopologicalX) check(g *D.Digraph) bool {
	// digraph is acyclic
	if topological.HasOrder() {
		found := make(map[int]bool)
		for i := 0; i < g.V; i++ {
			found[topological.Rank(i)] = true
		}

		for i := 0; i < g.V; i++ {
			if !found[i] {
				log.Fatalf("No vertex with rank  %d\n", i)
				return false
			}
		}

		// check that ranks provide a valid topological order
		for v := 0; v < g.V; v++ {
			for _, w := range g.Adj(v) {
				if topological.Rank(v) > topological.Rank(w) {
					log.Fatalf("%d-%d: rank(%d) = %d, rank(%d) = %d\n",
						v, w, v, topological.Rank(v), w, topological.Rank(w))

					return false
				}
			}
		}
		// check that order() is consistent with rank()
		r := 0
		for _, v := range topological.Order() {
			if topological.Rank(v) != r {
				log.Fatalln("order() and rank() inconsistent")
				return false
			}
			r++
		}
	}
	return true
}
