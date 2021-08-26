package main

import (
	symbolGraph "github.com/lee-hen/Algorithms/4_graphs/30_symbol_digraph"

	"fmt"
	"os"
)

func main()  {
	pwd, _ := os.Getwd()
	f, _  := os.Open(pwd + "/data/digraph/routes.txt")
	defer f.Close()
	sg := symbolGraph.New(f, " ")
	g := sg.Graph()
	fmt.Print(g)
	fmt.Println("----------------------------")
	for v := 0; v < g.V; v++ {
		fmt.Print(sg.NameOf(v), ": ")
		for _, w := range g.Adj(v) {
			fmt.Print(sg.NameOf(w), " ")
		}
		fmt.Println()
	}
}
