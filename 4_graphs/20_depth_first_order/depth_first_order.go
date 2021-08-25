package depth_first_order

import graph "github.com/lee-hen/Algorithms/4_graphs/13_digraph"

type DepthFirstOrder struct {
	marked map[int]bool  // marked[v] = has v been marked in dfs?
	Pre map[int]int      // pre[v]    = preorder  number of v
	Post map[int]int     // post[v]   = postorder number of v
	PreOrder []int       // vertices in preorder
	PostOrder []int      // vertices in postorder
	preCounter int       // counter or preorder numbering
	postCounter int      // counter for postorder numbering
}

// NewDigraph
// Determines a depth-first order for the digraph g.
func NewDigraph(g *graph.Digraph) *DepthFirstOrder {
	search := &DepthFirstOrder{}
	search.Pre = make(map[int]int)
	search.Post = make(map[int]int)
	search.PostOrder = make([]int, 0)
	search.PreOrder = make([]int, 0)
	search.marked = make(map[int]bool)

	for v := 0; v < g.V; v++ {
		if !search.marked[v] {
			search.dfs(g, v)
		}
	}

	return search
}

//
//func NewEdgeWeightedDigraph(g *graph.EdgeWeightedDigraph) *DepthFirstOrder {
//	search := &DepthFirstOrder{}
//	search.Pre = make(map[int]int)
//	search.Post = make(map[int]int)
//	search.PostOrder = make([]int, 0)
//	search.PreOrder = make([]int, 0)
//	search.marked = make(map[int]bool)
//
//	for v := 0; v < g.V; v++ {
//		if !search.marked[v] {
//			search.dfsEdgeWeighted(g, v)
//		}
//	}
//
//	return search
//}
//
//func (search *DepthFirstOrder) dfsEdgeWeighted(g *graph.EdgeWeightedDigraph, v int) {
//	search.marked[v] = true
//	search.Pre[v] = search.preCounter
//	search.preCounter++
//	search.PreOrder = append(search.PreOrder, v)
//
//	for i := len(g.Adj(v))-1; i >= 0; i-- {
//		w := g.Adj(v)[i]
//		if !search.marked[w] {
//			search.dfsEdgeWeighted(g, w)
//		}
//	}
//
//	search.Post[v] = search.postCounter
//	search.postCounter++
//	search.PostOrder = append(search.PostOrder, v)
//
//}

func (search *DepthFirstOrder) dfs(g *graph.Digraph, v int) {
	search.marked[v] = true
	search.Pre[v] = search.preCounter
	search.preCounter++
	search.PreOrder = append(search.PreOrder, v)

	for i := len(g.Adj(v))-1; i >= 0; i-- {
		w := g.Adj(v)[i]
		if !search.marked[w] {
			search.dfs(g, w)
		}
	}

	search.Post[v] = search.postCounter
	search.postCounter++
	search.PostOrder = append(search.PostOrder, v)

}

func (search *DepthFirstOrder) ReversePost() []int {
	reverse := make([]int, 0)
	for i := len(search.PostOrder)-1; i >= 0; i-- {
		v := search.PostOrder[i]
		reverse = append(reverse, v)
	}
	return reverse
}
