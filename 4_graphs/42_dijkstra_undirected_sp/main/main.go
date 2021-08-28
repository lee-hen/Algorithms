package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/23_edge_weighted_graph"
	dij "github.com/lee-hen/Algorithms/4_graphs/42_dijkstra_undirected_sp"

	"fmt"
	"log"
)

// 8
// 16
// 4 5 0.35
// 4 7 0.37
// 5 7 0.28
// 0 7 0.16
// 1 5 0.32
// 0 4 0.38
// 2 3 0.17
// 1 7 0.19
// 0 2 0.26
// 1 2 0.36
// 1 3 0.29
// 2 7 0.34
// 6 2 0.40
// 3 6 0.52
// 6 0 0.58
// 6 4 0.93
// 8 vertices, 16 edges
// 0: 0-7 0.16000 0-4 0.38000 0-2 0.26000 6-0 0.58000
// 1: 1-5 0.32000 1-7 0.19000 1-2 0.36000 1-3 0.29000
// 2: 2-3 0.17000 0-2 0.26000 1-2 0.36000 2-7 0.34000 6-2 0.40000
// 3: 2-3 0.17000 1-3 0.29000 3-6 0.52000
// 4: 4-5 0.35000 4-7 0.37000 0-4 0.38000 6-4 0.93000
// 5: 4-5 0.35000 5-7 0.28000 1-5 0.32000
// 6: 6-2 0.40000 3-6 0.52000 6-0 0.58000 6-4 0.93000
// 7: 4-7 0.37000 5-7 0.28000 0-7 0.16000 1-7 0.19000 2-7 0.34000

// 6
// 6 to 0 (0.58)  6-0 0.58000
// 6 to 1 (0.76)  1-2 0.36000   6-2 0.40000
// 6 to 2 (0.40)  6-2 0.40000
// 6 to 3 (0.52)  3-6 0.52000
// 6 to 4 (0.93)  6-4 0.93000
// 6 to 5 (1.02)  5-7 0.28000   2-7 0.34000   6-2 0.40000
// 6 to 6 (0.00)
// 6 to 7 (0.74)  2-7 0.34000   6-2 0.40000

func main() {
	g := graph.InitEdgeWeightedGraph()
	fmt.Println(g)

	var s int
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatalln(err)
	}

	sp := dij.New(g, s)

	// print shortest path
	for t := 0; t < g.V; t++ {
		if sp.HasPathTo(t) {
			fmt.Printf("%d to %d (%.2f)  ", s, t, sp.DistTo(t))
			for _, e := range sp.PathTo(t) {
				fmt.Print(e, "   ")
			}
			fmt.Println()
		} else {
			fmt.Printf("%d to %d         no path\n", s, t)
		}
	}
}

