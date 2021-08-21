package main

import (
	suffix "github.com/lee-hen/Algorithms/5_context_or_beyond/03_suffix_array"
	"github.com/lee-hen/Algorithms/util"

	"fmt"
)

func main() {
	s := "ABRACADABRA!"
	suffix := suffix.SuffixArray(s)

	fmt.Println("  i ind lcp rnk select")
	fmt.Println("---------------------------")

	for i := 0; i < len(s); i++ {
		index := suffix.Index(i)
		ith := "\"" + s[index: util.Min(index + 50, len(s))] + "\""
		if s[index:] != suffix.Select(i) {
			panic(fmt.Errorf("s[%d:] != suffix.Select(%d)", index, i))
		}
		rank := suffix.Rank(s[index:])
		if i == 0 {
			fmt.Printf("%3d %3d %3s %3d %s\n", i, index, "-", rank, ith)
		} else {
			lcp := suffix.Lcp(i)
			fmt.Printf("%3d %3d %3d %3d %s\n", i, index, lcp, rank, ith)
		}
	}
}
