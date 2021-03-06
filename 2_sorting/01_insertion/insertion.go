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

// Proposition B. Insertion sort uses ~n2/4 compares and ~n2/4 exchanges to sort a randomly ordered array of length n
// with distinct keys, on the average. The worst case is ~n2/2 compares and ~n2/2 exchanges and the best case is n − 1
// compares and 0 exchanges.
// Proof: Just as for PROPOSITION A, the number of compares and exchanges is easy to visualize in the n-by-n diagram
// that we use to illustrate the sort. We count entries below the diagonal—all of them, in the worst case,
// and none of them, in the best case. For randomly ordered arrays, we expect each item to go about halfway back,
// on the average, so we count one-half of the entries below the diagonal.
// The number of compares is the number of exchanges plus an additional term equal to n minus the number of times
// the item inserted is the smallest so far. This term is between 0 (array in reverse order) and n − 1 (array in order).

// Proposition C. The number of exchanges used by insertion sort is equal to the number of inversions in the array,
//and the number of compares is at least equal to the number of inversions and at most equal to the number of
// inversions plus the array size minus 1.
// Proof: Every exchange involves two inverted adjacent entries and thus reduces the number of inversions by one,
// and the array is sorted when the number of inversions reaches zero. Every exchange corresponds to a compare,
// and an additional compare might happen for each value of i from 1 to N-1
// (when a[i] does not reach the left end of the array).

// Property D. The running times of insertion sort and selection sort are quadratic and within a small constant
// factor of one another for randomly ordered arrays of distinct values.
// Evidence: This statement has been validated on many different computers over the past half-century.
// Insertion sort was about twice as fast as selection sort when the first edition of this book was written in 1980 and
// it still is today, even though it took several hours to sort 100,000 items with these algorithms then and
// just several seconds today. Is insertion sort a bit faster than selection sort on your computer?
// To find out, you can use the class SortCompare on the next page, which uses the sort() methods in the classes named
// as command-line arguments to perform the given number of experiments (sorting arrays of the given size) and prints
// the ratio of the observed running times of the algorithms.

func Sort(a []int, less func (i, j int) bool) {
	n := len(a)
	for i := 1; i < n; i++ {
		for j := i; j > 0 && less(j, j-1); j-- {
			exchange(a, j, j-1)
		}
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
