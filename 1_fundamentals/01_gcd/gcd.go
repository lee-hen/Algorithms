package main

import "fmt"

func main() {
	fmt.Println(gcd(888,54))
}

// gcd(888, 54)
// 888%54=24
// 54%24=6
// 6%0=6
// Euclid’s algorithm
func gcd(p, q int) int {
	if q == 0 {
		return p
	}
	return gcd(q, p%q)
}
