package cc

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"
)

// Proposition C. DFS uses preprocessing time and space proportional to V+E to support constant-time connectivity queries in a graph.
// Proof: Immediate from the code. Each adjacency-list entry is examined exactly once, and there are 2E such entries (two for each edge);
// initialzing the marked[] and id[] arrays takes time proportional to E. Instance methods examine or return one or two instance variables.

type CC struct {
	marked map[int]bool  // marked[v] = has vertex v been marked?
	// id[v] = id of connected component containing v
	// size[id] = number of vertices in given component
	id, size map[int]int
	// number of connected components
	count int
}

// Init
// Computes the connected components of the undirected graph g.
func Init(g *graph.Graph) *CC {
	cc := &CC{}
	cc.marked = make(map[int]bool)
	cc.id = make(map[int]int)
	cc.size = make(map[int]int)

	for v := 0; v < g.V; v++ {
		if !cc.marked[v] {
			cc.dfs(g, v)
			cc.count++
		}
	}

	return cc
}

// depth-first search for a Graph
func (cc *CC) dfs(g *graph.Graph, v int) {
	cc.marked[v] = true
	cc.id[v] = cc.count
	cc.size[cc.count]++

	for i := len(g.Adj(v))-1; i >= 0; i-- {
		w := g.Adj(v)[i]

		if !cc.marked[w] {
			cc.dfs(g, w)
		}
	}
}

// ID
// Returns the component id of the connected component containing vertex v.
func (cc *CC) ID(v int) int {
	return cc.id[v]
}

// Size
// Returns the number of vertices in the connected component containing vertex v
func (cc *CC) Size(v int) int {
	return cc.size[cc.id[v]]
}

// Count
// Returns the number of connected components in the graph g.
func (cc *CC) Count() int {
	return cc.count
}

// Connected
// Returns true if vertices v and w are in the same
// connected component.
func (cc *CC) Connected(v, w int) bool {
	return cc.id[v] == cc.id[w]
}

