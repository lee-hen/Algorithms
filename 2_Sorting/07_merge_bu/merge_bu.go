package main

import (
	"bufio"
	"fmt"
	"github.com/lee-hen/Algorithms/util"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// Bottom-up mergesort

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

func Sort(a []int) {
	aux := make(intSlice, len(a), len(a))
	for ln := 1; ln < len(a); ln *= 2 {
		for lo := 0; lo < len(a)-ln; lo += ln+ln {
			mid := lo + ln -1
			hi := util.Min(lo+ln+ln-1, len(a)-1)

			merge(a,aux, lo, mid, hi)
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
