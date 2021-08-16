package main

import (
	hash "github.com/lee-hen/Algorithms/3_searching/06_separate_chaining_hash_st"

	"fmt"
)

func main() {
	hashSt := hash.NewHashST()
	for i := 32; i <= 127; i++ {
		hashSt.Put(string(rune(i)), i)
	}
	fmt.Println(hashSt.Size())
	fmt.Println("----------------------------")
	for i := 32; i <= 127; i++ {
		key := string(rune(i))
		_, val := hashSt.Get(key)
		fmt.Println(key, val)
		hashSt.Delete(key)
	}
	fmt.Println("----------------------------")
	fmt.Println(hashSt.Size())
}
