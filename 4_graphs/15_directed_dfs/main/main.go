package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	dfs "github.com/lee-hen/Algorithms/4_graphs/15_directed_dfs"

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
// 1 2 6

// 0 1 2 3 4 5 6 8 9 10 11 12

func main() {
	g := graph.InitDigraph()

	var s1, s2, s3 int
	_, err := fmt.Scan(&s1, &s2, &s3)
	if err != nil {
		log.Fatalln(err)
	}

	sources := []int{s1, s2, s3}
	search := dfs.Multi(g, sources)
	for v := 0; v < g.V; v++ {
		if search.Marked(v) {
			fmt.Print(v, " ")
		}
	}
}
