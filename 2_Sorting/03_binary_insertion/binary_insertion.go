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

func Sort(a []int) {
	n := len(a)

	// find lo which a[lo] > a[i]
	// binary search to determine index j at which to insert a[i]
	for i := 1; i < n; i++ {
		v := a[i]
		lo, hi := 0, i
		for lo < hi {
			mid := lo + (hi-lo)/2
			if v < a[mid] {
				hi = mid
			} else {
				lo = mid+1
			}
		}

		// insertion sort with "half exchanges"
		// (insert a[i] at index j and shift a[j], ..., a[i-1] to right)
		for j := i; j > lo; j-- {
			a[j] = a[j-1]
		}
		a[lo] = v
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
