package main

import (
	gen "github.com/lee-hen/Algorithms/4_graphs/14_digraph_generator"
	"log"

	"fmt"
)

func main() {
	var v, e int
	_, err := fmt.Scan(&v, &e)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("complete graph")
	fmt.Println(gen.Complete(v))
	fmt.Println()

	fmt.Println("simple")
	fmt.Println(gen.Simple(v, e))
	fmt.Println()

	fmt.Println("path")
	fmt.Println(gen.Path(v))
	fmt.Println()

	fmt.Println("cycle")
	fmt.Println(gen.Cycle(v))
	fmt.Println()

	fmt.Println("Erdos-Renyi")
	p := float64(e) / (float64(v)*float64(v-1)/2.0)
	fmt.Println(gen.SimpleProb(v, p))
	fmt.Println()

	fmt.Println("Eulierian path")
	fmt.Println(gen.EulerianPath(v, e))
	fmt.Println()

	fmt.Println("Eulierian cycle")
	fmt.Println(gen.EulerianCycle(v, e))
	fmt.Println()

	fmt.Println("binary tree")
	fmt.Println(gen.BinaryTree(v))
	fmt.Println()

	fmt.Println("tournament")
	fmt.Println(gen.Tournament(v))
	fmt.Println()

	fmt.Println("DAG")
	fmt.Println(gen.Dag(v, e))
	fmt.Println()

	fmt.Println("rooted-in DAG")
	fmt.Println(gen.RootedInDAG(v, e))
	fmt.Println()

	fmt.Println("rooted-out DAG")
	fmt.Println(gen.RootedOutDAG(v, e))
	fmt.Println()

	fmt.Println("rooted-in tree")
	fmt.Println(gen.RootedInTree(v))
	fmt.Println()

	fmt.Println("rooted-out DAG")
	fmt.Println(gen.RootedOutTree(v))
	fmt.Println()
}
