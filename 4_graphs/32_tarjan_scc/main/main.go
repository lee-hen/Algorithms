package main

import (
	"fmt"
	digraph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	tarjanScc "github.com/lee-hen/Algorithms/4_graphs/32_tarjan_scc"
)

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

// 5  strong components
// 1
// 0 2 3 4 5
// 9 10 11 12
// 6 8
// 7

func main() {
	g := digraph.InitDigraph()
	scc := tarjanScc.New(g)

	// number of connected components
	m := scc.Count()
	fmt.Println(m, "strong components")

	// compute list of vertices in each strong component
	components := make([][]int, m, m)
	for v := 0; v < g.V; v++ {
		components[scc.ID(v)] = append(components[scc.ID(v)], v)
	}

	// print results
	for i := 0; i < m; i++ {
		for _, v := range components[i] {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
