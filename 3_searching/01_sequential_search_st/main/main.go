package main

import (
	seqSearchSt "github.com/lee-hen/Algorithms/3_searching/01_sequential_search_st"

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
	str := strings.Split(line[:len(line)-1], " ")
    st := seqSearchSt.NewSequentialSearchST()

    for i, s := range str{
    	st.Put(s, i)
	}

	for _, s := range st.Keys() {
		val, _ := st.Get(s)
		fmt.Printf("%s %d\n", s, val)
	}

	fmt.Println(st.Size())
}
