package main

import (
	gen "github.com/lee-hen/Algorithms/4_graphs/14_digraph_generator"
	directedEdge "github.com/lee-hen/Algorithms/4_graphs/22_directed_edge"
	graph "github.com/lee-hen/Algorithms/4_graphs/24_edge_weighted_digraph"
	topologicalX "github.com/lee-hen/Algorithms/4_graphs/28_topological_x"

	"fmt"
	"log"
	"math/rand"
)

// 20 14 5

func main() {
	var v, e, f int
	_, err := fmt.Scan(&v, &e, &f)
	if err != nil {
		log.Fatalln(err)
	}

	g1 := gen.Dag(v, e)
	g2 := graph.NewEdgeWeightedDigraph(v)
	for v := 0; v < g1.V; v++ {
		for i := len(g1.Adj(v))-1; i >= 0; i-- {
			w := g1.Adj(v)[i]
			g2.AddEdge(directedEdge.NewEdge(v, w, 0.0) )
		}
	}

	V := v
	for i := 0; i < f; i++ {
		v := rand.Intn(V)
		w := rand.Intn(V)
		g1.AddEdge(v, w)
		g2.AddEdge(directedEdge.NewEdge(v, w, 0.0))
	}

	fmt.Println(g1)
	topological1 := topologicalX.NewDigraph(g1)
	if !topological1.HasOrder() {
		fmt.Println("Not a DAG")
	} else {
		fmt.Print("Topological order: ")

		for i := len(topological1.Order())-1; i >= 0; i-- {
			v := topological1.Order()[i]
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

	// find a directed cycle
	topological2 := topologicalX.NewEdgeWeightedDigraph(g2)
	if !topological2.HasOrder() {
		fmt.Println("Not a DAG")
	} else {
		fmt.Print("Topological order: ")
		for i := len(topological2.Order())-1; i >= 0; i-- {
			v := topological2.Order()[i]
			fmt.Print(v, " ")
		}
		fmt.Println()
	}
}
