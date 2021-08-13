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

const R = 256

func Sort(a []string) {
	sort(0, len(a)-1, a)
}

func sort(lo, hi int, a []string) {
	st := make(util.Stack, 0)
	first := make([]int, R+2, R+2)
	next := make([]int, R+2, R+2)
	var d int

	st.Push(lo)
	st.Push(hi)
	st.Push(d)

	for !st.IsEmpty() {
		d = st.Pop()
		hi = st.Pop()
		lo = st.Pop()

		if hi <= lo {
			continue
		}

		// compute frequency counts
		for i := lo; i <= hi; i++ {
			c := util.CharAt(a[i], d) + 1  // account for -1 representing end-of-string
			first[c+1]++
		}

		// first[c] = location of first string whose dth character = c
		first[0] = lo
		for c := 0; c <= R; c++ {
			first[c+1] += first[c]

			if c > 0 && first[c+1]-1 > first[c] {
				// add subproblem for character c (excludes sentinel c == 0)
				st.Push(first[c])
				st.Push(first[c+1] - 1)
				st.Push(d + 1)
			}
		}

		// next[c] = location to place next string whose dth character = c
		for c := 0; c < R+2; c++ {
			next[c] = first[c]
		}

		// permute data in place
		for k := lo; k <= hi; k++ {
			c := util.CharAt(a[k], d) + 1
			for first[c] > k {
				exchange(a, k, next[c])
				next[c]++

				c = util.CharAt(a[k], d) + 1
			}
			next[c]++
		}

		// clear first[] and next[] arrays
		for c := 0; c < R+2; c++ {
			first[c] = 0
			next[c] = 0
		}
	}
}

func exchange(a []string, i, j int) {
	a[i], a[j] = a[j], a[i]
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
