package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const R = 256 // extend ASCII alphabet size

// Proposition B. LSD string sort stably sorts fixed-length strings.
// Proof: This fact depends crucially on the key-indexed counting implementation being stable, as indicated in PROPOSITION A. After sorting keys on their i trailing characters (in a stable manner),
// we know that any two keys appear in proper order in the array (considering just those characters) either because the first of their i trailing characters is different, in which case the sort on that character puts them in order, or because the first of their ith trailing characters is the same, in which case they are in order because of stability (and by induction, for i-1).
// Proposition B. LSD string sort uses ~7WN + 3WR array accesses and extra space proportional to N + R to sort N items whose keys are W-character strings taken from an R-character alphabet.
// Proof: The method is W passes of key-indexed counting, except that the aux[] array is initialized just once. The total is immediate from the code and PROPOSITION A.

func Sort(a []string, w int) {
	aux := make([]string, len(a), len(a))

	for d := w-1; d >= 0; d-- { // sort by key-indexed counting on dth character

		// compute frequency counts
		count := make([]int, R+1, R+1)
		for i := 0; i < len(a); i++ {
			count[a[i][d]+1]++
		}

		// compute cumulates
		for r := 0; r < R; r++ {
			count[r+1] += count[r]
		}

		//l := R
		//for i := R; i >= 0; i-- {
		//	l -= count[i]
		//	count[i] = l
		//}

		// move data
		for i := 0; i < len(a); i++ {
			aux[count[a[i][d]]] = a[i]
			count[a[i][d]]++
		}

		// copy back
		for i := 0; i < len(a); i++ {
			a[i] = aux[i]
		}
	}
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
	}

	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err == io.EOF {
		return
	}
	if err != nil {
		log.Fatal(err)
	}
	str := strings.Split(line[:len(line)-1], " ")
	Sort(str, n)
	for _, s := range str {
		fmt.Println(s)
	}
}
