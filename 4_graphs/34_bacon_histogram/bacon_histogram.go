package main

import (
	"fmt"
	paths "github.com/lee-hen/Algorithms/4_graphs/06_breadth_first_paths"
	symbolGraph "github.com/lee-hen/Algorithms/4_graphs/11_symbol_graph"
	"github.com/lee-hen/Algorithms/util"
	"os"
)

// Reads in a data file containing movie records (a movie followed by a list
// of actors appearing in that movie), and runs breadth first search to
// find the shortest distance from the source (Kevin Bacon) to each other
// actor and movie. After computing the Kevin Bacon numbers, the programs
// prints a histogram of the number of actors with each Kevin Bacon number.

func main() {
	pwd, _ := os.Getwd()
	f, _  := os.Open(pwd + "/data/movies.txt")
	defer f.Close()
	sg := symbolGraph.New(f, "/")

	var source = "Bacon, Kevin"

	g := sg.Graph()

	if !sg.Contains(source) {
		fmt.Println(source, "not in database")
		return
	}

	// run breadth-first search from s
	s := sg.IndexOf(source)
	bfs := paths.BreadthFirstPaths(g, s)

	// compute histogram of Kevin Bacon numbers - 100 for infinity
	maxBacon := 100
	hist := make([]int, maxBacon+1, maxBacon+1)
	for v := 0; v < g.V; v++ {
		if bfs.HasPathTo(v) {
			bacon := util.Min(maxBacon, bfs.DistTo(v))
			hist[bacon]++

			// to print actors and movies with large bacon numbers
			if bacon/2 >= 7 && bacon < maxBacon {
				fmt.Printf("%d %s\n", bacon/2, sg.NameOf(v))
			}
		}
	}

	// print out histogram - even indices are actors
	for i := 0; i < maxBacon; i+=2 {
		if hist[i] == 0 {
			break
		}

		fmt.Printf("%3d %8d\n", i/2, hist[i])
	}

	fmt.Printf("Inf %8d\n", hist[maxBacon])
}
