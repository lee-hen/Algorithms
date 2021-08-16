package linear_probing_hash_st

import (
	"github.com/lee-hen/Algorithms/util"
	"log"
)

// must be a power of 2
const initCapacity = 4

type HashST struct {
	n int  // number of key-value pairs
	keys []string // the keys
	values []int // the values
}

// NewHashST
// Initializes an empty symbol table
func NewHashST() *HashST {
	return &HashST{
		keys: make([]string, initCapacity, initCapacity),
		values: make([]int, initCapacity, initCapacity),
	}
}

// newHashST
// Initializes an empty symbol table with the specified initial capacity
func newHashST(m int) *HashST {
	return &HashST{
		keys: make([]string, m, m),
		values: make([]int, m, m),
	}
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

	found, _ := h.Get(key)
	return found
}

// hash function for keys - returns value between 0 and m-1

func (h *HashST) hash(key string) int {
	hashCode := util.String(key)
	return hashCode * 0x7fffffff % len(h.keys)
}

// resize
// the hash table to the given capacity by re-hashing all of the keys
func (h *HashST) resize(capacity int) {
	temp := newHashST(capacity)
	for i := 0; i < len(h.keys); i++ {
		if key := h.keys[i]; key != "" {
			_, val := h.Get(key)
			temp.Put(key, val)
		}
	}
	h.keys = temp.keys
	h.values = temp.values
	h.n = temp.n
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

	// double table size if 50% full
	if h.n >= len(h.keys)/2 {
		h.resize(2 * len(h.keys))
	}

	var i int
	for i = h.hash(key); h.keys[i] != ""; i = (i+1) % len(h.keys) {
		if h.keys[i] == key {
			h.values[i] = value
			return
		}
	}

	h.keys[i] = key
	h.values[i] = value
	h.n++
}

// Get
// Returns the value associated with the specified key.
func (h *HashST) Get(key string) (bool, int) {
	if key == "" {
		log.Fatalln("argument to get() is null")
	}

	for i := h.hash(key); h.keys[i] != ""; i = (i+1) % len(h.keys) {
		if h.keys[i] == key {
			return true, h.values[i]
		}
	}

	return false, 0
}

// Delete
// Removes the specified key and its associated value from this symbol table
// (if the key is in this symbol table).
func (h *HashST) Delete(key string) {
	if key == "" {
		log.Fatalln("argument to delete() is null")
	}

	if !h.Contains(key) {
		return
	}

	// find position i of key
	i := h.hash(key)
	for key != h.keys[i] {
		i = (i+1) % len(h.keys)
	}

	// delete key and associated value
	h.keys[i] = ""
	h.values[i] = 0

	// rehash all keys in same cluster
	i = (i+1) % len(h.keys)
	for h.keys[i] != "" {
		// delete keys[i] an vals[i] and reinsert
		keyToRehash, valToRehash := h.keys[i], h.values[i]
		h.keys[i] = ""
		h.values[i] = 0
		h.n--
		h.Put(keyToRehash, valToRehash)
		i = (i+1) % len(h.keys)
	}

	// deleted
	h.n--

	// halves size of array if it's 12.5% full or less
	// if n > 0 && n == len(a)/4 resize(len(a)/2)
	if i > 0 && h.n <= len(h.keys)/8 {
		h.resize(len(h.keys)/2)
	}
}

// Keys
// Returns all keys in this symbol table as an Iterable.
// To iterate over all of the keys in the symbol table named st,
// use the foreach notation:
func (h *HashST) Keys() []string {
	keys := make([]string, 0)
	for _, key := range h.keys {
		if key != "" {
			keys = append(keys, key)
		}
	}

	util.ReverseStringSlice(keys)
	return keys
}

// Check
// integrity check - don't check after each put() because
// integrity not maintained during a delete()
func Check(h *HashST) bool {
	// check that hash table is at most 50% full
	if len(h.keys) < 2 * h.n {
		log.Fatalf("Hash table size = %d; array size n = %d\n", len(h.keys), h.n)
	}

	// check that each key in table can be found by get()
	for i := 0; i < len(h.keys); i++ {
		if h.keys[i] == "" {
			continue
		}
		_, val := h.Get(h.keys[i])
		if val != h.values[i] {
			log.Fatalf("get[%s] = %d; vals[i] = %d\n", h.keys[i], val, h.values[i])
		}
	}

	return true
}