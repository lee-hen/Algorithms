package tst

import (
	"github.com/lee-hen/Algorithms/util"
	"log"
	"strings"
)

// Proposition J. The number of links in a TST built from N string keys of average length w is between 3N and 3Nw. Proof.
// Immediate, by the same argument as for PROPOSITION I.

// Proposition K. A search miss in a TST built from N random string keys requires ~ln N character compares, on the average.
// A search hit or an insertion in a TST uses ~ln N character compare for each character in the search key.
// Proof: The search hit/insertion cost is immediate from the code. The search miss cost is a consequence of the same arguments discussed in the proof sketch of PROPOSITION H.
// We assume that all but a constant number of the nodes on the search path (a few at the top) act as random BSTs on R character values with average path length ln R, so we multiply the time cost logR N = ln N / ln R by ln R.

// Proposition L. A search or an insertion in a TST built from N random string keys with no external one-way branching and Rt-way branching at the root requires roughly ln N − t ln R character compares, on the average.
// Proof: These rough estimates follow from the same argument we used to prove PROPOSITION K. We assume that all but a constant number of the nodes on the search path (a few at the top) act as random BSTs on R character values, so we multiply the time cost by ln R.

type TST struct {
	root *Node // root of TST
	n int // size
}

type data interface {}

type Node struct {
	c byte    // character
	left, mid , right *Node // left, middle, and right subtries
	value data // value associated with string
}

func newNode() *Node {
	return &Node{}
}

// Size
// Returns the number of key-value pairs in this symbol table.
func (t *TST) Size() int {
	return t.n
}

// Contains
// Does this symbol table contain the given key?
// return true if this symbol table contains key and
// return false otherwise
func (t *TST) Contains(key string) bool {
	if key == "" {
		log.Fatalln("argument to contains() is null")
	}

	_, found := t.Get(key)
	return found
}

// Get
// Returns the value associated with the given key.
func (t *TST) Get(key string) (int, bool){
	if len(key) == 0 {
		log.Fatalln("key must have length >= 1")
	}

	x := get(t.root, key, 0)
	if x == nil {
		return 0, false
	}

	return x.value.(int), true
}

// return subtrie corresponding to given key
func get(x *Node, key string, d int) *Node {
	if x == nil {
		return nil
	}

	c := key[d]
	if c < x.c {
		return get(x.left, key, d)
	} else if c > x.c {
		return get(x.right, key, d)
	} else if d < len(key)-1 {
		return get(x.mid, key, d+1)
	}

	return x
}

// Put
// Inserts the key-value pair into the symbol table, overwriting the old value
// with the new value if the key is already in the symbol table.
func (t *TST) Put(key string, value int) {
	if key == "" {
		log.Fatalln("calls put() with null key")
	}
	if !t.Contains(key) {
		t.n++
	}

	t.root = put(t.root, key, value, 0)
}

func put(x *Node, key string, value, d int) *Node {
	c := key[d]
	if x == nil {
		x = newNode()
		x.c = c
	}

	if c < x.c {
		x.left = put(x.left, key, value, d)
	} else if c > x.c {
		x.right = put(x.right, key, value, d)
	} else if d < len(key)-1 {
		x.mid = put(x.mid, key, value, d+1)
	} else {
		x.value = value
	}

	return x
}

// LongestPrefixOf
// Returns the string in the symbol table that is the longest prefix of query,
// or "", if no such string.
func (t *TST) LongestPrefixOf(query string) string {
	if query == "" {
		log.Fatalln("calls longestPrefixOf() with null argument")
	}

	if len(query) == 0 {
		return ""
	}

	var length, i int
	x := t.root
	for x != nil && i < len(query) {
		c := query[i]
		if c < x.c {
			x = x.left
		} else if c > x.c {
			x = x.right
		} else {
			i++
			if x.value != nil {
				length = i
			}
			x = x.mid
		}
	}

	return query[0:length]
}

// Keys
// Returns all keys in the symbol table as an Iterable.
func (t *TST) Keys() []string {
	results := make([]string, 0)
	collect(t.root, &strings.Builder{}, &results)

	return results
}

// KeysWithPrefix
// Returns all of the keys in the set that start with prefix
func (t *TST) KeysWithPrefix(prefix string) []string {
	if prefix == "" {
		log.Fatalln("calls keysWithPrefix() with null argument")
	}

	results := make([]string, 0)
	x := get(t.root, prefix, 0)
	if x == nil {
		return results
	}
	if x.value != nil {
		results = append(results, prefix)
	}

	stringBuilder := strings.Builder{}
	stringBuilder.WriteString(prefix)

	collect(x.mid, &stringBuilder, &results)
	util.ReverseStringSlice(results)
	return results
}

func collect(x *Node, prefix *strings.Builder, results *[]string) {
	if x == nil {
		return
	}

	collect(x.left, prefix, results)

	if x.value != nil {
		*results = append(*results, prefix.String() + string(x.c))
	}
	prefix.WriteByte(x.c)

	collect(x.mid, prefix, results)


	// delete last char
	temp := prefix.String()
	prefix.Reset()
	prefix.WriteString(temp[:len(temp)-1])

	collect(x.right, prefix, results)
}

// KeysThatMatch
// Returns all of the keys in the symbol table that match pattern,
// where the character '.' is interpreted as a wildcard character.
func (t *TST) KeysThatMatch(pattern string) []string {
	results := make([]string, 0)
	collectMatches(t.root, &strings.Builder{}, 0, pattern, &results)
	util.ReverseStringSlice(results)
	return results
}

func collectMatches(x *Node, prefix *strings.Builder, i int, pattern string, results *[]string) {
	if x == nil {
		return
	}

	c := pattern[i]

	if c == '.' || c < x.c {
		collectMatches(x.left, prefix, i, pattern, results)
	}
	if c == '.' || c == x.c {
		if i == len(pattern)-1 && x.value != nil {
			*results = append(*results, prefix.String() + string(x.c))
		}
		if i < len(pattern)-1 {
			prefix.WriteByte(x.c)
			collectMatches(x.mid, prefix, i+1, pattern, results)

			// delete last char
			temp := prefix.String()
			prefix.Reset()
			prefix.WriteString(temp[:len(temp)-1])
		}
	}

	if c == '.' || c > x.c {
		collectMatches(x.right, prefix, i, pattern, results)
	}
}
