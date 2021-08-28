package main

import (
	directedEdge "github.com/lee-hen/Algorithms/4_graphs/22_directed_edge"
	graph "github.com/lee-hen/Algorithms/4_graphs/24_edge_weighted_digraph"
	AcyclicLP "github.com/lee-hen/Algorithms/4_graphs/44_acyclic_lp"

	"fmt"
	"log"
)

// Parallel precedence-constrained scheduling. Given a set of jobs of specified duration to be completed,
// with precedence constraints that specify that certain jobs have to be completed before certain other jobs are begun,
// how can we schedule the jobs on identical processors (as many as needed) such that they are all completed in the minimum
// amount of time while still respecting the constraints?

// Definition. The critical path method for parallel scheduling is to proceed as follows: Create an edge-weighted DAG
// with a source s, a sink t, and two vertices for each job (a start vertex and an end vertex).
// For each job, add an edge from its start vertex to its end vertex with weight equal to its duration.
// For each precedence constraint v->w, add a zero-weight edge from the end vertex corresponding to v to the beginning vertex corresponding to w.
// Also add zero-weight edges from the source to each job’s start vertex and from each job’s end vertex to the sink. Now,
// schedule each job at the time given by the length of its longest path from the source.

// Proposition U. The critical path method solves the parallel precedence-constrained scheduling problem in linear time.
// Proof: Why does the CPM approach work? The correctness of the algorithm rests on two facts.
// First, every path in the DAG is a sequence of job starts and job finishes, separated by zero-weight precedence constraints—the
// length of any path from the source s to any vertex v in the graph is a lower bound on the start/finish time represented by v,
// because we could not do better than scheduling those jobs one after another on the same machine.
// In particular, the length of the longest path from s to the sink t is a lower bound on the finish time of all the jobs.
// Second, all the start and finish times implied by longest paths are feasible—every job starts after the finish of all the jobs where
// it appears as a successor in a precedence constraint, because the start time is the length of the longest path from the source to it.
// In particular, the length of the longest path from s to t is an upper bound on the finish time of all the jobs.
// The linear-time performance is immediate from PROPOSITION T.

// Proposition V. Parallel job scheduling with relative deadlines is a shortest-paths problem in edge-weighted digraphs (with cycles and negative weights allowed).
// Proof: Use the same construction as in PROPOSITION U, adding an edge for each deadline: if job v has to start within d time units of the start of job w,
// add an edge from v to w with negative weight d.
// Then convert to a shortest-paths problem by negating all the weights in the digraph. The proof of correctness applies,
// provided that the schedule is feasible. Determining whether a schedule is feasible is part of the computational burden, as you will see.


// 10
// 41.0
// 3
// 1 7 9
// 51.0
// 1
// 2
// 50.0
// 0
// 36.0
// 0
// 38.0
// 0
// 45.0
// 0
// 21.0
// 2
// 3 8
// 32.0
// 2
// 3 8
// 32.0
// 1
// 2
// 29.0
// 2
// 4 6

// 22 vertices, 41 edges
// 0: 0->10 41.00000
// 1: 1->11 51.00000
// 2: 2->12 50.00000
// 3: 3->13 36.00000
// 4: 4->14 38.00000
// 5: 5->15 45.00000
// 6: 6->16 21.00000
// 7: 7->17 32.00000
// 8: 8->18 32.00000
// 9: 9->19 29.00000
// 10: 10->21 0.00000 10->1 0.00000 10->7 0.00000 10->9 0.00000
// 11: 11->21 0.00000 11->2 0.00000
// 12: 12->21 0.00000
// 13: 13->21 0.00000
// 14: 14->21 0.00000
// 15: 15->21 0.00000
// 16: 16->21 0.00000 16->3 0.00000 16->8 0.00000
// 17: 17->21 0.00000 17->3 0.00000 17->8 0.00000
// 18: 18->21 0.00000 18->2 0.00000
// 19: 19->21 0.00000 19->4 0.00000 19->6 0.00000
// 20: 20->0 0.00000 20->1 0.00000 20->2 0.00000 20->3 0.00000 20->4 0.00000 20->5 0.00000 20->6 0.00000 20->7 0.00000 20->8 0.00000 20->9 0.00000
// 21:

// job   start  finish
// --------------------
// 0     0.0    41.0
// 1    41.0    92.0
// 2   123.0   173.0
// 3    91.0   127.0
// 4    70.0   108.0
// 5     0.0    45.0
// 6    70.0    91.0
// 7    41.0    73.0
// 8    91.0   123.0
// 9    41.0    70.0
// Finish time:173.0

func main() {
	// number of jobs
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		log.Fatalln(err)
	}

	// source and sink
	source := 2*n
	sink   := 2*n + 1

	// build network
	g := graph.NewEdgeWeightedDigraph(2*n + 2)
	for i := 0; i < n; i++ {
		var duration float64
		_, err = fmt.Scan(&duration)
		if err != nil {
			log.Fatalln(err)
		}

		g.AddEdge(directedEdge.NewEdge(source, i, 0.0))
		g.AddEdge(directedEdge.NewEdge(i+n, sink, 0.0))
		g.AddEdge(directedEdge.NewEdge(i, i+n,    duration))

		// precedence constraints
		var m int
		_, err = fmt.Scan(&m)
		if err != nil {
			log.Fatalln(err)
		}

		for j := 0; j < m; j++ {
			var precedent int
			_, err = fmt.Scan(&precedent)
			if err != nil {
				log.Fatalln(err)
			}
			g.AddEdge(directedEdge.NewEdge(n+i, precedent, 0.0))
		}
	}

	fmt.Println(g)

	// compute longest path
	lp := AcyclicLP.New(g, source)

	// print results
	fmt.Println(" job   start  finish")
	fmt.Println("--------------------")
	for i := 0; i < n; i++ {
		fmt.Printf("%4d %7.1f %7.1f\n", i, lp.DistTo(i), lp.DistTo(i+n))
	}
	fmt.Printf("Finish time:%7.1f\n", lp.DistTo(sink))
}
