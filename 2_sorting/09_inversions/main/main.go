package main

import (
	"github.com/lee-hen/Algorithms/2_sorting/09_inversions"

	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

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

	inversions := inversions.Count(nums)
	fmt.Println(nums, inversions)
}
