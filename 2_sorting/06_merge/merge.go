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

// Top-down mergesort

type intSlice []int

func (s intSlice) less (i, j int) bool {
	return s[i] < s[j]
}

func merge(a, aux intSlice, lo, mid, hi int) {
	// copy to aux[]
	for k := lo; k <= hi; k++ {
		aux[k] = a[k]
	}

	// merge back to a[]
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			a[k] = aux[j]
			j++
		} else if j > hi {
			a[k] = aux[i]
			i++
		} else if aux.less(j, i) {
			a[k] = aux[j]
			j++
		} else {
			a[k] = aux[i]
			i++
		}
	}
}

func mergeSort(a, aux intSlice, lo, hi int) {
	if hi <= lo {
		return
	}

	mid := lo + (hi-lo)/2
	mergeSort(a, aux, lo, mid)
	mergeSort(a, aux, mid+1, hi)
	merge(a, aux, lo, mid, hi)
}

// Sort
// Proposition G. Top-down mergesort uses at most 6N lgN array accesses to sort an array of length N.
// Proof: Each merge uses at most 6N array accesses (2N for the copy, 2N for the move back, and at most 2N for compares). The result follows from the same argument as for PROPOSITION F.
func Sort(a []int) {
	aux := make(intSlice, len(a), len(a))
	mergeSort(a, aux, 0, len(a)-1)
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
