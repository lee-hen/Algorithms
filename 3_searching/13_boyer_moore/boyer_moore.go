package main

import (
	"fmt"
	"github.com/lee-hen/Algorithms/util"
)

//  % go run boyer_moore.go
//  abracadabra
//  abacadabrabracabracadabrabrabracad
//  text:    abacadabrabracabracadabrabrabracad
//  pattern:               abracadabra
//
//  % go run boyer_moore.go
//  rab
//  abacadabrabracabracadabrabrabracad
//  text:    abacadabrabracabracadabrabrabracad
//  pattern:         rab
//
//  % go run boyer_moore.go
//  bcara
//  abacadabrabracabracadabrabrabracad
//  text:    abacadabrabracabracadabrabrabracad
//  pattern:                                   bcara
//
//  % go run boyer_moore.go
//  rabrabracad
//  abacadabrabracabracadabrabrabracad
//  text:    abacadabrabracabracadabrabrabracad
//  pattern:                        rabrabracad
//
//  % go run boyer_moore.go
//  abacad
//  abacadabrabracabracadabrabrabracad
//  text:    abacadabrabracabracadabrabrabracad
//  pattern: abacad

// Property O. On typical inputs, substring search with the Boyer-Moore mismatched character heuristic uses ~N/M character compares to search for a pattern of length M in a text of length N.
// Discussion: This result can be proved for various random string models, but such models tend to be unrealistic, so we shall skip the details.
// In many practical situations it is true that all but a few of the alphabet characters appear nowhere in the pattern, so nearly all compares lead to M characters being skipped, which gives the stated result.

const R = 256   // the radix
var right []int // the bad-character skip array

var pat string     // store the pattern as a string

// BoyerMoore
//  Preprocesses the pattern string.
func BoyerMoore(pattern string) {
	pat = pattern

	// position of rightmost occurrence of c in the pattern
	right = make([]int, R, R)
	for c := 0; c < R; c++ {
		right[c] = -1
	}

	for j := 0; j < len(pat); j++ {
		right[pat[j]] = j
	}
}

// Search
// Returns the index of the first occurrrence of the pattern string
// in the text string.
func Search(txt string) int {
	m, n := len(pat), len(txt)
	var skip int
	for i := 0; i <= n-m; i += skip {
		skip = 0
		for j := m-1; j >= 0; j-- {
			if pat[j] != txt[i+j] {
				skip = util.Max(1, j-right[txt[i+j]])
				break
			}
		}
		if skip == 0 {
			return i
		}
	}

	return n
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

	BoyerMoore(pattern)
	offset := Search(txt)

	fmt.Println("text:   ", txt)
	fmt.Print("pattern: ")
	for i := 0; i < offset; i++ {
		fmt.Print(" ")
	}
	fmt.Println(pattern)
}
