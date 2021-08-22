package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"

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

// 13 vertices, 13 edges
// 0: 5 1 2 6
// 1: 0
// 2: 0
// 3: 4 5
// 4: 3 6 5
// 5: 0 4 3
// 6: 4 0
// 7: 8
// 8: 7
// 9: 12 10 11
// 10: 9
// 11: 12 9
// 12: 9 11

func main() {
	g := graph.InitGraph()
	fmt.Println(g)
}
