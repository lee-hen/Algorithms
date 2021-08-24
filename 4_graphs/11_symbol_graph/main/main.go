package main

import (
	symbolGraph "github.com/lee-hen/Algorithms/4_graphs/11_symbol_graph"

	"fmt"
	"os"
)

func main()  {
	pwd, _ := os.Getwd()
	f, _  := os.Open(pwd + "/data/routes.txt")
	defer f.Close()
	sg := symbolGraph.New(f, " ")
	fmt.Println(sg.Graph())
}
