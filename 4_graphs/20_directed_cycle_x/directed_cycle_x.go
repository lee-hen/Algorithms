package directed_cycle_x

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	"github.com/lee-hen/Algorithms/util"
	"log"
)

type DirectedCycleX struct {
	cycle util.Stack // the directed cycle; null if digraph is acyclic
}

// New
// Determines whether digraph G has a cycle and,
// if so, finds such a cycle.
func New(g *graph.Digraph) *DirectedCycleX {
	cycle := &DirectedCycleX{}
	inDegree := make([]int, g.V, g.V)
	for v := 0; v < g.V; v++ {
		inDegree[v] = g.InDegree(v)
	}

	queue := make([]int, 0)
	for v := 0; v < g.V; v++ {
		if inDegree[v] == 0 {
			queue = append(queue, v)
		}
	}

	for len(queue) > 0 {
		var v int
		v, queue = queue[0], queue[1:]
		for _, w := range g.Adj(v){
			inDegree[w]--
			if inDegree[w] == 0 {
				queue = append(queue, w)
			}
		}
	}

	// there is a directed cycle in subgraph of vertices with indegree >= 1.
	edgeTo := make(map[int]int)
	root := -1

	for v := 0; v < g.V; v++ {
		if inDegree[v] == 0 {
			continue
		} else {
			root = v
		}
		for _, w := range g.Adj(v){
			if inDegree[w] > 0 {
				edgeTo[w] = v
			}
		}
	}

	if root != -1 {
		// find any vertex on cycle
		visited := make(map[int]bool)

		for !visited[root] {
			visited[root] = true
			root = edgeTo[root]
		}

		// extract cycle
		cycle.cycle = util.Stack{}
		v := root

		cycle.cycle.Push(v)
		v = edgeTo[v]
		for v != root {
			cycle.cycle.Push(v)
			v = edgeTo[v]
		}
		cycle.cycle.Push(v)
	}
	cycle.check()
	return cycle
}

// HasCycle
// Returns true if the graph g has a cycle.
func (cycle *DirectedCycleX) HasCycle() bool {
	return cycle.cycle != nil
}

// Cycle
// Returns a cycle in the graph g.
func (cycle *DirectedCycleX) Cycle() util.Stack {
	return cycle.cycle
}

// certify that digraph has a directed cycle if it reports one
func (cycle *DirectedCycleX) check() bool {
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
