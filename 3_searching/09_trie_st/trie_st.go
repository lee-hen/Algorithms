package trie_st

import (
	"github.com/lee-hen/Algorithms/util"
	"log"
	"strings"
)

// A string symbol table for extended ASCII strings, implemented
// using a 256-way trie.

// Proposition F. The linked structure (shape) of a trie is independent of the key insertion/deletion order: there is a unique trie for any given set of keys.
// Proof: Immediate, by induction on the subtries.

// Proposition G. The number of array accesses when searching in a trie or inserting a key into a trie is at most 1 plus the length of the key.
// Proof: Immediate from the code. The recursive get() and put() implementations carry an argument d that starts at 0, increments for each call, and is used to stop the recursion when it reaches the key length.

// Proposition H. The average number of nodes examined for search miss in a trie built from N random keys over an alphabet of size R is ~logR N.
// Proof sketch (for readers who are familiar with probabilistic analysis): The probability that each of the N keys in a random trie differs from a random search key in at least one of the leading t characters is (1 − R−t)N.
// Subtracting this quantity from 1 gives the probability that one of the keys in the trie matches the search key in all of the leading t characters. In other words, 1 − (1 − R−t)N is the probability that the search requires more than t character compares.
// From probabilistic analysis, the sum for t = 0, 1, 2, ... of the probabilities that an integer random variable is >t is the average value of that random variable, so the average search cost is 1 − (1 − R−1)N + 1 − (1 − R−2)N + ... + 1 − (1 − R−t)N + ... Using the elementary approximation (1−1/x)x ~ e−1,
// we find the search cost to be approximately (1 − e−N/R1) + (1 − e−N/R2) + ... + (1 − e−N/Rt) + ... The summand is extremely close to 1 for approximately lnR N terms with Rt substantially smaller than N; it is extremely close to 0 for all the terms with Rt substantially greater than N; and it is somewhere between 0 and 1 for the few terms with Rt ≈ N. So the grand total is about logR N.


// Proposition I. The number of links in a trie is between RN and RNw, where w is the average key length.
// Proof: Every key in the trie has a node containing its associated value that also has R links, so the number of links is at least RN.
// If the first characters of all the keys are different, then there is a node with R links for every key character, so the number of links is R times the total number of key characters, or RNw.

const R = 256 // extended ASCII

type TrieST struct {
	root *Node // root of trie
	n int // number of keys in trie
}

type data interface {}

// Node
// R-way trie node
type Node struct {
	next map[byte]*Node
	value data
}

func newNode() *Node {
	return &Node{
		next: make(map[byte]*Node),
		value: 0,
	}
}

// Get
// Returns the value associated with the given key.
func (t *TrieST) Get(key string) (int, bool){
	if key == "" {
		log.Fatalln("argument to get() is null")
	}

	x := get(t.root, key, 0)
	if x == nil {
		return 0, false
	}

	return x.value.(int), true
}

// Contains
// Does this symbol table contain the given key?
// return true if this symbol table contains key and
// return false otherwise
func (t *TrieST) Contains(key string) bool {
	if key == "" {
		log.Fatalln("first argument to contains() is null")
	}

	_, found := t.Get(key)
	return found
}

func get(x *Node, key string, d int) *Node {
	if x == nil {
		return nil
	}

	if d == len(key) {
		return x
	}

	c := key[d]
	return get(x.next[c], key, d+1)
}

// Put
// Inserts the key-value pair into the symbol table, overwriting the old value
// with the new value if the key is already in the symbol table.
func (t *TrieST) Put(key string, value int) {
	if key == "" {
		log.Fatalln("first argument to put() is null")
	}
	t.root = put(t.root, key, value, 0, &t.n)
}

func put(x *Node, key string, value, d int, n *int) *Node {
	if x == nil {
		x = newNode()
	}

	if d == len(key) {
		if x.value == nil {
			*n++
		}
		x.value = value
		return x
	}

	c := key[d]
	x.next[c] = put(x.next[c], key, value, d+1, n)
	return x
}

// Size
// Returns the number of key-value pairs in this symbol table.
func (t *TrieST) Size() int {
	return t.n
}

// IsEmpty
// Returns true if this symbol table is empty, false otherwise.
func (t *TrieST) IsEmpty() bool {
	return t.Size() == 0
}

// Keys
// Returns all keys in the symbol table as an Iterable.
func (t *TrieST) Keys() []string {
	return t.KeysWithPrefix("")
}

// KeysWithPrefix
// Returns all of the keys in the set that start with prefix
func (t *TrieST) KeysWithPrefix(prefix string) []string {
	results := make([]string, 0)
	x := get(t.root, prefix, 0)

	stringBuilder := strings.Builder{}
	stringBuilder.WriteString(prefix)

	collect(x, &stringBuilder, &results)
	util.ReverseStringSlice(results)
	return results
}

func collect(x *Node, prefix *strings.Builder, results *[]string) {
	if x == nil {
		return
	}

	if x.value != nil {
		*results = append(*results, prefix.String())
	}

	for c := 0; c < R; c++ {
		prefix.WriteByte(byte(c))
		collect(x.next[byte(c)], prefix, results)

		// delete last char
		temp := prefix.String()
		prefix.Reset()
		prefix.WriteString(temp[:len(temp)-1])
	}
}

// KeysThatMatch
// Returns all of the keys in the symbol table that match pattern,
// where the character '.' is interpreted as a wildcard character.
func (t *TrieST) KeysThatMatch(pattern string) []string {
	results := make([]string, 0)
	collectMatches(t.root, &strings.Builder{}, pattern, &results)
	util.ReverseStringSlice(results)
	return results
}

func collectMatches(x *Node, prefix *strings.Builder, pattern string, results *[]string) {
	if x == nil {
		return
	}

	d := prefix.Len()

	if d == len(pattern) && x.value != nil {
		*results = append(*results, prefix.String())
	}

	if d == len(pattern) {
		return
	}

	c := pattern[d]

	if c == '.' {
		for ch := 0; ch < R; ch++ {
			prefix.WriteByte(byte(ch))
			collectMatches(x.next[byte(ch)], prefix, pattern, results)

			// delete last char
			temp := prefix.String()
			prefix.Reset()
			prefix.WriteString(temp[:len(temp)-1])
		}

	} else {
		prefix.WriteByte(c)
		collectMatches(x.next[c], prefix, pattern, results)

		// delete last char
		temp := prefix.String()
		prefix.Reset()
		prefix.WriteString(temp[:len(temp)-1])
	}
}

// LongestPrefixOf
// Returns the string in the symbol table that is the longest prefix of {@code query},
// or "", if no such string.
func (t *TrieST) LongestPrefixOf(query string) string {
	if query == "" {
		log.Fatalln("argument to longestPrefixOf() is null")
	}

	length := longestPrefixOf(t.root, query, 0, -1)

	if length == -1 {
		return ""
	}

	return query[0:length]
}

// returns the length of the longest string key in the subtrie
// rooted at x that is a prefix of the query string,
// assuming the first d character match and we have already
// found a prefix match of given length (-1 if no such match)
func longestPrefixOf(x *Node, query string, d, length int) int {
	if x == nil {
		return length
	}

	if x.value != nil {
		length = d
	}

	if d == len(query) {
		return length
	}

	c := query[d]

	return longestPrefixOf(x.next[c], query, d+1, length)
}

// Delete
// Removes the key from the set if the key is present.
func (t *TrieST) Delete(key string) {
	if key == "" {
		log.Fatalln("argument to delete() is null")
	}

	t.root = del(t.root, key, 0, &t.n)
}

func del(x *Node, key string, d int, n *int) *Node {
	if x == nil {
		return nil
	}

	if d == len(key) {
		if x.value != nil {
			*n--
		}

		x.value = nil
	} else {
		c := key[d]
		x.next[c] = del(x.next[c], key, d+1, n)
	}

	// remove subtrie rooted at x if it is completely empty
	if x.value != nil {
		return x
	}

	for c := 0; c < R; c++ {
		if x.next[byte(c)] != nil {
			return x
		}

		delete(x.next, byte(c))
	}

	return nil
}
