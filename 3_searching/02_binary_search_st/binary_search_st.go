package binary_search_st

import (
	"fmt"
	"github.com/lee-hen/Algorithms/util"
	"log"
)

// Proposition A. Search misses and insertions in an (unordered) linked-list symbol table having n key-value pairs both require n compares,
// and search hits require n compares in the worst case. Proof: When searching for a key that is not in the list, we test every key in the table against the search key.
// Because of our policy of disallowing duplicate keys, we need to do such a search before each insertion.

// Corollary. Inserting n distinct keys into an initially empty linked-list symbol table uses ~n2/2 compares.

// Proposition B. Binary search in an ordered array with N keys uses no more than lg N + 1 compares for a search (successful or unsuccessful).
// Proof: This analysis is similar to (but simpler than) the analysis of mergesort (PROPOSITION F in CHAPTER 2).
// Let C(N) be the number of compares to search for a key in a symbol table of size N. We have C(0) = 0, C(1) = 1,
// and for N > 0 we can write a recurrence relationship that directly mirrors the recursive method: C(N) ≤ C(N/2) + 1
// Whether the search goes to the left or to the right, the size of the subarray is no more than N/2, and we use one compare to check for equality
// and to choose whether to go left or right. When N is one less than a power of 2 (say N = 2n−1), this recurrence is not difficult to solve.
// First, since N/2 = 2n−1−1, we have C(2n−1) ≤ C(2n−1−1) + 1 Applying the same equation to the first term on the right,
// we have C(2n−1) ≤ C(2n−2−1) + 1 + 1 Repeating the previous step n − 2 additional times gives C(2n−1) ≤ C(20) + n which
// leaves us with the solution C(N) = C(2n) ≤ n + 1 < lg N + 1 Exact solutions for general N are more complicated,
// but it is not difficult to extend this argument to establish the stated property for all values of N (see EXERCISE 3.1.20).
// With binary search, we achieve a logarithmic-time search guarantee.

// Proposition B (continued). Inserting a new key-value pair into an ordered symbol table of size n uses ~ 4n array accesses in the worst case,
// so inserting n keys-value pairs into an initially empty table uses ~ 2n2 array accesses in the worst case. Proof: Same as for PROPOSITION A.

type BinarySearchST struct {
	Keys []string
	Values []int
}

func NewBinarySearchST(capacity int) *BinarySearchST {
	binarySearchSt := BinarySearchST{
		Keys: make([]string, 0, capacity),
		Values: make([]int, 0, capacity),
	}

	return &binarySearchSt
}

// Size
// Returns the number of key-value pairs in this symbol table.
func (st *BinarySearchST) Size() int {
	return len(st.Keys)
}

// IsEmpty
// Returns true if this symbol table is empty.
func (st *BinarySearchST) IsEmpty() bool {
	return st.Size() == 0
}

// Contains
// Does this symbol table contain the given key?
// return true if this symbol table contains key and
// return false otherwise
func (st *BinarySearchST) Contains(key string) bool {
	if key == "" {
		log.Fatalln("first argument to contains() is null")
	}

	_, found := st.Get(key)
	return found
}

// Get
// Returns the value associated with the given key in this symbol table.
// the value associated with the given key if the key is in the symbol table
// return 0, false if the key is not in the symbol table
func (st *BinarySearchST) Get(key string) (int, bool) {
	if key == "" {
		log.Fatalln("argument to get() is null")
	}

	if st.IsEmpty() {
		return 0, false
	}

	i := st.Rank(key)

	if i < st.Size() && st.Keys[i] == key {
		return st.Values[i], true
	}

	return 0, false
}

// Rank
// Returns the number of keys in this symbol table strictly less than key.
// the number of keys in the symbol table strictly less than key
func (st *BinarySearchST) Rank(key string) int {
	if key == "" {
		log.Fatalln("argument to rank() is null")
	}

	lo, hi := 0, st.Size()-1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if key < st.Keys[mid] {
			hi = mid-1
		} else if key > st.Keys[mid] {
			lo = mid+1
		} else {
			return mid
		}
	}

	return lo
}

// Put
// Inserts the specified key-value pair into the symbol table, overwriting the old
// value with the new value if the symbol table already contains the specified key.
func (st *BinarySearchST) Put(key string, value int) {
	if key == "" {
		log.Fatalln("first argument to put() is null")
	}

	i, n := st.Rank(key), st.Size()

	if i < n && st.Keys[i] == key {
		st.Values[i] = value
		return
	}

	st.Keys = append(st.Keys, key)
	st.Values = append(st.Values, value)

	for j := n; j > i; j-- {
		st.Keys[j] = st.Keys[j-1]
		st.Values[j] = st.Values[j-1]
	}

	st.Keys[i] = key
	st.Values[i] = value
}

// Delete
// Removes the specified key and associated value from this symbol table
// (if the key is in the symbol table).
func (st *BinarySearchST) Delete(key string) {
	if key == "" {
		log.Fatalln("argument to delete() is null")
	}

	if st.IsEmpty() {
		return
	}

	i, n := st.Rank(key), st.Size()

	if i == n || st.Keys[i] != key {
		return
	}

	for j := i; j < n-1; j++ {
		st.Keys[j] = st.Keys[j+1]
		st.Values[j] = st.Values[j+1]
	}

	st.Keys = st.Keys[:n-1]
	st.Values = st.Values[:n-1]
}

// DeleteMin
// Removes the smallest key and associated value from this symbol table.
func (st *BinarySearchST) DeleteMin() {
	if st.IsEmpty() {
		log.Fatalln("Symbol table underflow error")
	}
	st.Delete(st.Min())
}

// DeleteMax
// Removes the largest key and associated value from this symbol table.
func (st *BinarySearchST) DeleteMax() {
	if st.IsEmpty() {
		log.Fatalln("Symbol table underflow error")
	}
	st.Delete(st.Max())
}

// Max
// Returns the largest key in this symbol table.
func (st *BinarySearchST) Max() string {
	if st.IsEmpty() {
		log.Fatalln("called max() with empty symbol table")
	}
	return st.Keys[st.Size()-1]
}

// Min
// Returns the smallest key in this symbol table.
func (st *BinarySearchST) Min() string {
	if st.IsEmpty() {
		log.Fatalln("called min() with empty symbol table")
	}
	return st.Keys[0]
}

// Select
// Return the kth smallest key in this symbol table.
func  (st *BinarySearchST) Select(k int) string {
	if k < 0 || k >= st.Size() {
		log.Fatalf("called select() with invalid argument: : %d\n", k)
	}

	return st.Keys[k]
}

// Floor
// the largest key in this symbol table less than or equal to key
func (st *BinarySearchST) Floor(key string) string {
	if key == "" {
		log.Fatalln("argument to floor() is null")
	}

	i := st.Rank(key)
	if i < st.Size() && key == st.Keys[i] {
		return st.Keys[i]
	}

	if i == 0 {
		fmt.Printf("argument: %s to floor() is too small\n", key)
		return ""
	}

	return st.Keys[i-1]
}

// Ceiling
// Returns the smallest key in this symbol table greater than or equal to key
func (st *BinarySearchST) Ceiling(key string) string {
	if key == "" {
		log.Fatalln("argument to ceiling() is null")
	}

	i := st.Rank(key)
	if i == st.Size() {
		fmt.Printf("argument: %s to ceiling() is too large\n", key)
		return ""
	}

	return st.Keys[i]
}

// SizeBetween
// Returns the number of keys in this symbol table in the specified range.
func (st *BinarySearchST) SizeBetween(lo, hi string) int {
	if lo == "" {
		log.Fatalln("first argument to size() is null")
	}
	if hi == "" {
		log.Fatalln("second argument to size() is null")
	}

	if lo > hi {
		return 0
	}
	if st.Contains(hi) {
		return st.Rank(hi) - st.Rank(lo) + 1
	}

	return st.Rank(hi) - st.Rank(lo)
}

// KeysBetween
// Returns all keys in this symbol table in the given range.
func (st *BinarySearchST) KeysBetween(lo, hi string)  []string {
	if lo == "" {
		log.Fatalln("first argument to keys() is null")
	}
	if hi == "" {
		log.Fatalln("second argument to keys() is null")
	}

	keys := make([]string, 0)
	if lo > hi {
		return keys
	}

	for i := st.Rank(lo); i < st.Rank(hi); i++ {
		keys = append(keys, st.Keys[i])
	}

	if st.Contains(hi) {
		keys = append(keys, st.Keys[st.Rank(hi)])
	}

	util.ReverseStringSlice(keys)

	return keys
}

func IsSorted(st *BinarySearchST) bool {
	for i := 1; i < st.Size(); i++ {
		if st.Keys[i] < st.Keys[i-1] {
			return false
		}
	}

	return true
}

func RankCheck(st *BinarySearchST) bool {
	for i := 0; i < st.Size(); i++ {
		if i != st.Rank(st.Select(i)) {
			return false
		}
	}

	for i := 0; i < st.Size(); i++ {
		if st.Keys[i] != st.Select(st.Rank(st.Keys[i])) {
			return false
		}
	}

	return true
}
