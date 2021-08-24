package main

import (
	"bufio"
	"fmt"
	path "github.com/lee-hen/Algorithms/4_graphs/06_breadth_first_paths"
	symbolGraph "github.com/lee-hen/Algorithms/4_graphs/11_symbol_graph"
	"log"
	"os"
	"strings"
)

func main()  {
	pwd, _ := os.Getwd()
	f, _  := os.Open(pwd + "/data/movies.txt")
	sg := symbolGraph.New(f, "/")

	reader := bufio.NewReader(os.Stdin)
	source, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalln(err)
	}
	source = strings.Replace(source, "\n", "", -1)

	if !sg.Contains(source) {
		log.Fatalln(source + " not in database.")
	}

	g, s := sg.Graph(), sg.IndexOf(source)
	bfs := path.BreadthFirstPaths(g, s)

	for {
		sink, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}
		sink = strings.Replace(sink, "\n", "", -1)

		if sg.Contains(sink) {
			t := sg.IndexOf(sink)
			if bfs.HasPathTo(t) {
				for _, v := range bfs.PathTo(t) {
					fmt.Println("   ", sg.NameOf(v))
				}
			} else {
				fmt.Println("Not connected")
			}
		} else {
			fmt.Println("   Not in database.")
		}
	}
}
