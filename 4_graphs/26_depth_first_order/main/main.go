package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	depthFirstOrder "github.com/lee-hen/Algorithms/4_graphs/26_depth_first_order"

	"fmt"
)

// 13
// 15
// 2 3
// 0 6
// 0 1
// 2 0
// 11 12
// 9 12
// 9 10
// 9 11
// 3 5
// 8 7
// 5 4
// 0 5
// 6 4
// 6 9
// 7 6
// v  pre post
// --------------
// 0    0    8
// 1    3    2
// 2    9   10
// 3   10    9
// 4    2    0
// 5    1    1
// 6    4    7
// 7   11   11
// 8   12   12
// 9    5    6
// 10    8    5
// 11    6    4
// 12    7    3
// Preorder:  12 1 2 6 9 11 8 10 0 3 4 5 7
// Postorder: 11 2 3 4 5 6 7 0 1 10 12 8 9
// Reverse postorder: 4 5 1 12 11 10 9 6 0 3 2 7 8

func main() {
	g := graph.InitDigraph()
	dfs := depthFirstOrder.NewDigraph(g)

	fmt.Println("   v  pre post")
	fmt.Println("--------------")
	for v := 0; v < g.V; v++ {
		fmt.Printf("%4d %4d %4d\n", v, dfs.Pre[v], dfs.Post[v])
	}

	fmt.Print("Preorder:  ")
	for _, v := range dfs.Pre {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Print("Postorder: ")
	for _, v := range dfs.Post {
		fmt.Print(v, " ")
	}
	fmt.Println()

	fmt.Print("Reverse postorder: ")
	for _, v := range dfs.ReversePost() {
		fmt.Print(v, " ")
	}
	fmt.Println()
}
