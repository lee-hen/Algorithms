package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/24_edge_weighted_digraph"
	AcyclicLP "github.com/lee-hen/Algorithms/4_graphs/44_acyclic_lp"

	"fmt"
	"log"
)

// 8
// 13
// 5 4 0.35
// 4 7 0.37
// 5 7 0.28
// 5 1 0.32
// 4 0 0.38
// 0 2 0.26
// 3 7 0.39
// 1 3 0.29
// 7 2 0.34
// 6 2 0.40
// 3 6 0.52
// 6 0 0.58
// 6 4 0.93
// 8 vertices, 13 edges
// 0: 0->2 0.26000
// 1: 1->3 0.29000
// 2:
// 3: 3->7 0.39000 3->6 0.52000
// 4: 4->7 0.37000 4->0 0.38000
// 5: 5->4 0.35000 5->7 0.28000 5->1 0.32000
// 6: 6->2 0.40000 6->0 0.58000 6->4 0.93000
// 7: 7->2 0.34000

// 5
// 5 to 0 (2.44)  4->0 0.38000   6->4 0.93000   3->6 0.52000   1->3 0.29000   5->1 0.32000
// 5 to 1 (0.32)  5->1 0.32000
// 5 to 2 (2.77)  7->2 0.34000   4->7 0.37000   6->4 0.93000   3->6 0.52000   1->3 0.29000   5->1 0.32000
// 5 to 3 (0.61)  1->3 0.29000   5->1 0.32000
// 5 to 4 (2.06)  6->4 0.93000   3->6 0.52000   1->3 0.29000   5->1 0.32000
// 5 to 5 (0.00)
// 5 to 6 (1.13)  3->6 0.52000   1->3 0.29000   5->1 0.32000
// 5 to 7 (2.43)  4->7 0.37000   6->4 0.93000   3->6 0.52000   1->3 0.29000   5->1 0.32000

func main() {
	g := graph.InitEdgeWeightedDigraph()
	fmt.Println(g)

	var s int
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatalln(err)
	}

	lp := AcyclicLP.New(g, s)

	// find shortest path from s to each other vertex in DAG
	for v := 0; v < g.V; v++ {
		if lp.HasPathTo(v) {
			fmt.Printf("%d to %d (%.2f)  ", s, v, lp.DistTo(v))
			for _, e := range lp.PathTo(v) {
				fmt.Print(e, "   ")
			}
			fmt.Println()
		} else {
			fmt.Printf("%d to %d         no path\n", s, v)
		}
	}
}
