package main

import (
	bst "github.com/lee-hen/Algorithms/3_searching/03_BST"

	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
)

// Perfect balance. Write a program that inserts a set of keys into an initially empty BST such that the tree produced is equivalent to binary search,
// in the sense that the sequence of compares done in the search for any key in the BST is the same as the sequence of compares used by binary search for the same set of keys.

// Read sequence of strings from standard input (no duplicates),
// and insert into a BST so that BST is perfectly balanced.
// % go run perfect_balance.go
// input: P E R F C T B I N A R Y S R H
// print: N E B A C H F I R R P R T S Y

type PerfectBST struct {
	*bst.BST
}

func (perfectBst *PerfectBST) Perfect(a []string) {
	sort.Strings(a)

	perfectBst.perfect(a, 0, len(a)-1)
}

func (perfectBst *PerfectBST) perfect(a []string, lo, hi int) {
	if hi < lo {
		return
	}
	mid := lo + (hi - lo) /2
	perfectBst.Put(a[mid], mid)
	fmt.Printf("%s: %d\n", a[mid], mid)
	perfectBst.perfect(a, lo, mid-1)
	perfectBst.perfect(a, mid+1, hi)
}

func NewPerfectBST() *PerfectBST{
	return &PerfectBST{bst.NewBST()}
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

	perfectBst := NewPerfectBST()
	perfectBst.Perfect(str)

	fmt.Println("---------------------------")
	for _, s := range perfectBst.LevelOrder() {
		_, val := perfectBst.Get(s)
		fmt.Printf("%s  %d\n", s, val)
	}
}
