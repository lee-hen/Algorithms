package main

import (
	"fmt"
	"github.com/lee-hen/Algorithms/util"
	"os"
)

// Lcs
// Compute length of LCS for all subproblems.
func Lcs(x, y string) string {
	m, n := len(x), len(y)
	opt := make([][]int, m+1, m+1)

	for i := range opt {
		opt[i] = make([]int, n+1, n+1)
	}

	for i := m-1; i >= 0; i-- {
		for j := n-1; j >= 0; j-- {
			if x[i] == y[j] {
				opt[i][j] = opt[i+1][j+1] + 1
			} else {
				opt[i][j] = util.Max(opt[i+1][j], opt[i][j+1])
			}
		}
	}

	lcs := ""
	var i, j int
	for i < m && j < n {
		if x[i] == y[j] {
			lcs += string(x[i])
			i++
			j++
		} else if opt[i+1][j] >= opt[i][j+1] {
			i++
		} else {
			j++
		}
	}

	return lcs
}

func main() {
	lcs := Lcs(os.Args[1], os.Args[2])
	fmt.Println(lcs)
}
