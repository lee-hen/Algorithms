package directed_cycle

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	"github.com/lee-hen/Algorithms/util"
	"log"
)

// Definition. A directed acyclic graph (DAG) is a digraph with no directed cycles.

type DirectedCycle struct {
	marked map[int]bool // marked[v] = has vertex v been marked?
	edgeTo map[int]int // edgeTo[v] = previous vertex on path to v
	onStack map[int]bool // onStack[v] = is vertex on the stack?

	cycle util.Stack // directed cycle (or null if no such cycle)
}

// New
// Determines whether digraph G has a cycle and,
// if so, finds such a cycle.
func New(g *graph.Digraph) *DirectedCycle {
	cycle := &DirectedCycle{}

	cycle.marked = make(map[int]bool)
	cycle.onStack = make(map[int]bool)
	cycle.edgeTo = make(map[int]int)
	for v := 0; v < g.V; v++ {
		if !cycle.marked[v] && cycle.cycle == nil {
			cycle.dfs(g, v)
		}
	}

	return cycle
}

// run DFS and find a directed cycle (if one exists)
func (cycle *DirectedCycle) dfs(g *graph.Digraph, v int) {
	cycle.marked[v] = true
	cycle.onStack[v] = true

	for _, w := range g.Adj(v) {
		// short circuit if cycle already found
		if cycle.cycle != nil {
			return
		}

		if !cycle.marked[w] {
			cycle.edgeTo[w] = v
			cycle.dfs(g, w)
		} else if cycle.onStack[w] {    // trace back directed cycle
			cycle.cycle = make(util.Stack, 0)
			for x := v; x != w; x = cycle.edgeTo[x] {
				cycle.cycle.Push(x)
			}

			cycle.cycle.Push(w)
			cycle.cycle.Push(v)
			cycle.check()
		}
	}
}


// HasCycle
// Returns true if the graph g has a cycle.
func (cycle *DirectedCycle) HasCycle() bool {
	return cycle.cycle != nil
}

// Cycle
// Returns a cycle in the graph g.
func (cycle *DirectedCycle) Cycle() util.Stack {
	return cycle.cycle
}

// certify that digraph has a directed cycle if it reports one
func (cycle *DirectedCycle) check() bool {
	if cycle.HasCycle() {
		first, last := -1, -1
		for _, v := range cycle.Cycle() {
			if first == -1 {
				first = v
			}
			last = v
		}

		if first != last {
			log.Fatalf("cycle begins with %d and ends with %d\n", first, last)
		}
	}
	return true
}
