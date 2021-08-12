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

// Sort
// Proposition S. Heapsort uses fewer than 2N lg N + 2N compares (and half that many exchanges) to sort N items.
// Proof: The 2 N term covers the cost of heap construction (see PROPOSITION R). The 2 N lg N term follows from bounding the cost of each sink operation during the sortdown by 2lg N (see PROPOSITION PQ).
// Rearranges the array in ascending order, using the natural order.
func Sort(pq []int) []int {
	// heapify phase
	for k := len(pq)/2; k >= 1; k-- {
		sink(k, len(pq), pq)
	}

	// sortdown phase
	k := len(pq)
	for k > 1 {
		exchange(1, k, pq)
		k--
		sink(1, k, pq)
	}
	return pq
}

// Helper functions to restore the heap invariant.

func sink(k, n int, pq []int) {
	for 2*k <= n {
		j := 2*k
		if j < n && less(j, j+1, pq) {
			j++
		}
		if !less(k, j, pq) {
			break
		}
		exchange(k, j, pq)
		k = j
	}
}

// * Helper functions for comparisons and swaps.
// * Indices are "off-by-one" to support 1-based indexing.

func less(i, j int, pq []int) bool {
	return pq[i-1] < pq[j-1]
}

func exchange(i, j int, pq []int) {
	pq[i-1], pq[j-1] = pq[j-1], pq[i-1]
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

	Sort(nums)
	fmt.Println(nums)
}
