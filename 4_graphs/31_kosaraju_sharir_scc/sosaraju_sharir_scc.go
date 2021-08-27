package kosaraju_sharir_scc

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	depthFirstOrder "github.com/lee-hen/Algorithms/4_graphs/26_depth_first_order"
	TC "github.com/lee-hen/Algorithms/4_graphs/29_transitive_closure"
)

// Definition. Two vertices v and w are strongly connected if they are mutually reachable: that is, if there is a directed path from v to w and a directed path from w to v.
// A digraph is strongly connected if all its vertices are strongly connected to one another.

type KosarajuSharirSCC struct {
	marked map[int]bool
	id map[int]int
	count int
}

// New
// Computes the strong components of the digraph g.
func New(g *graph.Digraph) *KosarajuSharirSCC {
	kosarajuSharirSCC := &KosarajuSharirSCC{}

	// compute reverse postorder of reverse graph
	dfs := depthFirstOrder.NewDigraph(g.Reverse())

	kosarajuSharirSCC.marked = map[int]bool{}
	kosarajuSharirSCC.id = map[int]int{}

	// run DFS on G, using reverse postorder to guide calculation
	for _, v := range dfs.ReversePost() {
		if !kosarajuSharirSCC.marked[v] {
			kosarajuSharirSCC.dfs(g, v)
			kosarajuSharirSCC.count++
		}
	}

	if kosarajuSharirSCC.check(g) {
		return kosarajuSharirSCC
	}

	return nil
}

// DFS on graph G
func (kosarajuSharirSCC *KosarajuSharirSCC) dfs(g *graph.Digraph, v int)  {
	kosarajuSharirSCC.marked[v] = true
	kosarajuSharirSCC.id[v] = kosarajuSharirSCC.count

	for _, w := range g.Adj(v) {
		if !kosarajuSharirSCC.marked[w] {
			kosarajuSharirSCC.dfs(g, w)
		}
	}
}

// Count
// Returns the number of strong components.
func (kosarajuSharirSCC *KosarajuSharirSCC) Count() int {
	return kosarajuSharirSCC.count
}

// StronglyConnected
// Are vertices v and w in the same strong component?
func (kosarajuSharirSCC *KosarajuSharirSCC) StronglyConnected(v, w int) bool {
	return kosarajuSharirSCC.id[v] == kosarajuSharirSCC.id[w]
}

// ID
// Returns the component id of the strong component containing vertex v.
func (kosarajuSharirSCC *KosarajuSharirSCC) ID(v int) int {
	return kosarajuSharirSCC.id[v]
}

// does the id[] array contain the strongly connected components?
func (kosarajuSharirSCC *KosarajuSharirSCC) check(g *graph.Digraph) bool {
	TC.TransitiveClosure(g)
	for v := 0; v < g.V; v++ {
		for w := 0; w < g.V; w++ {
			if kosarajuSharirSCC.StronglyConnected(v, w) != (TC.Reachable(v, w) && TC.Reachable(w, v)) {
				return false
			}
		}
	}

	return true
}