package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	cycle "github.com/lee-hen/Algorithms/4_graphs/10_cycle"

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

// 3 4 5 3

func main() {
	g := graph.InitGraph()

	finder := cycle.New(g)
	if finder.HasCycle() {
		for _, v := range finder.Cycle() {
			fmt.Print(v, " ")
		}
		fmt.Println()
	} else {
		fmt.Println("Graph is acyclic")
	}
}
