package main

import (
	gen "github.com/lee-hen/Algorithms/4_graphs/02_graph_generator"
	bipartite "github.com/lee-hen/Algorithms/4_graphs/09_bipartite_x"

	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	var v1, v2, e, f int

	_, err := fmt.Scan(&v1, &v2, &e, &f)
	if err != nil {
		log.Fatalln(err)
	}

	// create random bipartite graph with V1 vertices on left side,
	// V2 vertices on right side, and E edges; then add F random edges
	g := gen.Bipartite(v1, v2, e)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < f; i++ {
		v := rand.Intn(v1 + v2)
		w := rand.Intn(v1 + v2)
		g.AddEdge(v, w)
	}
	fmt.Println(g)

	b := bipartite.New(g)

	if b.IsBipartite() {
		fmt.Println("Graph is bipartite")
		for v := 0; v < g.V; v++ {
			fmt.Println(v, ":", b.Color(v))
		}
	} else {
		fmt.Println("Graph has an odd-length cycle: ")
		for _, x := range b.OddCycle() {
			fmt.Print(x, ' ')
		}
		fmt.Println()
	}
}
