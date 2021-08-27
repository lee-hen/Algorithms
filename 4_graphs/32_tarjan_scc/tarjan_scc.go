package tarjan_scc

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	TC "github.com/lee-hen/Algorithms/4_graphs/29_transitive_closure"
	"github.com/lee-hen/Algorithms/util"
)

type TarjanSCC struct {
	marked map[int]bool // marked[v] = has v been visited?
	id map[int]int      // id[v] = id of strong component containing v
	low map[int]int     // low[v] = low number of v
	pre int             // preorder number counter
	count int           // number of strongly-connected components

	stack util.Stack
}

func New(g *graph.Digraph) *TarjanSCC {
	tarjanSCC := &TarjanSCC{}
	tarjanSCC.marked = make(map[int]bool)
	tarjanSCC.stack = make(util.Stack, 0)
	tarjanSCC.id = make(map[int]int)
	tarjanSCC.low = make(map[int]int)

	for v := 0; v < g.V; v++ {
		if !tarjanSCC.marked[v] {
			tarjanSCC.dfs(g, v)
		}
	}

	if tarjanSCC.check(g) {
		return tarjanSCC
	}

	return nil
}

func (tarjanSCC *TarjanSCC) dfs(g *graph.Digraph, v int)  {
	tarjanSCC.marked[v] = true
	tarjanSCC.low[v] = tarjanSCC.pre
	tarjanSCC.pre++
	min := tarjanSCC.low[v]
	tarjanSCC.stack.Push(v)

	for _, w := range g.Adj(v) {
		if !tarjanSCC.marked[w] {
			tarjanSCC.dfs(g, w)
		}

		if tarjanSCC.low[w] < min {
			min = tarjanSCC.low[w]
		}
	}

	if min < tarjanSCC.low[v] {
		tarjanSCC.low[v] = min
		return
	}

	w := tarjanSCC.stack.Pop()
	tarjanSCC.id[w] = tarjanSCC.count
	tarjanSCC.low[w] = g.V

	for w != v {
		w = tarjanSCC.stack.Pop()
		tarjanSCC.id[w] = tarjanSCC.count
		tarjanSCC.low[w] = g.V
	}

	tarjanSCC.count++
}

// Count
// Returns the number of strong components.
func (tarjanSCC *TarjanSCC) Count() int {
	return tarjanSCC.count
}

// StronglyConnected
// Are vertices v and w in the same strong component?
func (tarjanSCC *TarjanSCC) StronglyConnected(v, w int) bool {
	return tarjanSCC.id[v] == tarjanSCC.id[w]
}

// ID
// Returns the component id of the strong component containing vertex v.
func (tarjanSCC *TarjanSCC) ID(v int) int {
	return tarjanSCC.id[v]
}

// does the id[] array contain the strongly connected components?
func (tarjanSCC *TarjanSCC) check(g *graph.Digraph) bool {
	TC.TransitiveClosure(g)
	for v := 0; v < g.V; v++ {
		for w := 0; w < g.V; w++ {
			if tarjanSCC.StronglyConnected(v, w) != (TC.Reachable(v, w) && TC.Reachable(w, v)) {
				return false
			}
		}
	}

	return true
}
