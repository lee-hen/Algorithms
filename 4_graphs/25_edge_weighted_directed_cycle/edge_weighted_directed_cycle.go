package edge_weighted_directed_cycle

import (
	directedEdge "github.com/lee-hen/Algorithms/4_graphs/22_directed_edge"
	graph "github.com/lee-hen/Algorithms/4_graphs/24_edge_eeighted_digraph"
	"github.com/lee-hen/Algorithms/util"

	"log"
)

type EdgeWeightedDirectedCycle struct {
	marked map[int]bool // marked[v] = has vertex v been marked?
	edgeTo map[int]*directedEdge.Edge // edgeTo[v] = previous vertex on path to v
	onStack map[int]bool // onStack[v] = is vertex on the stack?

	cycle util.DirectedEdgeStack // directed cycle (or null if no such cycle)
}

// New
// Determines whether edge-weighted digraph G has a cycle and,
// if so, finds such a cycle.
func New(g *graph.EdgeWeightedDigraph) *EdgeWeightedDirectedCycle {
	cycle := &EdgeWeightedDirectedCycle{}

	cycle.marked = make(map[int]bool)
	cycle.onStack = make(map[int]bool)
	cycle.edgeTo = make(map[int]*directedEdge.Edge)

	for v := 0; v < g.V; v++ {
		if !cycle.marked[v] {
			cycle.dfs(g, v)
		}
	}

	//cycle.check()
	return cycle
}

// run DFS and find a edge-weighted digraph cycle (if one exists)
func (cycle *EdgeWeightedDirectedCycle) dfs(g *graph.EdgeWeightedDigraph, v int) {
	cycle.marked[v] = true
	cycle.onStack[v] = true

	for _, e := range g.Adj(v) {
		w := e.To()
		// short circuit if cycle already found
		if cycle.cycle != nil {
			return
		}

		if !cycle.marked[w] {
			cycle.edgeTo[w] = e
			cycle.dfs(g, w)
		} else if cycle.onStack[w] {    // trace back directed cycle
			cycle.cycle = make(util.DirectedEdgeStack, 0)

			f := e
			for f.From() != w {
				cycle.cycle.Push(f)
				f = cycle.edgeTo[f.From()]
			}
			cycle.cycle.Push(f)
			return
		}
	}
	cycle.onStack[v] = false
}

// HasCycle
// Returns true if the edge-weighted digraph g has a cycle.
func (cycle *EdgeWeightedDirectedCycle) HasCycle() bool {
	return cycle.cycle != nil
}

// Cycle
// Returns a cycle in the graph g.
func (cycle *EdgeWeightedDirectedCycle) Cycle() util.DirectedEdgeStack {
	return cycle.cycle
}

// certify that digraph has a directed cycle if it reports one
func (cycle *EdgeWeightedDirectedCycle) check() bool {
	if cycle.HasCycle() {
		var first, last *directedEdge.Edge

		for i := len(cycle.Cycle())-1; i >= 0; i-- {
			e := cycle.cycle[i]
			if first == nil {
				first = e
			}
			if last != nil {
				if last.To() != e.From() {
					log.Fatalf("cycle edges %s and %s not incident\n", last, e)
					return false
				}
			}
			last = e
		}

		if first.From() != last.To() {
			log.Fatalf("cycle edges %s and %s not incident\n", last, first)
			return false
		}
	}
	return true
}
