package main

import (
	topological "github.com/lee-hen/Algorithms/4_graphs/27_topological"
	symbolGraph "github.com/lee-hen/Algorithms/4_graphs/30_symbol_digraph"

	"fmt"
	"os"
)

func main()  {
	pwd, _ := os.Getwd()
	f, _  := os.Open(pwd + "/data/jobs.txt")
	defer f.Close()
	sg := symbolGraph.New(f, "/")
	g := sg.Graph()
	fmt.Print(g)
	fmt.Println("----------------------------")
	topologi := topological.NewDigraph(sg.Graph())
	for _, v := range topologi.Order() {
		fmt.Println(sg.NameOf(v))
	}
}
