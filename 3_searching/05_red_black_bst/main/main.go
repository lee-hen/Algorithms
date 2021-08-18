package main

import (
	bst "github.com/lee-hen/Algorithms/3_searching/05_red_black_bst"

	"fmt"
	"strings"
)

func main() {
	test := "S E A R C H E X A M P L E"
	keys := strings.Split(test, " ")

	n := len(keys)
	st := bst.NewRedBlackBST()

	for i := 0; i < n; i++ {
		st.Put(keys[i], i)
	}

	fmt.Printf("size = %d\n", st.Size())
	fmt.Printf("min = %s\n", st.Min())
	fmt.Printf("max = %s\n", st.Max())
	fmt.Println("--------------------------------")
	for _, s := range st.Keys() {
		val, _ := st.Get(s)
		fmt.Printf("%s  %d\n", s, val)
	}
	fmt.Println("")

	for i := 0; i < st.Size() / 2; i++ {
		st.DelMin()
	}
	fmt.Printf("After deleting the smallest %d keys\n",  st.Size() / 2)
	fmt.Println("--------------------------------")
	for _, s := range st.Keys() {
		val, _ := st.Get(s)
		fmt.Printf("%s  %d\n", s, val)
	}
	fmt.Println()

	for !st.IsEmpty() {
		st.Del(st.Select(st.Size() / 2))
	}
	fmt.Println("After deleting the remaining keys")
	fmt.Println("--------------------------------")
	for _, s := range st.Keys() {
		val, _ := st.Get(s)
		fmt.Printf("%s  %d\n", s, val)
	}
	fmt.Println()

	fmt.Println("After adding back N keys")
	fmt.Println("--------------------------------")
	for i := 0; i < len(keys); i++ {
		st.Put(keys[i], i)
	}
	for _, s := range st.Keys() {
		val, _ := st.Get(s)
		fmt.Printf("%s  %d\n", s, val)
	}
	fmt.Println()
}
