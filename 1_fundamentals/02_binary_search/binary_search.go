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

func IndexOf(a []int, key int) int {
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
	sort.Ints(nums)

	for {
		var key int
		_, err = fmt.Scan(&key)
		if err != nil {
			log.Fatal(err)
		}
		if IndexOf(nums, key) == -1 {
			fmt.Println(key)
		}
	}
}
