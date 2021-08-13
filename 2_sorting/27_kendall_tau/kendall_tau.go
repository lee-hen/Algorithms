package main

import (
	inversions "github.com/lee-hen/Algorithms/2_sorting/09_inversions"
	"github.com/lee-hen/Algorithms/util"

	"fmt"
	"log"
)

// Generate two random permutations of size N and compute their
// Kendall tau distance (number of inversions).
// https://en.wikipedia.org/wiki/Kendall_tau_distance

//b 9 7 5 8 1 6 3 2 4 0
//  5 6 7 3 1 8 0 2 4 9

//a 3 1 2 8 4 9 7 5 6 0
//  0 1 2 3 4 5 6 7 8 9

func distance(a, b []int) int {
	if len(a) != len(b) {
		log.Fatalln("Array dimensions disagree")
	}

	// number of swaps to sort b as the order of a

	ainv := make([]int, len(a))

	for i, num := range a {
		ainv[num] = i
	}

	bnew := make([]int, len(b))

	for i, num := range b {
		bnew[i] = ainv[num]
	}

	return inversions.Count(bnew)
}

func permutation(n int) []int {
	a := make([]int, n)

	for i := 0; i < n; i++ {
		a[i] = i
	}

	util.ShuffleIntSlice(a)
	return a
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
	}

	a := permutation(n)
	b := permutation(n)

	for i := 0; i < n; i++ {
		fmt.Printf("%d %d\n",  a[i], b[i])
	}

	fmt.Printf("inversions = %d", distance(a, b))
}
