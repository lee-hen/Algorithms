package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	cycle "github.com/lee-hen/Algorithms/4_graphs/19_directed_cycle"

	"fmt"
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
// 3

// 3 2 3

func main() {
	g := graph.InitDigraph()

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
