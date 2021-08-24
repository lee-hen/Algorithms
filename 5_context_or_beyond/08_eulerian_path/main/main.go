package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	gen "github.com/lee-hen/Algorithms/4_graphs/02_graph_generator"
	eulerianPath "github.com/lee-hen/Algorithms/5_context_or_beyond/08_eulerian_path"

	"fmt"
	"log"
	"math/rand"
	"time"
)

func unitTest(g *graph.Graph, description string) {
	log.Println(description)
	log.Println("-------------------------------------")
	log.Println(g)

	euler := eulerianPath.New(g)
	log.Print("Eulerian path: ")
	if euler.HasEulerianPath() {
		for _, v := range euler.Path() {
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

	// add one random edge
	g3 := graph.CloneGraph(g2)
	rand.Seed(time.Now().UnixNano())

	g3.AddEdge(rand.Intn(v), rand.Intn(v))
	unitTest(g3, "one random edge added to Eulerian path")

	// self loop
	g4 := graph.NewGraph(v)
	v4 := rand.Intn(v)
	g4.AddEdge(v4, v4)
	unitTest(g4, "single self loop")

	// single edge
	g5 := graph.NewGraph(v)
	g5.AddEdge(rand.Intn(v), rand.Intn(v))
	unitTest(g5, "single edge")

	// empty graph
	g6 := graph.NewGraph(v)
	unitTest(g6, "empty graph")

	// random graph
	g7 := gen.Simple(v, e)
	unitTest(g7, "simple graph")
}
