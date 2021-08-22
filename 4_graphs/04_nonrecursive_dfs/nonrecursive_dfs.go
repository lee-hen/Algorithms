package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	"github.com/lee-hen/Algorithms/util"

	"fmt"
	"log"
)

func NonRecursiveDFS(g *graph.Graph, s int) map[int]bool {
	marked := make(map[int]bool)

	// depth-first search using an explicit stack
	stack := make(util.Stack, 0)
	marked[s] = true
	stack.Push(s)

	var next int
	for !stack.IsEmpty() {
		v := stack.Peek()
		if next < g.Degree(v){
			w := g.Adj(v)[next]
			if !marked[w] {
				marked[w] = true
				stack.Push(w)
			}
			next++
		} else {
			stack.Pop()
			next = 0
		}
	}
	return marked
}

// 13
// 13
// 0 5
// 4 3
// 0 1
// 9 12
// 6 4
// 5 4
// 0 2
// 11 12
// 9 10
// 0 6
// 7 8
// 9 11
// 5 3
// 9
// 9 10 11 12
// https://algs4.cs.princeton.edu/41graph/tinyG.txt

func main()  {
	g := graph.InitGraph()

	var s int
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatalln(err)
	}

	marked := NonRecursiveDFS(g, s)
	for v := 0; v < g.V; v++ {
		if marked[v] {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
}
