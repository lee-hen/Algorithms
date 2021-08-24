package bipartite_x

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	"github.com/lee-hen/Algorithms/util"

	"fmt"
	"log"
)

const (
	WHITE = false
	BLACK = true
)

type BipartiteX struct {
	isBipartite bool
	color, marked map[int]bool
	edgeTo []int
	cycle []int
}

// New
// Determines whether an undirected graph is bipartite and finds either a
// bipartition or an odd-length cycle.
func New(g *graph.Graph) *BipartiteX {
	b := &BipartiteX{}
	b.isBipartite = true
	b.color = make(map[int]bool)
	b.marked = make(map[int]bool)
	for v := 0; v < g.V; v++ {
		b.marked[v] = false
		b.color[v] = false
	}
	b.edgeTo = make([]int, g.V, g.V)

	for v := 0; v < g.V && b.isBipartite; v++ {
		if !b.marked[v] {
			b.bfs(g, v)
		}
	}

	b.check(g)
	return b
}

func (b *BipartiteX) bfs(g *graph.Graph, s int) {
	queue := make([]int, 0)
	b.color[s] = WHITE
	b.marked[s] = true
	queue = append(queue, s)

	for len(queue) > 0 {
		var v int
		v, queue = queue[0], queue[1:]

		for _, w := range g.Adj(v) {
			if !b.marked[w] {
				b.marked[w] = true
				b.edgeTo[w] = v
				b.color[w] = !b.color[v]

				queue = append(queue, w)
			} else if b.color[w] == b.color[v] {
				b.isBipartite = false

				// to form odd cycle, consider s-v path and s-w path
				// and let x be closest node to v and w common to two paths
				// then (w-x path) + (x-v path) + (edge v-w) is an odd-length cycle
				// Note: distTo[v] == distTo[w];
				b.cycle = make([]int, 0)
				stack := make(util.Stack, 0)
				x, y := v, w
				for x != y {
					stack.Push(x)
					b.cycle = append(b.cycle, y)
					x = b.edgeTo[x]
					y = b.edgeTo[y]
				}
				stack.Push(x)

				for !stack.IsEmpty() {
					b.cycle = append(b.cycle, stack.Pop())
				}
				return
			}
		}
	}
}


// IsBipartite
// Returns true if the graph is bipartite.
func (b *BipartiteX) IsBipartite() bool{
	return b.isBipartite
}


// Color
// Returns the side of the bipartite that vertex v is on.
func (b *BipartiteX) Color(v int) bool{
	b.validateVertex(v)

	if !b.isBipartite {
		log.Fatalln("graph is not bipartite")
	}
	return b.color[v]
}

// OddCycle
// Returns an odd-length cycle if the graph is not bipartite, and
// nil otherwise.
func (b *BipartiteX) OddCycle() util.Stack {
	return b.cycle
}

func (b *BipartiteX) check(g *graph.Graph) bool {
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

func (b *BipartiteX) validateVertex(v int) {
	if v < 0 || v >= len(b.marked) {
		panic(fmt.Sprintf("vertex %d is not between 0 and %d", v, len(b.marked)-1))
	}
}
