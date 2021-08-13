package main

import (
	minPQ "github.com/lee-hen/Algorithms/2_sorting/27_cube_sum/min_pq"

	"fmt"
)

/******************************************************************************
 *  Dependencies: min_pq.go
 *
 *  Print out integers of the form a^3 + b^3 in sorted order, where
 *  0 <= a <= b <= n.
 *
 *  % go run main.go 12
 *  0 = 0^3 + 0^3
 *  1 = 0^3 + 1^3
 *  2 = 1^3 + 1^3
 *  8 = 0^3 + 2^3
 *  9 = 1^3 + 2^3
 *  ...
 *  1729 = 9^3 + 10^3
 *  1729 = 1^3 + 12^3
 *  ...
 *  3456 = 12^3 + 12^3
 *
 *  Remarks
 *  -------
 *   - Easily extends to handle sums of the form f(a) + g(b)
 *   - Prints out a sum more than once if it can be obtained
 *     in more than one way, e.g., 1729 = 9^3 + 10^3 = 1^3 + 12^3
 *
 ******************************************************************************/

func CubeSum(i, j int) minPQ.CubeSum {
	return minPQ.CubeSum {
		I:i,
		J: j,
		Sum: i*i*i + j*j*j,
	}
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
	}

	cubeSums := make([]minPQ.CubeSum, 0)

	for idx := 0; idx <= n; idx++ {
		cubeSums = append(cubeSums, CubeSum(idx, idx))
	}

	pq := minPQ.NewMinPQ(cubeSums)

	for !pq.IsEmpty() {
		s := pq.DelMin()
		fmt.Printf( "%d = %d^3 + %d^3\n", s.Sum, s.I, s.J)
		if s.J < n {
			pq.Insert(CubeSum(s.I, s.J+1))
		}
	}
}
