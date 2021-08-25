package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	gen "github.com/lee-hen/Algorithms/4_graphs/14_digraph_generator"
	eulerianCycle "github.com/lee-hen/Algorithms/5_context_or_beyond/09_directed_eulerian_cycle"
	"github.com/lee-hen/Algorithms/util"

	"fmt"
	"log"
	"math/rand"
)

func unitTest(g *graph.Digraph, description string) {
	log.Println(description)
	log.Println("-------------------------------------")
	log.Println(g)

	euler := eulerianCycle.New(g)
	log.Print("Eulerian cycle: ")
	if euler.HasEulerianCycle() {
		for _, v := range euler.Cycle() {
			log.Print(v, " ")
		}
	} else {
		log.Println("none")
	}
	log.Println()
}

func main() {
	var v, e int
	_, err := fmt.Scan(&v, &e)
	if err != nil {
		log.Fatalln(err)
	}

	// Eulerian cycle
	g1 := gen.EulerianCycle(v, e)
	unitTest(g1, "Eulerian cycle")

	// Eulerian path
	g2 := gen.EulerianPath(v, e)
	unitTest(g2, "Eulerian path")

	// empty digraph
	g6 := graph.NewDigraph(v)
	unitTest(g6, "empty graph")

	// self loop
	g4 := graph.NewDigraph(v)
	v4 := rand.Intn(v)
	g4.AddEdge(v4, v4)
	unitTest(g4, "single self loop")

	// union of two disjoint cycles
	h1 := gen.EulerianCycle(v/2, e/2)
	h2 := gen.EulerianCycle(v-v/2, e-e/2)
	perm := make([]int, v, v)
	for i := 0; i < v; i++ {
		perm[i] = i
	}
	util.ShuffleIntSlice(perm)
	g5 := graph.NewDigraph(v)
	for v := 0; v < h1.V; v++ {
		for _, w := range h1.Adj(v) {
			g5.AddEdge(perm[v], perm[w])
		}
	}
	for v := 0; v < h2.V; v++ {
		for _, w := range h2.Adj(v) {
			g5.AddEdge(perm[g5.V/2 + v], perm[g5.V/2 + w])
		}
	}
	unitTest(g5, "Union of two disjoint cycles")

	// random graph
	g7 := gen.Simple(v, e)
	unitTest(g7, "simple graph")
}