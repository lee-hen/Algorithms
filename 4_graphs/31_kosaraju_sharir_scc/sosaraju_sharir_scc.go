package kosaraju_sharir_scc

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	depthFirstOrder "github.com/lee-hen/Algorithms/4_graphs/26_depth_first_order"
	TC "github.com/lee-hen/Algorithms/4_graphs/29_transitive_closure"
)

// Definition. Two vertices v and w are strongly connected if they are mutually reachable: that is, if there is a directed path from v to w and a directed path from w to v.
// A digraph is strongly connected if all its vertices are strongly connected to one another.

// Postorder lemma. Let C be a strong component in a digraph G and let v be any vertex not in C. If there is an edge e pointing from any vertex in C to v, then vertex v appears before every vertex in C in the reverse postorder of GR.
// Proof: See EXERCISE 4.2.15. Proposition H. The Kosaraju—Sharir algorithm identifies the strong components of a digraph G.
// Proof: By induction on the number of strong components identified in the DFS of G.
// After the algorithm has identified the first i components, we assume (by our inductive hypothesis) that the vertices in the first i components are marked and the vertices in the remaining components are unmarked.
// Let s be the unmarked vertex that appears first in the reverse postorder of GR.
// Then, the constructor call dfs(G, s) will visit every vertex in the strong component containing s (which we refer to as component i+1) and only those vertices
// because:
// • Vertices in the first i components will not be visited (because they are already marked).
// • Vertices in component i+1 are not yet marked and are reachable from s using only other vertices in component i+1 (so will be visited and marked).
// • Vertices in components after i+1 will not be visited (or marked): Consider (for the sake of contradiction) the first such vertex v that is visited.
// Let e be an edge that goes from a vertex in component i+1 to v. By the postorder lemma, v appears in the reverse postorder before every vertex in component i+1 (including s).
// This contradicts the definition of s.


// Proposition I. The Kosaraju–Sharir algorithm uses preprocessing time and space proportional to V+E to support constant-time strong connectivity queries in a digraph.
// Proof: The algorithm computes the reverse of the digraph and does two depth-first searches. Each of these three steps takes time proportional to V+E. The reverse copy of the digraph uses space proportional to V+E.


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