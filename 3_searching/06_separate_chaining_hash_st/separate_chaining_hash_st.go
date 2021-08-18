package separate_chaining_hash_st

import (
	bst "github.com/lee-hen/Algorithms/3_searching/05_red_black_bst"
	"github.com/lee-hen/Algorithms/util"

	"log"
)

const initCapacity = 4

type HashST struct {
	n int  // number of key-value pairs
	st []bst.RedBlackBST // array of Red Black BST symbol tables
}

// NewHashST
// Initializes an empty symbol table
func NewHashST() *HashST {
	return &HashST{
		st: make([]bst.RedBlackBST, initCapacity, initCapacity),
	}
}

// newHashST
// Initializes an empty symbol table with m chains.
func newHashST(m int) *HashST {
	return &HashST{
		st: make([]bst.RedBlackBST, m, m),
	}
}

// hash function for keys - returns value between 0 and m-1

func (h *HashST) hash(key string) int {
	hashCode := util.String(key)
	return hashCode * 0x7fffffff % len(h.st)
}

// resize
// the hash table to have the given number of chains,
// rehashing all of the keys
func (h *HashST) resize(chains int) {
	temp := newHashST(chains)
	for i := 0; i < len(h.st); i++ {
		for _, key := range h.st[i].Keys() {
			val, _ := h.st[i].Get(key)
			temp.Put(key, val)
		}
	}
	h.st = temp.st
	h.n = temp.n
}

// Size
// Returns the number of key-value pairs in this symbol table.
func (h *HashST) Size() int {
	return h.n
}

// IsEmpty
// Returns true if this symbol table is empty
func (h *HashST) IsEmpty() bool {
	return h.Size() == 0
}

// Contains
// Returns true if this symbol table contains the specified key.
func (h *HashST) Contains(key string) bool {
	if key == "" {
		log.Fatalln("argument to contains() is null")
	}

	_,  found := h.Get(key)
	return found
}

// Get
// Returns the value associated with the specified key in this symbol table.
func (h *HashST) Get(key string) (int, bool) {
	if key == "" {
		log.Fatalln("argument to get() is null")
	}

	i := h.hash(key)
	return h.st[i].Get(key)
}

// Put
// Inserts the specified key-value pair into the symbol table, overwriting the old
// value with the new value if the symbol table already contains the specified key.
// Deletes the specified key (and its associated value) from this symbol table
// if the specified value is null
func (h *HashST) Put(key string, value int) {
	if key == "" {
		log.Fatalln("first argument to put() is null")
	}

	// double table size if average length of list >= 10
	if h.n >= 10 * len(h.st) {
		h.resize(2 * len(h.st))
	}

	i := h.hash(key)
	if !h.st[i].Contains(key) {
		h.n++
	}
	h.st[i].Put(key, value)
}

// Delete
// Removes the specified key and its associated value from this symbol table
// (if the key is in this symbol table).
func (h *HashST) Delete(key string) {
	if key == "" {
		log.Fatalln("argument to delete() is null")
	}

	i := h.hash(key)
	if h.st[i].Contains(key) {
		h.n--
	}
	h.st[i].Del(key)

	// halve table size if average length of list <= 2
	if len(h.st) > initCapacity && h.n <= 2*len(h.st) {
		h.resize(len(h.st)/2)
	}
}

// Keys
// return keys in symbol table as an Iterable
func (h *HashST) Keys() []string {
	keys := make([]string, 0)
	for i := 0; i < len(h.st); i++ {
		for _, key := range h.st[i].Keys() {
			keys = append(keys, key)
		}
	}

	util.ReverseStringSlice(keys)
	return keys
}
