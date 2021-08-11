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
	util.ShuffleIntSlice(a)

	sort(0, len(a)-1, a)
}

func sort(lo, hi int, a []int) {
	if hi <= lo {
		return
	}

	pv := a[lo]
	lt, mid, gt := lo, lo+1, hi

	for mid <= gt {
		if a[mid] < pv {
			exchange(a, mid, lt)
			mid++
			lt++
		} else if a[mid] > pv {
			exchange(a, mid, gt)
			gt--
		} else {
			mid++
		}
	}

	sort(lo, lt-1, a)
	sort(gt+1, hi, a)
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

	Sort(nums)
	fmt.Println(nums)
}
