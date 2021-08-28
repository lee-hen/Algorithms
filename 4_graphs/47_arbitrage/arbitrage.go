package main

import (
	directedEdge "github.com/lee-hen/Algorithms/4_graphs/22_directed_edge"
	graph "github.com/lee-hen/Algorithms/4_graphs/24_edge_weighted_digraph"
	blf "github.com/lee-hen/Algorithms/4_graphs/46_bellman_ford_sp"

	"fmt"
	"log"
	"math"
)

// Proposition Z. The arbitrage problem is a negative-cycle-detection problem in edge-weighted digraphs.
// Proof: Replace each weight by its logarithm, negated. With this change, computing path weights by multiplying edge weights in the original problem corresponds to adding them in the transformed problem.
// Specifically, any product w1w2 . . . wk corresponds to a sum −ln(w1) − ln(w2) − . . . − ln(wk).
// The transformed edge weights might be negative or positive, a path from v to w gives a way of converting from currency v to currency w, and any negative cycle is an arbitrage opportunity.


// 5
// USD 1      0.741  0.657  1.061  1.005
// EUR 1.349  1      0.888  1.433  1.366
// GBP 1.521  1.126  1      1.614  1.538
// CHF 0.942  0.698  0.619  1      0.953
// CAD 0.995  0.732  0.650  1.049  1

// 1000.00000 CAD =  995.00000 USD
// 995.00000 EUR = 1359.17000 CAD
// 1359.17000 USD = 1007.14497 EUR

func main() {
	// V currencies
	var V int
	_, err := fmt.Scan(&V)
	if err != nil {
		log.Fatalln(err)
	}

	name := make([]string, V)

	// create complete network
	g := graph.NewEdgeWeightedDigraph(V)
	for v := 0; v < V; v++ {
		_, err = fmt.Scan(&name[v])
		if err != nil {
			log.Fatalln(err)
		}

		for w := 0; w < V; w++ {
			var rate float64

			_, err = fmt.Scan(&rate)
			if err != nil {
				log.Fatalln(err)
			}

			g.AddEdge(directedEdge.NewEdge(v, w, -math.Log(rate)))
		}
	}

	// find negative cycle
	spt := blf.New(g, 0)
	if spt.HasNegativeCycle() {
		stake := 1000.0
		for _, e := range spt.NegativeCycle() {
			fmt.Printf("%10.5f %s ", stake, name[e.From()])
			stake *= math.Exp(-e.Weight())

			fmt.Printf("= %10.5f %s\n", stake, name[e.To()])
		}
	} else {
		fmt.Println("No arbitrage opportunity")
	}
}
