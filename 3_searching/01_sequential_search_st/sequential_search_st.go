package sequential_search_st

// Proposition A. Search misses and insertions in an (unordered) linked-list symbol table having n key-value pairs both require n compares, and search hits require n compares in the worst case.
// Proof: When searching for a key that is not in the list, we test every key in the table against the search key. Because of our policy of disallowing duplicate keys, we need to do such a search before each insertion.
// Corollary. Inserting n distinct keys into an initially empty linked-list symbol table uses ~n2/2 compares.

import (
	"github.com/lee-hen/Algorithms/util"
	"log"
)

type SequentialSearchST struct {
	N int
	First *Node
}

type Node struct {
	Key string
	Value int
	Next *Node
}

func NewSequentialSearchST() *SequentialSearchST {
	return &SequentialSearchST{
		N :0,
	}
}

func NewNode(key string, value int) *Node {
	return &Node {
		Key: key,
		Value: value,
	}
}

func (st *SequentialSearchST) Size() int {
	return st.N
}

func (st *SequentialSearchST) IsEmpty() bool {
	return st.Size() == 0
}

func (st *SequentialSearchST) Contains(key string) bool {
	if key == "" {
		log.Fatalln("argument to contains() is null")
	}
	_, found := st.Get(key)
	return found
}

func (st *SequentialSearchST) Get(key string) (int, bool) {
	if key == "" {
		log.Fatalln("argument to get() is null")
	}

	for x := st.First; x != nil; x = x.Next {
		if key == x.Key {
			return  x.Value, true
		}
	}
	return 0, false
}

func (st *SequentialSearchST) Keys() []string {
	keys := make([]string, 0)
	for x := st.First; x != nil; x = x.Next {
		keys = append(keys, x.Key)
	}

	util.ReverseStringSlice(keys)
	return keys
}

func (st *SequentialSearchST) Put(key string, value int) {
	if key == "" {
		log.Fatalln("first argument to put() is null")
	}
	for x := st.First; x != nil; x = x.Next {
		if key == x.Key {
			x.Value = value
			return
		}
	}

	node := NewNode(key,value)
	node.Next = st.First
	st.First = node
	st.N++
}

func (st *SequentialSearchST) Delete(key string) {
	if key == "" {
		log.Fatalln("first argument to put() is null")
	}
	st.First = st.delete(st.First, key)
}

func (st *SequentialSearchST) delete(x *Node, key string) *Node {
	if x == nil {
		return nil
	}

	if key == x.Key {
		st.N--
		return x.Next
	}

	x.Next = st.delete(x.Next, key)
	return x
}
