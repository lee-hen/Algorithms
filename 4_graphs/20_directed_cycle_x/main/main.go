package main

import (
	gen "github.com/lee-hen/Algorithms/4_graphs/14_digraph_generator"
	cycle "github.com/lee-hen/Algorithms/4_graphs/20_directed_cycle_x"

	"fmt"
	"log"
	"math/rand"
	"time"
)

func main() {
	var v, e, f int
	_, err := fmt.Scan(&v, &e, &f)
	if err != nil {
		log.Fatalln(err)
	}
	g := gen.Dag(v, e)

	rand.Seed(time.Now().UnixNano())

	V := g.V
	for i := 0; i < f; i++ {
		v := rand.Intn(V)
		w := rand.Intn(V)
		g.AddEdge(v, w)
	}

	fmt.Println(g)

	finder := cycle.New(g)
	if finder.HasCycle() {
		fmt.Println("Directed cycle: ")
		for _, v := range finder.Cycle() {
			fmt.Print(v, " ")
		}
		fmt.Println()
	} else {
		fmt.Println("Graph is acyclic")
	}
}
