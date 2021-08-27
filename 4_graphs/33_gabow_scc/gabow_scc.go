package gabow_scc

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	TC "github.com/lee-hen/Algorithms/4_graphs/29_transitive_closure"
	"github.com/lee-hen/Algorithms/util"
)

type GabowSCC struct {
	marked map[int]bool // marked[v] = has v been visited?
	id map[int]int      // id[v] = id of strong component containing v
	preOrder map[int]int // preOrder[v] = preorder of v
	pre int             // preorder number counter
	count int           // number of strongly-connected components

	stack1, stack2 util.Stack
}

func New(g *graph.Digraph) *GabowSCC {
	gabowSCC := &GabowSCC{}
	gabowSCC.marked = make(map[int]bool)
	gabowSCC.stack1 = make(util.Stack, 0)
	gabowSCC.stack2 = make(util.Stack, 0)
	gabowSCC.id = make(map[int]int)
	gabowSCC.preOrder = make(map[int]int)

	for v := 0; v < g.V; v++ {
		gabowSCC.id[v] = -1
	}

	for v := 0; v < g.V; v++ {
		if !gabowSCC.marked[v] {
			gabowSCC.dfs(g, v)
		}
	}

	if gabowSCC.check(g) {
		return gabowSCC
	}

	return nil
}

func (gabowSCC *GabowSCC) dfs(g *graph.Digraph, v int)  {
	gabowSCC.marked[v] = true
	gabowSCC.preOrder[v] = gabowSCC.pre
	gabowSCC.pre++

	gabowSCC.stack1.Push(v)
	gabowSCC.stack2.Push(v)

	for _, w := range g.Adj(v) {
		if !gabowSCC.marked[w] {
			gabowSCC.dfs(g, w)
		} else if gabowSCC.id[w] == -1 {
					for gabowSCC.preOrder[gabowSCC.stack2.Peek()] > gabowSCC.preOrder[w] {
				gabowSCC.stack2.Pop()
			}
		}
	}


	if gabowSCC.stack2.Peek() == v {
		gabowSCC.stack2.Pop()

		w := gabowSCC.stack1.Pop()
		gabowSCC.id[w] = gabowSCC.count

		for w != v {
			w = gabowSCC.stack1.Pop()
			gabowSCC.id[w] = gabowSCC.count
		}

		gabowSCC.count++
	}
}

// Count
// Returns the number of strong components.
func (gabowSCC *GabowSCC) Count() int {
	return gabowSCC.count
}

// StronglyConnected
// Are vertices v and w in the same strong component?
func (gabowSCC *GabowSCC) StronglyConnected(v, w int) bool {
	return gabowSCC.id[v] == gabowSCC.id[w]
}

// ID
// Returns the component id of the strong component containing vertex v.
func (gabowSCC *GabowSCC) ID(v int) int {
	return gabowSCC.id[v]
}

// does the id[] array contain the strongly connected components?
func (gabowSCC *GabowSCC) check(g *graph.Digraph) bool {
	TC.TransitiveClosure(g)
	for v := 0; v < g.V; v++ {
		for w := 0; w < g.V; w++ {
			if gabowSCC.StronglyConnected(v, w) != (TC.Reachable(v, w) && TC.Reachable(w, v)) {
				return false
			}
		}
	}

	return true
}
