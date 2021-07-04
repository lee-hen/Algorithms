package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// The study of the performance characteristics of ShellSort requires mathematical arguments that are beyond the scope of this book.
// If you want to be convinced, start by thinking about how you would prove the following fact: when an h-sorted array is k-sorted,
// it remains h-sorted. As for the performance of ALGORITHM 2.3,
// the most important result in the present context is the knowledge that the running time of ShellSort is not necessarily quadratic—for example,
// it is known that the worst-case number of compares for ALGORITHM 2.3 is proportional to N3/2.
// That such a simple modification can break the quadratic-running-time barrier is quite interesting,
// as doing so is a prime goal for many algorithm design problems.
// No mathematical results are available about the average-case number of compares for ShellSort for randomly ordered input.
// Increment sequences have been devised that drive the asymptotic growth of the worst-case number of compares down to N4/3, N5/4, N6/5, . . .,
// but many of these results are primarily of academic interest because these functions are hard to distinguish
// from one another (and from a constant factor of N) for practical values of N.

// Property E. The number of compares used by ShellSort with the increments 1, 4, 13, 40, 121, 364, . . .
// is bounded by a small multiple of N times the number of increments used.
// Evidence: Instrumenting ALGORITHM 2.3 to count compares and divide by the number of increments used is a straightforward exercise (see EXERCISE 2.1.12).
// Extensive experiments suggest that the average number of compares per increment might be N1/5, but it is quite difficult to discern the growth in that function unless N is huge.
// This property also seems to be rather insensitive to the input model.

func Sort(a []int, less func (i, j int) bool) {
	n := len(a)

	h := 1
	// 3x+1 increment sequence:  1, 4, 13, 40, 121, 364, 1093, ...
	for h < n/3 {
		h = 3*h+1
	}

	for h >= 1 {
		// h-sort the array
		for i := h; i < n; i++ {
			for j := i; j >= h && less(j, j-h); j-=h {
				exchange(a, j, j-h)
			}
		}

		h /= 3
	}

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

	Sort(nums, func(i, j int) bool{
		return nums[i] < nums[j]
	})
	fmt.Println(nums)
}
