package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"
	"github.com/lee-hen/Algorithms/util"

	"fmt"
	"log"
)

func NonrecursiveDirectedDFS(g *graph.Digraph, s int) map[int]bool {
	marked := make(map[int]bool)

	// depth-first search using an explicit stack
	stack := make(util.Stack, 0)
	marked[s] = true
	stack.Push(s)

	var next int
	for !stack.IsEmpty() {
		v := stack.Peek()
		if next < g.OutDegree(v){
			w := g.Adj(v)[next]
			if !marked[w] {
				marked[w] = true
				stack.Push(w)
				next = 0
			} else {
				next++
			}
		} else {
			stack.Pop()
			next = 0
		}
	}
	return marked
}

// 13
// 22
// 4  2
// 2  3
// 3  2
// 6  0
// 0  1
// 2  0
// 11 12
// 12  9
// 9 10
// 9 11
// 7  9
// 10 12
// 11  4
// 4  3
// 3  5
// 6  8
// 8  6
// 5  4
// 0  5
// 6  4
// 6  9
// 7  6
// 2
// 0 1 2 3 4 5
// https://algs4.cs.princeton.edu/42digraph/tinyDG.txt

func main()  {
	g := graph.InitDigraph()

	fmt.Println(g)

	var s int
	_, err := fmt.Scan(&s)
	if err != nil {
		log.Fatalln(err)
	}

	marked := NonrecursiveDirectedDFS(g, s)
	for v := 0; v < g.V; v++ {
		if marked[v] {
			fmt.Print(v, " ")
		}
	}
	fmt.Println()
}
