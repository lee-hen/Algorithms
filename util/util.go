package util

import (
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
