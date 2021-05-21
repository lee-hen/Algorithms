package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func PrintAll(a []int) {
	n := len(a)
	sort.Ints(a)
	if containsDuplicates(a) {
		log.Fatal("array contains duplicate integers")
	}

	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			k := binarySearch(a, -(a[i] + a[j]))
			if k > j {
				fmt.Println(a[i], a[j], a[k])
			}
		}
	}
}

func Count(a []int) int {
	n := len(a)
	sort.Ints(a)
	if containsDuplicates(a) {
		log.Fatal("array contains duplicate integers")
	}
	var count int
	for i := 0; i < n; i++ {
		for j := i+1; j < n; j++ {
			k := binarySearch(a, -(a[i] + a[j]))
			if k > j {
				count++
			}
		}
	}
	return count
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

	fmt.Println(Count(nums))
	PrintAll(nums)
}


func binarySearch(a []int, key int) int {
	lo := 0
	hi := len(a)-1
	for lo <= hi {
		mid := lo + (hi-lo) / 2
		if key < a[mid] {
			hi = mid-1
		} else if key > a[mid]{
			lo = mid+1
		} else {
			return mid
		}
	}
	return -1
}

func containsDuplicates(a []int) bool {
	for i := 1; i < len(a); i++ {
		if a[i] == a[i-1] {
			return true
		}
	}
	return false
}
