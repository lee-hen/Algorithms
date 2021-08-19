package main

import (
	"fmt"
)

//  % go run kmp.go
//  abracadabra
//  abacadabrabracabracadabrabrabracad
//  text:    abacadabrabracabracadabrabrabracad
//  pattern:               abracadabra
//
//  % go run kmp.go
//  rab
//  abacadabrabracabracadabrabrabracad
//  text:    abacadabrabracabracadabrabrabracad
//  pattern:         rab
//
//  % go run kmp.go
//  bcara
//  abacadabrabracabracadabrabrabracad
//  text:    abacadabrabracabracadabrabrabracad
//  pattern:                                   bcara
//
//  % go run kmp.go
//  rabrabracad
//  abacadabrabracabracadabrabrabracad
//  text:    abacadabrabracabracadabrabrabracad
//  pattern:                        rabrabracad
//
//  % go run kmp.go
//  abacad
//  abacadabrabracabracadabrabrabracad
//  text:    abacadabrabracabracadabrabrabracad
//  pattern: abacad

// Proposition N. Knuth-Morris-Pratt substring search accesses no more than M + N characters to search for a pattern of length M in a text of length N.
// Proof. Immediate from the code: we access each pattern character once when computing dfa[][] and each text character once (in the worst case) in search().

const R = 256 // the radix
var m int // length of pattern
var dfa [][]int  // the KMP automoton

// KMP
// Preprocesses the pattern string.
func KMP(pat string) {
	m = len(pat)

	// build DFA from pattern
	dfa = make([][]int, R, R)
	for i := range dfa {
		dfa[i] = make([]int, m, m)
	}

	dfa[pat[0]][0] = 1
	for x, j := 0, 1; j < m; j++ {
		for c := 0; c < R; c++ {
			dfa[c][j] = dfa[c][x] // Copy mismatch cases.
		}

		dfa[pat[j]][j] = j+1 // Set match case.
		x = dfa[pat[j]][x]  // Update restart state.
	}
}

// Search
// Returns the index of the first occurrrence of the pattern string
// in the text string.
func Search(txt string) int {
	// simulate operation of DFA on text
	n := len(txt)
	var i, j int
	for ; i < n && j < m; i++ {
		j = dfa[txt[i]][j]
	}

	if j == m {
		return i-m  // found
	}

	return n // not found
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

	KMP(pattern)
	offset := Search(txt)

	fmt.Println("text:   ", txt)
	fmt.Print("pattern: ")
	for i := 0; i < offset; i++ {
		fmt.Print(" ")
	}
	fmt.Println(pattern)
}
