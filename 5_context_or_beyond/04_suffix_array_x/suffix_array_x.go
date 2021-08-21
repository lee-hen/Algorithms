package suffix_array_x

import "github.com/lee-hen/Algorithms/util"

const CUTOFF = 5

// SuffixArrayX
// Initializes a suffix array for the given text string.
func SuffixArrayX(str string) *SuffixSlice {
	text := []byte(str)
	index := make([]int, len(text), len(text))
	for i := 0; i < len(text); i++ {
		index[i] = i
	}

	sort(0, len(text)-1, 0, text, index)
	return &SuffixSlice {text, index }
}

func insertion(index []int, lo, hi int, less func (i, j int) bool) {
	for i := lo; i <= hi; i++ {
		for j := i; j > lo && less(index[j], index[j-1]); j-- {
			exchange(index, j, j-1)
		}
	}
}

// 3-way string quicksort lo..hi starting at dth character
func sort(lo, hi, d int, text []byte, index []int) {
	if hi <= lo + CUTOFF {
		insertion(index, lo, hi, func(i, j int)  bool {
			if i == j {
				return false
			}
			i = i+d
			j = j+d
			for i < len(text) && j < len(text) {
				if text[i] < text[j] {
					return true
				}
				if text[i] > text[j] {
					return false
				}
				i++
				j++
			}

			return i > j
		})

		return
	}

	lt, gt := lo, hi
	v := util.ByteAt(text, index[lo] + d)

	mid := lo + 1
	for mid <= gt {
		t := util.ByteAt(text, index[mid] + d)
		if t < v {
			exchange(index, lt, mid)
			lt++
			mid++
		} else if t > v {
			exchange(index, mid, gt)
			gt--
		} else {
			mid++
		}
	}

	// a[lo..lt-1] < v = a[lt..gt] < a[gt+1..hi].
	sort(lo, lt-1, d, text, index)
	if v > 0 {
		sort(lt, gt, d+1, text, index)
	}
	sort(gt+1, hi, d, text, index)
}

func exchange(index []int, i, j int) {
	index[i], index[j] = index[j], index[i]
}

type SuffixSlice struct {
	text []byte
	index []int
}

// Index
// Returns the index into the original string of the ith smallest suffix.
// That is, text.substring(sa.index(i)) is the ith smallest suffix.
func (s *SuffixSlice) Index(i int) int     { return s.index[i] }

// Lcp
// Returns the length of the longest common prefix of the ith
// smallest suffix and the i1st smallest suffix.
func (s *SuffixSlice) Lcp(i int) int     { return lcp(s.text, s.index[i], s.index[i-1]) }

// longest common prefix of text[i..n) and text[j..n)
func lcp(text []byte, i, j int) int {
	var n int
	for i < len(text) && j < len(text) {
		if text[i] != text[j] {
			return n
		}
		i++
		j++
		n++
	}
	return n
}

// Select
// Returns the ith smallest suffix as a string.
func (s *SuffixSlice) Select(i int) string     { return string(s.text[s.index[i]:])}

// Rank
// Returns the number of suffixes strictly less than the query string.
func (s *SuffixSlice) Rank(query string) int {
	lo, hi := 0, len(s.text)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		cmp := s.compare(query, s.index[mid])
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
func (s *SuffixSlice) compare(query string, i int) int {
	var j int
	for i < len(s.text) && j < len(query) {
		if util.CharAt(query, j) != util.CharAt(string(s.text), i) {
			return util.CharAt(query, j) - util.CharAt(string(s.text), i)
		}

		i++
		j++
	}
	if i < len(s.text) {
		return -1
	}

	if j < len(query) {
		return 1
	}

	return 0
}
