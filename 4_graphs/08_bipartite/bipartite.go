package bipartite

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"

	"fmt"
	"github.com/lee-hen/Algorithms/util"
	"log"
)


// Given a graph, find either (i) a bipartition or (ii) an odd-length cycle.
// Runs in O(E + V) time.

type Bipartite struct {
	isBipartite bool
	color, marked map[int]bool
	edgeTo []int
	cycle util.Stack
}

func New(g *graph.Graph) *Bipartite {
	b := &Bipartite{}
	b.isBipartite = true
	b.color = make(map[int]bool)
	b.marked = make(map[int]bool)
	b.edgeTo = make([]int, g.V, g.V)

	for v := 0; v < g.V; v++ {
		if !b.marked[v] {
			b.dfs(g, v)
		}
	}

	b.check(g)
	return b
}

func (b *Bipartite) dfs(g *graph.Graph, v int) {
	b.marked[v] = true
	for i := len(g.Adj(v))-1; i >= 0; i-- {
		w := g.Adj(v)[i]

		// short circuit if odd-length cycle found
		if len(b.cycle) > 0 {
			return
		}

		// found uncolored vertex, so recur
		if !b.marked[w] {
			b.edgeTo[w] = v
			b.color[w] = !b.color[v]
			b.dfs(g, w)
		} else if b.color[w] == b.color[v] {
			b.isBipartite = false
			b.cycle = make(util.Stack, 0)
			b.cycle.Push(w) // don't need this unless you want to include start vertex twice
			for x := v; x != w; x = b.edgeTo[x] {
				b.cycle.Push(x)
			}
			b.cycle.Push(w)
		}
	}
}

// IsBipartite
// Returns true if the graph is bipartite.
func (b *Bipartite) IsBipartite() bool{
	return b.isBipartite
}

// Color
// Returns the side of the bipartite that vertex v is on.
func (b *Bipartite) Color(v int) bool{
	b.validateVertex(v)

	if !b.isBipartite {
		log.Fatalln("graph is not bipartite")
	}
	return b.color[v]
}

// OddCycle
// Returns an odd-length cycle if the graph is not bipartite, and
// nil otherwise.
func (b *Bipartite) OddCycle() util.Stack {
	return b.cycle
}

func (b *Bipartite) check(g *graph.Graph) bool {
	// graph is bipartite
	if b.isBipartite {
		for v := 0; v < g.V; v++ {
			for _, w := range g.Adj(v) {
				if b.color[v] == b.color[w] {
					log.Fatalf("edge %d-%d with %d and %d in same side of bipartition\n", v, w, v, w)
					return false
				}
			}
		}
	} else {  // graph has an odd-length cycle
		// verify cycle
		first, last := -1, -1
		for _, v := range b.OddCycle() {
			if first == -1 {
				first = v
			}

			last = v
		}

		if first != last {
			log.Fatalf("cycle begins with %d and ends with %d\n", first, last)
			return false
		}
	}

	return true
}

func (b *Bipartite) validateVertex(v int) {
	if v < 0 || v >= len(b.marked) {
		panic(fmt.Sprintf("vertex %d is not between 0 and %d", v, len(b.marked)-1))
	}
}







