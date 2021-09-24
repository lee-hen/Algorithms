package computer_science

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Sqrt
// Newton's method:
//    - initialize t = c
//    - replace t with the average of c/t and t
//    - repeat until desired accuracy reached
// 2
// 1.414213562373095
// ---
// 1000000
// 1000.0
// ---
// 0.4
// 0.6324555320336759
// ---
// 1048575
// 1023.9995117186336
// ---
// 16664444
// 4082.2106756021303
// ---
// 0
// 0.0
// ---
// 1e-50
// 9.999999999999999E-26
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
// 5
// 0 1
// 1 2
// 2 4
// 3 8
// 4 16
// 5 32
//
// 6
// 0 1
// 1 2
// 2 4
// 3 8
// 4 16
// 5 32
// 6 64
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

// Harmonic
// returns 1/1 + 1/2 + 1/3 + ... + 1/n
// 10
// 2.9289682539682538
// ---
// 10000
// 9.787606036044348
func Harmonic(n int) float64 {
	sum := 0.0
	for i := 1; i <= n; i++ {
		sum += 1.0 / float64(i)
	}

	return sum
}

// Binary
// Limitations
// Does not handle negative integers.
// 5    101
// 106  1101010
// 0    0
// 16   10000
func Binary(n int) string {
	power := 1

	// set power to the largest power of 2 that is <= n
	for power <= n/2 {
		power *= 2
	}

	var binary strings.Builder
	// check for presence of powers of 2 in n, from largest to smallest
	for power > 0 {
		// power is not present in n
		if n < power {
			binary.WriteByte(0)
		} else { // power is present in n, so subtract power from n
			binary.WriteByte(1)
			n -= power
		}

		// next smallest power of 2
		power /= 2
	}
	return binary.String()
}

// IntegerToBinary
// 8 1000
// 366 101101110
func IntegerToBinary(n int) string {
	if n == 0 {
		return ""
	}
	return IntegerToBinary(n/2) + strconv.Itoa(n % 2)
}
