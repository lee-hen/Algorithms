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

type intSlice []int

func (s intSlice) less (i, j int) bool {
	return s[i] < s[j]
}

func merge(src, dst intSlice, lo, mid, hi int) {
	i, j := lo, mid+1
	for k := lo; k <= hi; k++ {
		if i > mid {
			dst[k] = src[j]
			j++
		} else if j > hi {
			dst[k] = src[i]
			i++
		} else if src.less(j, i) {
			dst[k] = src[j]
			j++
		} else {
			dst[k] = src[i]
			i++
		}
	}
}

func mergeSort(src, dst intSlice, lo, hi int) {
	if hi <= lo {
		return
	}

	mid := lo + (hi-lo)/2
	mergeSort(dst, src, lo, mid)
	mergeSort(dst, src, mid+1, hi)
	merge(src, dst, lo, mid, hi)
}

// Sort
// Proposition G. Top-down mergesort uses at most 6N lgN array accesses to sort an array of length N.
// Proof: Each merge uses at most 6N array accesses (2N for the copy, 2N for the move back, and at most 2N for compares). The result follows from the same argument as for PROPOSITION F.
func Sort(a []int) {
	aux := make([]int, len(a), len(a))
	copy(aux, a)
	mergeSort(aux, a, 0, len(a)-1)
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
