package suffix_array

import (
	"github.com/lee-hen/Algorithms/util"
	"sort"
)

//   i ind lcp rnk  select
//  ---------------------------
//   0  11   -   0  "!"
//   1  10   0   1  "A!"
//   2   7   1   2  "ABRA!"
//   3   0   4   3  "ABRACADABRA!"
//   4   3   1   4  "ACADABRA!"
//   5   5   1   5  "ADABRA!"
//   6   8   0   6  "BRA!"
//   7   1   3   7  "BRACADABRA!"
//   8   4   0   8  "CADABRA!"
//   9   6   0   9  "DABRA!"
//  10   9   0  10  "RA!"
//  11   2   2  11  "RACADABRA!"

// SuffixArray
// Initializes a suffix array for the given text string.
func SuffixArray(text string) *SuffixSlice {
	n := len(text)
	suffixes := make(SuffixSlice, n, n)
	for i := 0; i < n; i++ {
		suffixes[i] = Suffix{text, i}
	}

	sort.Sort(suffixes)
	return &suffixes
}

type Suffix struct {
	text string
	index int
}

func (s *Suffix) Len() int {
	return len(s.text) - s.index
}

func (s *Suffix) CharAt(i int) byte {
	return s.text[s.index + i]
}

func (s *Suffix) CompareTo(t *Suffix) int {
	if s == t {
		return 0
	}

	n := util.Min(s.Len(), t.Len())
	for i := 0; i < n; i++ {
		if s.CharAt(i) < t.CharAt(i) {
			return -1
		}
		if s.CharAt(i) > t.CharAt(i) {
			return 1
		}
	}
	return s.Len() - t.Len()
}

func (s *Suffix) String() string {
	return s.text[s.index:]
}

type SuffixSlice []Suffix

func (s SuffixSlice) Len() int           { return len(s) }
func (s SuffixSlice) Less(i, j int) bool { return s[i].CompareTo(&s[j]) < 0 }
func (s SuffixSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// Index
// Returns the index into the original string of the ith smallest suffix.
// That is, text.substring(sa.index(i)) is the ith smallest suffix.
func (s SuffixSlice) Index(i int) int     { return s[i].index }

// Lcp
// Returns the length of the longest common prefix of the ith
// smallest suffix and the i-1st smallest suffix.
func (s SuffixSlice) Lcp(i int) int     { return lcpSuffix(s[i], s[i-1]) }

// longest common prefix of s and t
func lcpSuffix(s, t Suffix) int {
	n := util.Min(s.Len(), t.Len())
	for i := 0; i < n; i++ {
		if s.CharAt(i) != t.CharAt(i) {
			return i
		}
	}
	return n
}

// Select
// Returns the ith smallest suffix as a string.
func (s SuffixSlice) Select(i int) string     { return s[i].String() }

// Rank
// Returns the number of suffixes strictly less than the query string.
func (s SuffixSlice) Rank(query string) int {
	lo, hi := 0, s.Len()-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		cmp := compare(query, s[mid])
		if cmp < 0 {
			hi = mid-1
		} else if cmp > 0 {
			lo = mid+1
		} else {
			return mid
		}
	}

	return lo
}

// compare query string to suffix
func compare(query string, suffix Suffix) int {
	n := util.Min(len(query), suffix.Len())
	for i := 0; i < n; i++ {
		if query[i] < suffix.CharAt(i) {
			return -1
		}
		if query[i] > suffix.CharAt(i) {
			return 1
		}
	}
	return len(query) - suffix.Len()
}
