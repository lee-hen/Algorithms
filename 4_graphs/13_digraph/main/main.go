package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"

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

// 13 vertices, 22 edges
// 0: 1 5
// 1:
// 2: 3 0
// 3: 2 5
// 4: 2 3
// 5: 4
// 6: 0 8 4 9
// 7: 9 6
// 8: 6
// 9: 10 11
// 10: 12
// 11: 12 4
// 12: 9

func main() {
	g := graph.InitDigraph()
	fmt.Println(g)
}
