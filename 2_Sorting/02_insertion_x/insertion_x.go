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

// 2.1.23 Deck sort. Ask a few friends to sort a deck of cards (see EXERCISE 2.1.13).
// Observe them carefully and write down the method(s) that they use.

// 2.1.24 Insertion sort with sentinel. Develop an implementation of insertion sort that eliminates the j>0 test
// in the inner loop by first putting the smallest item into position.
// Use SortCompare to evaluate the effectiveness of doing so.
// Note: It is often possible to avoid an index-out-of-bounds test in this way—the element that enables the test to be
// eliminated is known as a sentinel.

// Sort
// Sorts a sequence of strings from standard input using an optimized
// version of insertion sort that uses half exchanges instead of
// full exchanges to reduce data movement..
func Sort(a []int, less func (i, j int) bool) {
	n := len(a)
	exchanges := 0

	for i := n-1; i > 0; i-- {
		if less(i, i-1) {
			exchange(a, i, i-1)
			exchanges++
		}
	}
	if exchanges == 0 {
		return
	}

	// 10 9 8 7 6 5 4
	//[4 10 9 8 7 6 5]  v 9
	//[4 9 10 8 7 6 5]  v 8 a[j]=8   a[j-1]= 10
	//[4 9 10 10 7 6 5] v 8 a[j]=10  a[j-1] 9
	//[4 8 9 10 7 6 5]  a[j] = v -> 8
	// ...
	// insertion sort with half-exchanges
	for i := 2; i < n; i++ {
		v := a[i]
		j := i
		for v < a[j-1] {
			// fmt.Println(a)
			a[j] = a[j-1]
			j--
		}
		a[j] = v
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

