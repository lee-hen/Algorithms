package main

import (
	"github.com/lee-hen/Algorithms/util"

	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const R = 256 // extend ASCII alphabet size

// Sort
// Proposition C. To sort N random strings from an R-character alphabet, MSD string sort examines about N logR N characters, on average.
// Proof sketch: We expect the subarrays to be all about the same size, so the recurrence CN = RCN/R + N approximately describes the performance,
// which leads to the stated result, generalizing our argument for quicksort in CHAPTER 2.
// Again, this description of the situation is not entirely accurate, because N/R is not necessarily an integer, and the subarrays are the same size only on the average (and because the number of characters in real keys is finite).
// These effects turn out to be less significant for MSD string sort than for standard quicksort, so the leading term of the running time is the solution to this recurrence.
// The detailed analysis that proves this fact is a classical example in the analysis of algorithms, first done by Knuth in the early 1970s.
// Proposition D. MSD string sort uses between 8N + 3R and ~7wN + 3WR array accesses to sort N strings taken from an R-character alphabet, where w is the average string length.
// Proof: Immediate from the code, PROPOSITION A, and PROPOSITION B. In the best case MSD sort uses just one pass; in the worst case, it performs like LSD string sort.
// Proposition D. To sort N strings taken from an R-character alphabet, the amount of space needed by MSD string sort is proportional to R times the length of the longest string (plus N), in the worst case.
// Proof: The count[] array must be created within sort(), so the total amount of space needed is proportional to R times the depth of recursion (plus N for the auxiliary array). Precisely, the depth of the recursion is the length of the longest string that is a prefix of two or more of the strings to be sorted.
func Sort(a []string) {
	aux := make([]string, len(a), len(a))

	sort(0, len(a)-1, 0, a, aux)
}

// sort from a[lo] to a[hi], starting at the dth character
func sort(lo, hi, d int, a, aux []string) {
	if hi <= lo {
		return
	}

	// compute frequency counts
	count := make([]int, R+2, R+2)
	for i := lo; i <= hi; i++ {
		c := util.CharAt(a[i], d)
		count[c+2]++
	}

	// compute cumulates
	for r := 0; r < R+1; r++ {
		count[r+1] += count[r]
	}

	// distribute
	for i := lo; i <= hi; i++ {
		c := util.CharAt(a[i], d)
		aux[count[c+1]] = a[i]
		count[c+1]++
	}

	// copy back
	for i := lo; i <= hi; i++ {
		a[i] = aux[i - lo]
	}

	// recursively sort for each character (excludes sentinel -1)
	for r := 0; r < R; r++ {
		// lo+count[r], lo+count[r+1]-1 group by same char and then sort each
		sort(lo+count[r], lo+count[r+1]-1, d+1, a, aux)
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err == io.EOF {
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Split(line[:len(line)-1], " ")

	Sort(str)
	for _, s := range str {
		fmt.Println(s)
	}
}
