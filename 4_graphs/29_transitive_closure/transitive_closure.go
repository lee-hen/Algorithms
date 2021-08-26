package main

import (
	"fmt"
	D "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	dfs "github.com/lee-hen/Algorithms/4_graphs/15_directed_dfs"
)

var tc []*dfs.DirectedDFS

// TransitiveClosure
// Computes the transitive closure of the digraph G.
func TransitiveClosure(g *D.Digraph) {
	tc = make([]*dfs.DirectedDFS, g.V, g.V)
	for v := 0; v < g.V; v++ {
		tc[v] = dfs.New(g, v)
	}
}

// Reachable
// Is there a directed path from vertex v to vertex w in the digraph?
func Reachable(v, w int) bool {
	return tc[v].Marked(w)
}

// 13
// 22
// 4  2
// 2  3
// 3  2
// 6  0
// 0  1
// 2  0
// 11 12
// 12  9
// 9 10
// 9 11
// 7  9
// 10 12
// 11  4
// 4  3
// 3  5
// 6  8
// 8  6
// 5  4
// 0  5
// 6  4
// 6  9
// 7  6
// 0  1  2  3  4  5  6  7  8  9 10 11 12
// --------------------------------------------
// 0:   T  T  T  T  T  T
// 1:      T
// 2:   T  T  T  T  T  T
// 3:   T  T  T  T  T  T
// 4:   T  T  T  T  T  T
// 5:   T  T  T  T  T  T
// 6:   T  T  T  T  T  T  T     T  T  T  T  T
// 7:   T  T  T  T  T  T  T  T  T  T  T  T  T
// 8:   T  T  T  T  T  T  T     T  T  T  T  T
// 9:   T  T  T  T  T  T           T  T  T  T
// 10:   T  T  T  T  T  T           T  T  T  T
// 11:   T  T  T  T  T  T           T  T  T  T
// 12:   T  T  T  T  T  T           T  T  T  T

func main() {
	g := D.InitDigraph()
	TransitiveClosure(g)

	// print header
	fmt.Print("     ")
	for v := 0; v < g.V; v++ {
		fmt.Printf("%3d", v)
	}
	fmt.Println()
	fmt.Println("--------------------------------------------")

	// print transitive closure
	for v := 0; v < g.V; v++ {
		fmt.Printf("%3d: ", v)
		for  w := 0; w < g.V; w++ {
			if Reachable(v, w) {
				fmt.Print("  T")
			} else {
				fmt.Print("   ")
			}
		}
		fmt.Println()
	}

}