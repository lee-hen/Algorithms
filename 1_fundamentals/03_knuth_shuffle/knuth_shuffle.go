package main

import (
	"github.com/lee-hen/Algorithms/util"

	"bufio"
	"fmt"
	"io"
	"log"
	"os"
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
	a := strings.Split(line[:len(line)-1], "")
	util.ShuffleStringSlice(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
