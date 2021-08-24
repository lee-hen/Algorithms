package cycle

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	"github.com/lee-hen/Algorithms/util"
)

type Cycle struct {
	marked map[int]bool
	edgeTo map[int]int

	cycle util.Stack
}

// New
// Determines whether the undirected graph {@code G} has a cycle and,
// if so, finds such a cycle.
func New(g *graph.Graph) *Cycle {
	cycle := &Cycle{}

	// need special case to identify parallel edge as a cycle
	if cycle.hasParallelEdges(g) {
		return cycle
	}

	// don't need special case to identify self-loop as a cycle
	// if (hasSelfLoop(G)) return;

	cycle.marked = make(map[int]bool)
	cycle.edgeTo = make(map[int]int)
	for v := 0; v < g.V; v++ {
		if !cycle.marked[v] {
			cycle.dfs(g, -1, v)
		}
	}

	return cycle
}

// does this graph have a self loop?
// side effect: initialize cycle to be self loop
func (cycle *Cycle) hasSelfLoop(g *graph.Graph) bool {
	for v := 0; v < g.V; v++ {
		for _, w := range g.Adj(v) {
			if v == w {
				cycle.cycle = make(util.Stack, 0)
				cycle.cycle.Push(v)
				cycle.cycle.Push(v)

				return true
			}
		}
	}

	return false
}

// does this graph have two parallel edges?
// side effect: initialize cycle to be two parallel edges
func (cycle *Cycle) hasParallelEdges(g *graph.Graph) bool {
	cycle.marked = make(map[int]bool)

	for v := 0; v < g.V; v++ {
		// check for parallel edges incident to v
		for _, w := range g.Adj(v) {
			if cycle.marked[w] {
				cycle.cycle = make(util.Stack, 0)
				cycle.cycle.Push(v)
				cycle.cycle.Push(w)
				cycle.cycle.Push(v)
				return true
			}
			cycle.marked[w] = true
		}

		// reset so marked[v] = false for all v
		for _, w := range g.Adj(v) {
			cycle.marked[w] = false
		}
	}

	return false
}

// HasCycle
// Returns true if the graph g has a cycle.
func (cycle *Cycle) HasCycle() bool {
	return cycle.cycle != nil
}

// Cycle
// Returns a cycle in the graph g.
func (cycle *Cycle) Cycle() util.Stack {
	return cycle.cycle
}

func (cycle *Cycle) dfs(g *graph.Graph, u, v int) {
	cycle.marked[v] = true

	for _, w := range g.Adj(v) {
		// short circuit if cycle already found
		if cycle.cycle != nil {
			return
		}

		if !cycle.marked[w] {
			cycle.edgeTo[w] = v
			cycle.dfs(g, v, w)
		} else if w != u {  // check for cycle (but disregard reverse of edge leading to v)
			cycle.cycle = make(util.Stack, 0)

			for x := v; x != w; x = cycle.edgeTo[x] {
				cycle.cycle.Push(x)
			}

			cycle.cycle.Push(w)
			cycle.cycle.Push(v)
		}
	}
}
