package main

import (
	"bufio"
	"fmt"
	"github.com/lee-hen/Algorithms/util"
	"io"
	"log"
	"os"
	"strings"
)

const CUTOFF = 15

// Sort
// Proposition E. To sort an array of N random strings, 3-way string quicksort uses ~2Nln N character compares, on the average.
// Proof: There are two instructive ways to understand this result.
// First, considering the method to be equivalent to quicksort partitioning on the leading character,
// then (recursively) using the same method on the subarrays, we should not be surprised that the total number of operations is about the same as for normal quicksort—but they are single-character compares, not full-key compares.
// Second, considering the method as replacing key-indexed counting by quicksort, we expect that the N logR N running time from PROPOSITION D should be multiplied by a factor of 2 ln R because it takes quicksort 2R ln R steps to sort R characters, as opposed to R steps for the same characters in the MSD string sort. We omit the full proof.
func Sort(a []string) {
	util.ReverseStringSlice(a)
	sort(0, len(a)-1, 0, a)
}

// 3-way string quicksort a[lo..hi] starting at dth character
func sort(lo, hi, d int, a []string) {
	if hi <= lo + CUTOFF {
		insertion(a, lo, hi, func(v, w string)  bool {
			for i := d; i < util.Min(len(v), len(w)); i++ {
				if util.CharAt(v, i) < util.CharAt(w, i) {
					return true
				}
				if util.CharAt(v, i) > util.CharAt(w, i) {
					return false
				}
			}
			return len(v) < len(w)
		})
	 return
	}

	lt, gt := lo, hi
	v := util.CharAt(a[lo], d)
	mid := lo + 1
	for mid <= gt {
		t := util.CharAt(a[mid], d)
		if t < v {
			exchange(a, lt, mid)
			lt++
			mid++
		} else if t > v {
			exchange(a, mid, gt)
			gt--
		} else {
			mid++
		}
	}

	// a[lo..lt-1] < v = a[lt..gt] < a[gt+1..hi].
	sort(lo, lt-1, d, a)
	if v >= 0 {
		sort(lt, gt, d+1, a)
	}
	sort(gt+1, hi, d, a)
}

func insertion(a []string, lo, hi int, less func (v, w string) bool) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && less(a[j], a[j-1]); j-- {
			exchange(a, j, j-1)
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
