package main

import (
	"github.com/lee-hen/Algorithms/util"

	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func Sort(a []int) {
	sort(0, len(a)-1, a)
}

func sort(lo, hi int, a []int) {
	if lo >= hi {
		return
	}

	pivot := partition(lo, hi, a)
	sort(lo, pivot-1, a)
	sort(pivot+1, hi, a)
}

func partition(lo, hi int, a []int) int {
	n := hi-lo+1

	mid := median3(lo, lo + n/2, hi, a)
	exchange(a, mid, lo)

	i, j := lo, hi+1
	v := a[lo]

	for i = i+1; util.Less(a[i], v); i++ {
		if i == hi {
			exchange(a, lo, hi)
			return hi
		}
	}

	for j = j-1; util.Less(v, a[j]); j-- {
		if j == lo+1 {
			return lo
		}
	}

	for i < j {
		exchange(a, i, j)

		for i = i+1; util.Less(a[i], v); i++ {}
		for j = j-1; util.Less(v, a[j]); j-- {}
	}
	exchange(a, lo, j)

	return j
}

func exchange(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

// return the index of the median element among a[i], a[j], and a[k]
func median3(i, j, k int, a []int) int {
	if util.Less(a[i], a[j]) {
		if util.Less(a[j], a[k]) {
			return j
		} else if util.Less(a[i], a[k]) {
			return k
		} else {
			return i
		}
	} else {
		if util.Less(a[k], a[j]) {
			return j
		} else if util.Less(a[k], a[i]) {
			return k
		} else {
			return i
		}
	}
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
