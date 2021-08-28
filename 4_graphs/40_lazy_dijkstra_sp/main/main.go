package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/24_edge_weighted_digraph"
	dij "github.com/lee-hen/Algorithms/4_graphs/40_lazy_dijkstra_sp"

	"fmt"
	"log"
)

// 8
// 15
// 4 5 0.35
// 5 4 0.35
// 4 7 0.37
// 5 7 0.28
// 7 5 0.28
// 5 1 0.32
// 0 4 0.38
// 0 2 0.26
// 7 3 0.39
// 1 3 0.29
// 2 7 0.34
// 6 2 0.40
// 3 6 0.52
// 6 0 0.58
// 6 4 0.93
// 8 vertices, 15 edges
// 0: 0->4 0.38000 0->2 0.26000
// 1: 1->3 0.29000
// 2: 2->7 0.34000
// 3: 3->6 0.52000
// 4: 4->5 0.35000 4->7 0.37000
// 5: 5->4 0.35000 5->7 0.28000 5->1 0.32000
// 6: 6->2 0.40000 6->0 0.58000 6->4 0.93000
// 7: 7->5 0.28000 7->3 0.39000

// 0
// 0 to 0 (0.00)
// 0 to 1 (1.05)  5->1 0.32000   4->5 0.35000   0->4 0.38000
// 0 to 2 (0.26)  0->2 0.26000
// 0 to 3 (0.99)  7->3 0.39000   2->7 0.34000   0->2 0.26000
// 0 to 4 (0.38)  0->4 0.38000
// 0 to 5 (0.73)  4->5 0.35000   0->4 0.38000
// 0 to 6 (1.51)  3->6 0.52000   7->3 0.39000   2->7 0.34000   0->2 0.26000
// 0 to 7 (0.60)  2->7 0.34000   0->2 0.26000

func main() {
	g := graph.InitEdgeWeightedDigraph()
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

