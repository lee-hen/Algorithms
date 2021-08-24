package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	cc "github.com/lee-hen/Algorithms/4_graphs/07_cc"

	"fmt"
)

// 13
// 13
// 0 5
// 4 3
// 0 1
// 9 12
// 6 4
// 5 4
// 0 2
// 11 12
// 9 10
// 0 6
// 7 8
// 9 11
// 5 3

// 3 components
// 0 1 2 3 4 5 6
// 7 8
// 9 10 11 12

func main() {
	g := graph.InitGraph()

	c := cc.Init(g)
	// number of connected components
	m := c.Count()
	fmt.Println(m, "components")


	// compute list of vertices in each connected component
	components := make([][]int, m, m)
	for i := 0; i < m; i++ {
		components[i] = make([]int, 0)
	}

	for v := 0; v < g.V; v++ {
		components[c.ID(v)] = append(components[c.ID(v)], v)
	}

	// print results
	for i := 0; i < m; i++ {
		for _, v := range components[i] {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
