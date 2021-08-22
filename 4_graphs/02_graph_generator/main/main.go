package main

import (
	gen "github.com/lee-hen/Algorithms/4_graphs/02_graph_generator"
	"log"

	"fmt"
)

func main() {
	var v, e int
	_, err := fmt.Scan(&v, &e)
	if err != nil {
		log.Fatalln(err)
	}

	v1 := v/2
	v2 := v-v1

	fmt.Println("complete graph")
	fmt.Println(gen.Complete(v))
	fmt.Println()

	fmt.Println("simple")
	fmt.Println(gen.Simple(v, e))
	fmt.Println()

	fmt.Println("Erdos-Renyi")
	p := float64(e) / (float64(v)*float64(v-1)/2.0)
	fmt.Println(gen.SimpleProb(v, p))
	fmt.Println()

	fmt.Println("complete bipartite")
	fmt.Println(gen.CompleteBipartite(v1, v2))
	fmt.Println()

	fmt.Println("bipartite")
	fmt.Println(gen.Bipartite(v1, v2, e))
	fmt.Println()

	fmt.Println("Erdos Renyi bipartite");
	 q := float64(e) / float64(v1*v2)
	fmt.Println(gen.BipartiteProb(v1, v2, q))
	fmt.Println()

	fmt.Println("path")
	fmt.Println(gen.Path(v))
	fmt.Println()

	fmt.Println("cycle")
	fmt.Println(gen.Cycle(v))
	fmt.Println()

	fmt.Println("binary tree")
	fmt.Println(gen.BinaryTree(v))
	fmt.Println()

	fmt.Println("4-regular")
	fmt.Println(gen.Regular(v, 4))
	fmt.Println()

	fmt.Println("star")
	fmt.Println(gen.Star(v))
	fmt.Println()

	fmt.Println("wheel")
	fmt.Println(gen.Wheel(v))
	fmt.Println()

	fmt.Println("tree")
	fmt.Println(gen.Tree(v))
	fmt.Println()
}
