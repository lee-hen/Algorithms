package main

import (
	directedEdge "github.com/lee-hen/Algorithms/4_graphs/22_directed_edge"
	graph "github.com/lee-hen/Algorithms/4_graphs/24_edge_eeighted_digraph"
	cycle "github.com/lee-hen/Algorithms/4_graphs/25_edge_weighted_directed_cycle"
	"github.com/lee-hen/Algorithms/util"

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

	g := graph.NewEdgeWeightedDigraph(v)

	rand.Seed(time.Now().UnixNano())
	vertices := make([]int, v)
	for i := 0; i < v; i++ {
		vertices[i] = i
	}
	util.ShuffleIntSlice(vertices)
	V := v
	for i := 0; i < e; i++ {
		v := rand.Intn(V)
		w := rand.Intn(V)
		for v >= w {
			v = rand.Intn(V)
			w = rand.Intn(V)
		}

		weight := rand.Float64()

		g.AddEdge(directedEdge.NewEdge(v, w, weight))
	}

	for i := 0; i < f; i++ {
		v := rand.Intn(V)
		w := rand.Intn(V)
		weight := rand.Float64()
		g.AddEdge(directedEdge.NewEdge(v, w, weight))
	}
	fmt.Println(g)

	finder := cycle.New(g)
	if finder.HasCycle() {
		fmt.Println("Cycle: ")
		for _, e := range finder.Cycle() {
			fmt.Print(e, " ")
		}
		fmt.Println()
	} else {
		fmt.Println("No directed cycle")
	}
}
