package main

import (
	"github.com/lee-hen/Algorithms/util"

	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// Sort
// Proposition K. Quicksort uses ~ 2N ln N compares (and one-sixth that many exchanges) on the average to sort an array of length N with distinct keys.
// Proof: Let CN be the average number of compares needed to sort N items with distinct values.
// We have C0 = C1 = 0 and for N > 1 we can write a recurrence relationship that directly mirrors the recursive program:
// CN = N + 1 + (C0 + C1 + . . . + CN−2 + CN−1) / N + (CN−1 + CN−2 + . . . + C0)/N
// The first term is the cost of partitioning (always N + 1),
// the second term is the average cost of sorting the left subarray (which is equally likely to be any size from 0 to N − 1),
// and the third term is the average cost for the right subarray (which is the same as for the left subarray).
// Multiplying by N and collecting terms transforms this equation to NCN = N(N + 1) + 2(C0 + C1+ . . . +CN−2+CN−1)
// Subtracting the same equation for N − 1 from this equation gives NCN − (N − 1)CN−1 = 2N + 2CN−1 Rearranging terms and dividing by N(N + 1) leaves CN/(N + 1) = CN−1/N + 2/(N + 1) which telescopes to give the result
// CN ~ 2 (N + 1)(1/3 + 1/4 + . . . + 1/(N + 1) )
// The parenthesized quantity is the discrete estimate of the area under the curve 2/x from 3 to N + 1 and CN ~ 2N lnN by integration. Note that 2N ln N ≈ 1.39N lg N,
// so the average number of compares is only about 39 percent higher than in the best case. A similar (but much more complicated) analysis is needed to establish the stated result for exchanges.
// Proposition L. Quicksort uses ~ N2/2 compares in the worst case, but random shuffling protects against this case.
// Proof: By the argument just given, the number of compares used when one of the subarrays is empty for every partition is
// N + (N − 1) + (N − 2) + . . . + 2 + 1 = (N + 1) N / 2
// This behavior means not only that the time required will be quadratic but also that the space required to handle the recursion will be linear,
// which is unacceptable for large arrays. But (with quite a bit more work)
// it is possible to extend the analysis that we did for the average to find that the standard deviation of the number of compares is about .65 N,
// so the running time tends to the average as N grows and is unlikely to be far from the average.
// For example, even the rough estimate provided by Chebyshev’s inequality says that the probability that the running time is more than ten times the average for an array with a million elements is less than .00001 (and the true probability is far smaller).
// The probability that the running time for a large array is close to quadratic is so remote that we can safely ignore the possibility (see EXERCISE 2.3.10).
// For example, the probability that quicksort will use as many compares as insertion sort or selection sort when sorting a large array on your computer is much less than the probability that your computer will be struck by lightning during the sort!
func Sort(a []int) {
	util.ShuffleIntSlice(a)

	sort(0, len(a)-1, a)
}

func Select(a []int, k int) int {
	util.ShuffleIntSlice(a)

	lo, hi := 0, len(a)-1

	for hi > lo {
		pivot := partition(lo, hi, a)
		if pivot > k {
			hi = pivot-1
		} else if pivot < k {
			lo = pivot+1
		} else {
			return a[pivot]
		}
	}
	return a[lo]
}

func sort(lo, hi int, a []int) {
	if lo >= hi {
		return
	}

	pivot := partition(lo, hi, a)
	sort(lo, pivot-1, a)
	sort(pivot+1, hi, a)
}

func partition(pivot, hi int, a []int) int {
	i, j := pivot+1, hi

	for j >= i {
		if a[i] > a[pivot] && a[j] < a[pivot] {
			a[i], a[j] = a[j], a[i]
		}

		if a[i] <= a[pivot] {
			i++
		}

		if a[j] >= a[pivot] {
			j--
		}
	}

	exchange(a, pivot, j)
	return j
}

func exchange(a []int, i, j int) {
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
	nums := make([]int, len(str))
	for i, s := range str {
		if nums[i], err = strconv.Atoi(s); err != nil {
			log.Fatal(err)
		}
	}

	rand.Seed(time.Now().UnixNano())
	k := rand.Intn(len(nums))

	fmt.Printf("quick select with random idx %d\n", k)
	fmt.Println(Select(nums, k))

	fmt.Println("----------------------------------------")

	Sort(nums)
	fmt.Println(nums)
}
