package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

func Shuffle(a []string) {
	n := len(a)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		r := rand.Intn(n)
		a[i], a[r] = a[r], a[i]
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
	a := strings.Split(line[:len(line)-1], "")
	Shuffle(a)

	for i := 0; i < len(a); i++ {
		fmt.Println(a[i])
	}
}
