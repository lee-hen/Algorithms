package main

import "github.com/lee-hen/Algorithms/util"

const R = 256

func Sort(a []string) {
	sort(0, len(a)-1, a)
}

func sort(lo, hi int, a []string) {
	st := make(util.Stack, 0)
	count := make([]int, R+2, R+2)
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
			count[c]++
		}

		// accumulate counts relative to a[0], so that
		// count[c] is the number of keys <= c
		count[0] += lo
		for c := 0; c <= R; c++ {
			count[c+1] += count[c]

			if c > 0 && count[c+1]-1 > count[c] {
				// add subproblem for character c (excludes sentinel c == 0)
				st.Push(count[c])
				st.Push(count[c+1] - 1)
				st.Push(d + 1)
			}
		}

		// permute data in place
		// for details and proof see Knuth Theorem 5.1.2B and ch 5.2 excercise 13.
		for r := hi; r >= lo; r-- {
			c := util.CharAt(a[r], d) + 1
			for r >= lo && count[c]-1 <= r {
				if count[c]-1 == r {
					count[c]--
				}
				r--
				if r >= lo {
					c = util.CharAt(a[r], d) + 1
				}
			}

			// if r < lo the subarray is sorted.
			if r < lo {
				break
			}

			// permute a[r] until correct element is in place
			for count[c] = count[c]-1; count[c] != r; count[c]-- {
				exchange(a, r, count[c])
				c = util.CharAt(a[r], d) + 1
			}

		}

		// clear count[] array
		for c := 0; c < R+2; c++ {
			count[c] = 0
		}
	}
}

func exchange(a []string, i, j int) {
	a[i], a[j] = a[j], a[i]
}
