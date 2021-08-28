package trie_set

import (
	"log"
	"strings"

	"github.com/lee-hen/Algorithms/util"
)

// An set for extended ASCII strings, implemented  using a 256-way trie.

const R = 256 // extended ASCII

type TrieSET struct {
	root *Node // root of trie
	n    int   // number of keys in trie
}

// Node
// R-way trie node
type Node struct {
	next     map[byte]*Node
	isString bool
}

func newNode() *Node {
	return &Node{
		next: make(map[byte]*Node),
	}
}

// Contains
// Does this symbol table contain the given key?
// return true if this symbol table contains key and
// return false otherwise
func (t *TrieSET) Contains(key string) bool {
	if key == "" {
		log.Fatalln("argument to contains() is null")
	}

	x := get(t.root, key, 0)

	if x == nil {
		return false
	}

	return x.isString
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

// Add
// Adds the key to the set if it is not already present.
func (t *TrieSET) Add(key string) {
	if key == "" {
		log.Fatalln("argument to add() is null")
	}
	t.root = add(t.root, key, 0, &t.n)
}

func add(x *Node, key string, d int, n *int) *Node {
	if x == nil {
		x = newNode()
	}

	if d == len(key) {
		if !x.isString {
			*n++
		}
		x.isString = true
		return x
	}

	c := key[d]
	x.next[c] = add(x.next[c], key, d+1, n)
	return x
}

// Size
// Returns the number of key-value pairs in the set
func (t *TrieSET) Size() int {
	return t.n
}

// IsEmpty
// Returns true if this set is empty, false otherwise.
func (t *TrieSET) IsEmpty() bool {
	return t.Size() == 0
}

// Keys
// Returns all keys in the symbol table as an Iterable.
func (t *TrieSET) Keys() []string {
	return t.KeysWithPrefix("")
}

// KeysWithPrefix
// Returns all of the keys in the set that start with prefix.
func (t *TrieSET) KeysWithPrefix(prefix string) []string {
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

	if x.isString {
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
// Returns all of the keys in the set that match pattern,
// where the character '.' is interpreted as a wildcard character.
func (t *TrieSET) KeysThatMatch(pattern string) []string {
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

	if d == len(pattern) && x.isString {
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
// Returns the string in the set that is the longest prefix of query,
// or "", if no such string.
func (t *TrieSET) LongestPrefixOf(query string) string {
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

	if x.isString {
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
func (t *TrieSET) Delete(key string) {
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
		if x.isString {
			*n--
		}

		x.isString = false
	} else {
		c := key[d]
		x.next[c] = del(x.next[c], key, d+1, n)
	}

	// remove subtrie rooted at x if it is completely empty
	if x.isString {
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
