package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/24_edge_weighted_digraph"
	dij "github.com/lee-hen/Algorithms/4_graphs/39_dijkstra_sp"
	"github.com/lee-hen/Algorithms/util"

	"fmt"
	"math"
)

// All-pairs shortest paths. Given an edge-weighted digraph, support queries of the form
// Given a source vertex s and a target vertex t, is there a path from s to t? If so, find a shortest such path (one whose total weight is minimal).

var all []*dij.DijkstraSP

// DijkstraAllPairsSP
// Computes a shortest paths tree from each vertex to to every other vertex in
//  the edge-weighted digraph g.
func DijkstraAllPairsSP(g *graph.EdgeWeightedDigraph) {
	all = make([]*dij.DijkstraSP, g.V, g.V)
	for v := 0; v < g.V; v++ {
		all[v] = dij.New(g, v)
	}
}


// Path
// Returns a shortest path from vertex s to vertex t.
func Path(s, t int) util.DirectedEdgeStack {
	return all[s].PathTo(t)
}

// HasPath
// Returns true if there is a path from the source vertex s to vertex v.
func HasPath(s, t int) bool {
	return Dist(s, t) < math.MaxFloat64
}

// Dist
// Returns the length of a shortest path from vertex s to vertex t.
func Dist(s, t int) float64 {
	return all[s].DistTo(t)
}

//
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

// 0      1      2      3      4      5      6      7
// 0:   0.00   1.05   0.26   0.99   0.38   0.73   1.51   0.60
// 1:   1.39   0.00   1.21   0.29   1.74   1.83   0.81   1.55
// 2:   1.83   0.94   0.00   0.73   0.97   0.62   1.25   0.34
// 3:   1.10   1.86   0.92   0.00   1.45   1.54   0.52   1.26
// 4:   1.86   0.67   1.68   0.76   0.00   0.35   1.28   0.37
// 5:   1.71   0.32   1.53   0.61   0.35   0.00   1.13   0.28
// 6:   0.58   1.34   0.40   1.13   0.93   1.02   0.00   0.74
// 7:   1.49   0.60   1.31   0.39   0.63   0.28   0.91   0.00

// 0 to 0 ( 0.00)
// 0 to 1 ( 1.05)  0  1  2
// 0 to 2 ( 0.26)  0
// 0 to 3 ( 0.99)  0  1  2
// 0 to 4 ( 0.38)  0
// 0 to 5 ( 0.73)  0  1
// 0 to 6 ( 1.51)  0  1  2  3
// 0 to 7 ( 0.60)  0  1
// 1 to 0 ( 1.39)  0  1  2
// 1 to 1 ( 0.00)
// 1 to 2 ( 1.21)  0  1  2
// 1 to 3 ( 0.29)  0
// 1 to 4 ( 1.74)  0  1  2
// 1 to 5 ( 1.83)  0  1  2  3  4
// 1 to 6 ( 0.81)  0  1
// 1 to 7 ( 1.55)  0  1  2  3
// 2 to 0 ( 1.83)  0  1  2  3
// 2 to 1 ( 0.94)  0  1  2
// 2 to 2 ( 0.00)
// 2 to 3 ( 0.73)  0  1
// 2 to 4 ( 0.97)  0  1  2
// 2 to 5 ( 0.62)  0  1
// 2 to 6 ( 1.25)  0  1  2
// 2 to 7 ( 0.34)  0
// 3 to 0 ( 1.10)  0  1
// 3 to 1 ( 1.86)  0  1  2  3  4
// 3 to 2 ( 0.92)  0  1
// 3 to 3 ( 0.00)
// 3 to 4 ( 1.45)  0  1
// 3 to 5 ( 1.54)  0  1  2  3
// 3 to 6 ( 0.52)  0
// 3 to 7 ( 1.26)  0  1  2
// 4 to 0 ( 1.86)  0  1  2  3
// 4 to 1 ( 0.67)  0  1
// 4 to 2 ( 1.68)  0  1  2  3
// 4 to 3 ( 0.76)  0  1
// 4 to 4 ( 0.00)
// 4 to 5 ( 0.35)  0
// 4 to 6 ( 1.28)  0  1  2
// 4 to 7 ( 0.37)  0
// 5 to 0 ( 1.71)  0  1  2  3
// 5 to 1 ( 0.32)  0
// 5 to 2 ( 1.53)  0  1  2  3
// 5 to 3 ( 0.61)  0  1
// 5 to 4 ( 0.35)  0
// 5 to 5 ( 0.00)
// 5 to 6 ( 1.13)  0  1  2
// 5 to 7 ( 0.28)  0
// 6 to 0 ( 0.58)  0
// 6 to 1 ( 1.34)  0  1  2  3
// 6 to 2 ( 0.40)  0
// 6 to 3 ( 1.13)  0  1  2
// 6 to 4 ( 0.93)  0
// 6 to 5 ( 1.02)  0  1  2
// 6 to 6 ( 0.00)
// 6 to 7 ( 0.74)  0  1
// 7 to 0 ( 1.49)  0  1  2
// 7 to 1 ( 0.60)  0  1
// 7 to 2 ( 1.31)  0  1  2
// 7 to 3 ( 0.39)  0
// 7 to 4 ( 0.63)  0  1
// 7 to 5 ( 0.28)  0
// 7 to 6 ( 0.91)  0  1
// 7 to 7 ( 0.00)

func main() {
	g := graph.InitEdgeWeightedDigraph()
	fmt.Println(g)

	DijkstraAllPairsSP(g)

	// print all-pairs shortest path distances
	fmt.Printf("  ")
	for  v := 0; v < g.V; v++ {
		fmt.Printf("%6d ", v)
	}

	fmt.Println()
	for  v := 0; v < g.V; v++ {
		fmt.Printf("%3d: ", v)
		for w := 0; w < g.V; w++ {
			if HasPath(v, w) {
				fmt.Printf("%6.2f ", Dist(v, w))
			} else {
				fmt.Printf("  Inf ")
			}
		}
		fmt.Println()
	}
	fmt.Println()

	// print all-pairs shortest paths
	for v := 0; v < g.V; v++ {
		for w := 0; w < g.V; w++ {
			if HasPath(v, w) {
				fmt.Printf("%d to %d (%5.2f)  ", v, w, Dist(v, w))
				for e := range Path(v, w) {
					fmt.Print(e, "  ")
				}

				fmt.Println()
			} else {
				fmt.Printf("%d to %d no path\n", v, w)
			}
		}
	}
}
