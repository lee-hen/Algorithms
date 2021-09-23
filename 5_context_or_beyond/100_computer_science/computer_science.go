package computer_science

import (
	"fmt"
	"math"
)

// Sqrt
// Newton's method:
//    - initialize t = c
//    - replace t with the average of c/t and t
//    - repeat until desired accuracy reached
func Sqrt(c float64) float64 {
	epsilon := 1.0e-15 // relative error tolerance 0.0000000000000010
	t := c              // estimate of the square root of c

	// repeatedly apply Newton update step until desired precision is achieved
	for math.Abs(t- c/t) > epsilon * t {
		t = (c/t + t)/2.0
	}
	return t
}


// PowersOfTwo
// This program takes a command-line argument n and prints a table of
// the powers of 2 that are less than or equal to 2^n.
func PowersOfTwo(n int) int {
	i := 0 // count from 0 to N
	powerOfTwo := 1 // the ith power of two

	for i <= n {
		fmt.Println(i, powerOfTwo)
		powerOfTwo = 2 * powerOfTwo // double to get the next one
		i = i + 1
	}

	return powerOfTwo
}
