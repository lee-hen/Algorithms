package main

import (
	"crypto/rand"
	"fmt"
)

//  % go run rabin_karp.go
//  abracadabra
//  abacadabrabracabracadabrabrabracad
//  text:    abacadabrabracabracadabrabrabracad
//  pattern:               abracadabra
//
//  % go run rabin_karp.go
//  rab
//  abacadabrabracabracadabrabrabracad
//  text:    abacadabrabracabracadabrabrabracad
//  pattern:         rab
//
//  % go run rabin_karp.go
//  bcara
//  abacadabrabracabracadabrabrabracad
//  text:    abacadabrabracabracadabrabrabracad
//  pattern:                                   bcara
//
//  % go run rabin_karp.go
//  rabrabracad
//  abacadabrabracabracadabrabrabracad
//  text:    abacadabrabracabracadabrabrabracad
//  pattern:                        rabrabracad
//
//  % go run rabin_karp.go
//  abacad
//  abacadabrabracabracadabrabrabracad
//  text:    abacadabrabracabracadabrabrabracad
//  pattern: abacad

// Property P. The Monte Carlo version of Rabin-Karp substring search is linear-time and extremely likely to be correct,
// and the Las Vegas version of Rabin-Karp substring search is correct and extremely likely to be linear-time. Discussion:
// The use of the very large value of Q, made possible by the fact that we need not maintain an actual hash table,
// makes it extremely unlikely that a collision will occur. Rabin and Karp showed that when Q is properly chosen,
// we get a hash collision for random strings with probability 1/Q, which implies that, for practical values of the variables,
// there are no hash matches when there are no substring matches and only one hash match if there is a substring match.
// Theoretically, a text position could lead to a hash collision and not a substring match, but in practice it can be relied upon to find a match.

const R = 256 // radix

var pat string     // the pattern  // needed only for Las Vegas
var patHash int64  // pattern hash value
var m int          // pattern length
var q int64        // a large prime, small enough to avoid long overflow
var RM int64       //  R^(M-1) % Q

// RabinKarp
// Preprocesses the pattern string.
func RabinKarp(pattern string) {
	pat = pattern // save pattern (needed only for Las Vegas)
	m = len(pat)
	q = longRandomPrime()

	// precompute R^(m-1) % q for use in removing leading digit
	RM = 1
	for i := 1; i <= m-1; i++ {
		RM = (R * RM) % q
	}

	patHash = hash(pat, m)
}

// Compute hash for key[0..m-1].
func hash(key string, m int) int64 {
	var h int64

	for j := 0; j < m; j++ {
		h = (int64(R) * h +  int64(key[j])) % q
	}

	return h
}

// Las Vegas version: does pat[] match txt[i..i-m+1] ?
func check(txt string, i int) bool {
	for j := 0; j < m; j++ {
		if pat[j] != txt[i+j] {
			return false
		}
	}

	return true
}

// Search
// Returns the index of the first occurrrence of the pattern string
// in the text string.
func Search(txt string) int {
	n := len(txt)
	if n < m {
		return n
	}

	txtHash := hash(txt, m)

	// check for match at offset 0
	if patHash == txtHash && check(txt, 0) {
		return 0
	}

	// check for hash match; if hash match, check for exact match
	for i := m; i < n; i++ {
		// Remove leading digit, add trailing digit, check for match.
		txtHash = (txtHash + q - RM * int64(txt[i-m]) % q) % q
		txtHash = (txtHash * R + int64(txt[i])) % q

		// match
		offset := i - m + 1
		if patHash == txtHash && check(txt, offset) {
			return offset
		}
	}

	// no match
	return n
}

// a random 31-bit prime
func longRandomPrime() int64 {
	p, _ := rand.Prime(rand.Reader, 31)
	return p.Int64()
}

func main() {
	var pattern, txt string
	var err error

	_, err = fmt.Scan(&pattern)
	if err != nil {
		fmt.Println(err)
	}
	_, err = fmt.Scan(&txt)
	if err != nil {
		fmt.Println(err)
	}

	RabinKarp(pattern)
	offset := Search(txt)

	fmt.Println("text:   ", txt)
	fmt.Print("pattern: ")
	for i := 0; i < offset; i++ {
		fmt.Print(" ")
	}
	fmt.Println(pattern)
}
