package main

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"
	gen "github.com/lee-hen/Algorithms/4_graphs/02_graph_generator"
	"github.com/lee-hen/Algorithms/util"

	"fmt"
	"log"
)

//  Identifies bridge edges and prints them out. This decomposes
//  a directed graph into two-edge connected components.
//  Runs in O(E + V) time.

//  Key quantity:  low[v] = minimum DFS preorder number of v
//  and the set of vertices w for which there is a back edge (x, w)
//  with x a descendant of v and w an ancestor of v.

//  Note: code assumes no parallel edges, e.g., two parallel edges
//  would be (incorrectly) identified as bridges.

// 4.1.36 Two-edge connectivity.
// A bridge in a graph is an edge that, if removed, would increase the number of connected components.
// A graph that has no bridges is said to be two-edge connected.
// Develop a DFS-based data type for determining whether a given graph is edge connected.


var bridges, cnt int
var pre, low map[int]int

func Bridge(g *graph.Graph) {
	low = make(map[int]int)
	pre = make(map[int]int)

	for v := 0; v < g.V; v++ {
		low[v] = -1
		pre[v] = -1
	}

	for v := 0; v < g.V; v++ {
		if pre[v] == -1 {
			dfs(g, v, v)
		}
	}
}

func Components() int { return bridges + 1 }

func dfs(g *graph.Graph, p, v int) {
	pre[v] = cnt
	cnt++
	low[v] = pre[v]

	for _, w := range g.Adj(v) {
		if pre[w] == -1 {
			dfs(g, v, w)
			low[v] = util.Min(low[v], low[w])
			if low[w] == pre[w] {
				fmt.Println(v, "-", w, "is a bridge")
				bridges++
			}
		} else if w != p {
			low[v] = util.Min(low[v], pre[w])
		}
	}
}

func main() {
	var v, e int
	_, err := fmt.Scan(&v, &e)
	if err != nil {
		log.Fatalln(err)
	}

	g := gen.Simple(v, e)
	fmt.Println(g)

	//input := [][]int{
	//	{1, 2},
	//	{2},
	//	{3},
	//	{4, 5},
	//	{5},
	//	{},
	//}
	//
	//g := graph.NewGraph(len(input))
	//for v := range input {
	//	for _, w := range input[v] {
	//		g.AddEdge(v, w)
	//	}
	//}

	Bridge(g)
	fmt.Println("Edge connected components =", Components())
}
