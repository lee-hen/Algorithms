package main

import (
	"fmt"
	"sort"
)

func main() {
	domains := []string{
		"com.google",
		"com.cnn",
		"edu.princeton.cs.www",
		"edu.princeton.cs.bolle",
		"edu.princeton.cs",
		"com.apple",
		"edu.princeton",
		"com.amazon",
		"edu.princeton.ee",
	}

	sort.Strings(domains)

	for i := 0; i < len(domains); i++ {
		fmt.Println(domains[i])
	}
}
