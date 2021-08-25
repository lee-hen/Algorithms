package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	path "github.com/lee-hen/Algorithms/4_graphs/18_breadth_first_directed_paths"

	"fmt"
	"log"
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

// 3 to 0 (2):  3-2-0
// 3 to 1 (3):  3-2-0-1
// 3 to 2 (1):  3-2
// 3 to 3 (0):  3
// 3 to 4 (2):  3-5-4
// 3 to 5 (1):  3-5
// 3 to 6:  not connected
// 3 to 7:  not connected
// 3 to 8:  not connected
// 3 to 9:  not connected
// 3 to 10:  not connected
// 3 to 11:  not connected
// 3 to 12:  not connected

func main() {
	g := graph.InitDigraph()

	var s int
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatalln(err)
	}

	bfs := path.BreadthFirstDirectedPaths(g, s)
	for v := 0; v < g.V; v++ {
		if bfs.HasPathTo(v) {
			fmt.Printf("%d to %d (%d):  ", s, v, bfs.DistTo(v))
			for i := len(bfs.PathTo(v))-1; i >= 0; i-- {
				x := bfs.PathTo(v)[i]
				if x == s {
					fmt.Print(x)
				} else {
					fmt.Print("-", x)
				}
			}
			fmt.Println()
		} else {
			fmt.Printf("%d to %d:  not connected\n", s, v)
		}
	}
}

