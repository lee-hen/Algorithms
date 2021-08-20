package main

import (
	avl "github.com/lee-hen/Algorithms/5_context_or_beyond/02_avl_tree_st"

	"fmt"
	"strings"
)

func main() {
	test := "S E A R C H E X A M P L E"
	keys := strings.Split(test, " ")

	n := len(keys)
	st := avl.NewAVLTree()

	for i := 0; i < n; i++ {
		st.Put(keys[i], i)
	}

	fmt.Printf("size = %d\n", st.Size())
	fmt.Printf("min = %s\n", st.Min())
	fmt.Printf("max = %s\n", st.Max())
	fmt.Println("")

	fmt.Println("Testing keys()")
	fmt.Println("--------------------------------")
	for _, s := range st.Keys() {
		val, _ := st.Get(s)
		fmt.Printf("%s  %d\n", s, val)
	}
	fmt.Println("")


	fmt.Println("Testing select")
	fmt.Println("--------------------------------")
	for i := 0; i < st.Size(); i++ {
		fmt.Printf("%d %s\n", i, st.Select(i))
	}
	fmt.Println("")


	fmt.Println("key rank floor ceil")
	fmt.Println("-------------------")
	for i := 'A'; i <= 'X'; i++ {
		s := string(i)
		fmt.Printf("%2s %4d %4s %4s\n", s, st.Rank(s), st.Floor(s), st.Ceiling(s))
	}
	fmt.Println("")

	from := []string{"A", "Z", "X", "0", "B", "C"}
	to := []string{"Z", "A", "X", "Z", "G", "L"}
	fmt.Println("range search")
	fmt.Println("-------------------")
	for i := 0; i < len(from); i++ {
		fmt.Printf("%s-%s (%2d) : ", from[i], to[i], st.SizeBetween(from[i], to[i]))
		for _, s := range st.KeysBetween(from[i], to[i]) {
			fmt.Printf(s + " ")
		}
		fmt.Println()
	}
	fmt.Println()

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
