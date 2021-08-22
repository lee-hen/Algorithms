package util

import (
	"hash/crc32"
	"log"
	"math/rand"
	"time"
)

func Max(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr < num {
			curr = num
		}
	}
	return curr
}

func Min(arg int, rest ...int) int {
	curr := arg
	for _, num := range rest {
		if curr > num {
			curr = num
		}
	}
	return curr
}

func Abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Less(v, w int) bool {
	return v < w
}

func Greater(v, w int) bool {
	return v > w
}

func Eq(v, w int) bool {
	return v == w
}

// CharAt
// return dth character of s, -1 if d = length of string
func CharAt(s string, d int) int {
	if d == len(s) {
		return -1
	}

	return int(s[d])
}

// ByteAt
// return dth byte of b, -1 if d = length of []byte
func ByteAt(b []byte, d int) int {
	if d == len(b) {
		return -1
	}

	return int(b[d])
}

// Gcd
// gcd(888, 54)
// 888%54=24
// 54%24=6
// 6%0=6
// Euclid’s algorithm
func Gcd(p, q int) int {
	if q == 0 {
		return p
	}
	return Gcd(q, p%q)
}

type Interface interface {
	Swap(i, j int)
	Len() int
}

func ShuffleStringSlice(a []string) {
	ShuffleSlice(StringSlice(a))
}

func ShuffleIntSlice(a []int) {
	ShuffleSlice(IntSlice(a))
}

func ShuffleSlice(slice Interface) {
	n := slice.Len()
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < n; i++ {
		r := rand.Intn(n)
		slice.Swap(i, r)
	}
}

func ReverseStringSlice(a []string) {
	ReverseSlice(StringSlice(a))
}

func ReverseIntSlice(a []int) {
	ReverseSlice(IntSlice(a))
}

func ReverseSlice(slice Interface) Interface {
	for i, j := 0, slice.Len()-1; i < j; i, j = i+1, j-1 {
		slice.Swap(i, j)
	}
	return slice
}

type IntSlice []int

func (p IntSlice) Len() int           { return len(p) }
func (p IntSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type StringSlice []string

func (p StringSlice) Len() int           { return len(p) }
func (p StringSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }

type Stack []int

func (stack *Stack) Push(arr int) {
	*stack = append(*stack, arr)
}

func (stack *Stack) Pop() int {
	s := *stack
	last := s[len(s)-1]
	*stack = s[:len(s)-1]
	return last
}

func (stack *Stack) Peek() int {
	s := *stack
	return s[len(s)-1]
}

func (stack *Stack) Size() int {
	return len(*stack)
}

func (stack *Stack) IsEmpty() bool {
	return len(*stack) == 0
}

// String hashes a string to a unique hashcode.
// crc32 returns a uint32, but for our use we need
// a non negative integer. Here we cast to an integer
// and invert it if the result is negative.
func String(s string) int {
	v := int(crc32.ChecksumIEEE([]byte(s)))
	if v >= 0 {
		return v
	}
	if -v >= 0 {
		return -v
	}
	return 0
}

// Bernoulli
// Returns a random boolean from a Bernoulli distribution with success
// probability p
func Bernoulli(p float64) bool {
	if !(p >= 0.0 && p <= 1.0) {
		log.Fatalln("probability p must be between 0.0 and 1.0: ", p)
	}
	rand.Seed(time.Now().UnixNano())
	return rand.Float64() < p
}
