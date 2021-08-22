package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	path "github.com/lee-hen/Algorithms/4_graphs/05_depth_first_paths"

	"fmt"
	"log"
)

// 6
// 8
// 0 5
// 2 4
// 2 3
// 1 2
// 0 1
// 3 4
// 3 5
// 0 2
// 0

// 0 to 0:  0
// 0 to 1:  0-2-1
// 0 to 2:  0-2
// 0 to 3:  0-2-3
// 0 to 4:  0-2-3-4
// 0 to 5:  0-2-3-5

func main() {
	g := graph.InitGraph()

	var s int
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatalln(err)
	}

	dfs := path.DepthFirstPaths(g, s)
	for v := 0; v < g.V; v++ {
		if dfs.HasPathTo(v) {
			fmt.Printf("%d to %d:  ", s, v)
			for i := len(dfs.PathTo(v))-1; i >= 0; i-- {
				x := dfs.PathTo(v)[i]
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

