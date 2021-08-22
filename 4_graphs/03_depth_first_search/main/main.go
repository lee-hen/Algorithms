package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	dfs "github.com/lee-hen/Algorithms/4_graphs/03_depth_first_search"

	"fmt"
	"log"
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
// 0

// 0 1 2 3 4 5 6
// NOT connected

func main() {
	g := graph.InitGraph()

	var s int
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatalln(err)
	}

	search := dfs.New(g, s)
	for v := 0; v < g.V; v++ {
		if search.Marked(v) {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()

	if search.Count() != g.V {
		fmt.Println("NOT connected")
	} else {
		fmt.Println("connected")
	}
}


