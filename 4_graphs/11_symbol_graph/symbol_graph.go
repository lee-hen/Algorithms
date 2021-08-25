package symbol_graph

import (
	graph "github.com/lee-hen/Algorithms/4_graphs/01_graph"

	"bufio"
	"bytes"
	"io"
	"strings"
)

type SymbolGraph struct {
	st map[string]int   // string -> index
	keys []string       // index  -> string
    graph *graph.Graph  // the underlying graph
}

// New
// Initializes a graph from a file using the specified delimiter.
// Each line in the file contains
// the name of a vertex, followed by a list of the names
// of the vertices adjacent to that vertex, separated by the delimiter.
func New(reader io.Reader, delimiter string) *SymbolGraph {
	tBuf := new(bytes.Buffer)
   	buf := io.TeeReader(reader, tBuf)

	fileName := bufio.NewReader(buf)
	sg := &SymbolGraph{}
	sg.st = make(map[string]int)

	// First pass builds the index by reading strings to associate
	// distinct strings with an index
	line, _, err := fileName.ReadLine()
	for ; err != io.EOF; line, _, err = fileName.ReadLine() {
		a := strings.Split(string(line), delimiter)
		for _, s := range a {
			if _, ok := sg.st[s]; !ok {
				sg.st[s] = len(sg.st)
			}
		}
	}

	// inverted index to get string keys in an array
	sg.keys = make([]string, len(sg.st), len(sg.st))

	for name := range sg.st {
		sg.keys[sg.st[name]] = name
	}

	// second pass builds the graph by connecting first vertex on each
	// line to all others
	sg.graph = graph.NewGraph(len(sg.st))

	fileName = bufio.NewReader(tBuf)
	line, _, err = fileName.ReadLine()
	for ; err != io.EOF; line, _, err = fileName.ReadLine() {
		a := strings.Split(string(line), delimiter)

		v := sg.st[a[0]]
		for i := 1; i < len(a); i++ {
			w := sg.st[a[i]]
			sg.graph.AddEdge(v, w)
		}
	}

	return sg
}

// Contains
// Does the graph contain the vertex named s
func (sg *SymbolGraph) Contains(s string) bool {
	if _, ok := sg.st[s]; ok {
		return true
	}

	return false
}

// IndexOf
// Returns the integer associated with the vertex named s
func (sg *SymbolGraph) IndexOf(s string) int {
	if _, ok := sg.st[s]; ok {
		return sg.st[s]
	}

	return -1
}

// NameOf
// Returns the name of the vertex associated with the integer v.
func (sg *SymbolGraph) NameOf(v int) string {
	return sg.keys[v]
}

// Graph
// Returns the graph assoicated with the symbol graph. It is the client's responsibility
// not to mutate the graph.
func (sg *SymbolGraph) Graph() *graph.Graph {
	return sg.graph
}
