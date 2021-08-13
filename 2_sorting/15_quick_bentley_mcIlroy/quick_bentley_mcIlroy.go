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

const MEDIAN_OF_3_CUTOFF = 40

func Sort(a []int) {
	sort(0, len(a)-1, a)
}

func sort(lo, hi int, a []int) {
	n := partitioningElement(lo, hi, a)
	if n <= 1 {
		return
	}

	i, j := lo, hi+1
	p, q := lo, hi+1
	v := a[lo]

	for  {
		for i = i+1; i < hi && util.Less(a[i], v); i++ {}
		for j = j-1; j > lo && util.Less(v, a[j]); j-- {}

		if i == j && util.Eq(a[i], v) {
			p++
			exchange(a, p, i)
		}
		if i >= j {
			break
		}
		exchange(a, i, j)

		if util.Eq(a[i], v) {
			p++
			exchange(a, p, i)
		}

		if util.Eq(a[j], v) {
			q--
			exchange(a, q, j)
		}
	}

	i = j+1
	for k := lo; k <= p; k++ {
		exchange(a, k, j)
		j--
	}

	for k := hi; k >= q; k-- {
		exchange(a, k, i)
		i++
	}

	sort(lo, j, a)
	sort(i, hi, a)
}

func exchange(a []int, i, j int) {
	a[i], a[j] = a[j], a[i]
}

func partitioningElement(lo, hi int, a []int) int {
	n := hi - lo + 1
	var mid int

	if n <= MEDIAN_OF_3_CUTOFF {
		mid = median3(lo, lo + n/2, hi, a)
		exchange(a, mid, lo)
	} else {
		eps := n/8
		mid = lo + n/2
		m1 := median3(lo, lo + eps, lo + eps + eps, a)
		m2 := median3(mid - eps, mid, mid + eps, a)
		m3 := median3(hi - eps - eps, hi - eps, hi, a)
		ninther := median3(m1, m2, m3, a)
		exchange(a, ninther, lo)
	}

	return n
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
