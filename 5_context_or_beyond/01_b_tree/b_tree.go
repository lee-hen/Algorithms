package b_tree

import (
	"fmt"
	"log"
	"strings"
)

const M = 4 // max children per B-tree node = M-1 (must be even and greater than 2)

type BTree struct {
	root *Node // root of the B-tree
	// height of the B-tree
	// number of key-value pairs in the B-tree
	height, n int
}

type Node struct {
	m int      // number of children
	children []*Entry // the array of children
}

type Data interface {}

type Entry struct {
	key string
	value Data
	next *Node  // helper field to iterate over array entries
}

func newNode(k int) *Node {
	return &Node{
		m: k,
		children: make([]*Entry, M),
	}
}

func newEntry(key string, value Data, next *Node) *Entry{
	return &Entry {
		key,
		value,
		next,
	}
}

// NewBTree
// Initializes an empty B-tree.
func NewBTree() *BTree {
	return &BTree {
		root: newNode(0),
	}
}

// IsEmpty
// Returns true if this symbol table is empty.
func (tree *BTree) IsEmpty() bool {
	return tree.Size() == 0
}

// Size
//Returns the number of key-value pairs in this symbol table.
func (tree *BTree) Size() int {
	return tree.n
}

// Height
// Returns the height of this B-tree (for debugging).
func (tree *BTree) Height() int {
	return tree.height
}

// Get
// Returns the value associated with the given key.
func (tree *BTree) Get(key string) Data {
	if key == "" {
		log.Fatalln("argument to get() is null")
	}

	return tree.root.search(key, tree.height)
}

func (h *Node) search(key string, ht int) Data {
	children := h.children

	// external node
	if ht == 0 {
		for j := 0; j < h.m; j++ {
			if key == children[j].key {
				return children[j].value
			}
		}
	} else {  // internal node
		for j := 0; j < h.m; j++ {
			if j == h.m-1 || key < children[j+1].key {
				return children[j].next.search(key, ht-1)
			}
		}
	}

	return nil
}

// Put
// Inserts the key-value pair into the symbol table, overwriting the old value
// with the new value if the key is already in the symbol table.
func (tree *BTree) Put(key string, value Data) {
	if key == "" {
		log.Fatalln("argument key to put() is null")
	}

	u := tree.root.insert(key, value, tree.height)
	tree.n++
	if u == nil {
		return
	}

	// need to split root
	t := newNode(2)
	t.children[0] = newEntry(tree.root.children[0].key, nil, tree.root)
	t.children[1] = newEntry(u.children[0].key, nil, u)
	tree.root = t
	tree.height++
}

func (h *Node) insert(key string, value Data, ht int) *Node {
	var j int
	t := newEntry(key, value, nil)

	// external node
	if ht == 0 {
		for j = 0; j < h.m; j++ {
			if key <= h.children[j].key {
				break
			}
		}
	} else {  // internal node
		for j = 0; j < h.m; j++ {
			if j == h.m-1 || key < h.children[j+1].key {
				x := h.children[j]
				j++
				u := x.next.insert(key, value, ht-1)

				if u == nil {
					return nil
				}

				t.key = u.children[0].key
				t.value = nil
				t.next = u

				break
			}
		}
	}

	for i := h.m; i > j; i-- {
		h.children[i] = h.children[i-1]
	}

	h.children[j] = t
	h.m++

	if h.m == M {
		return split(h)
	}

	return nil
}

// split node in half
func split(h *Node) *Node {
	t := newNode(M/2)
	h.m = M/2
	for j := 0; j < M/2; j++ {
		t.children[j] = h.children[M/2+j]
		h.children[M/2+j] = nil
	}
	return t
}

// Returns a string representation of this B-tree (for debugging).
func (tree *BTree) String() string {
	return tree.root.string(tree.height, "") + "\n"
}

func (h *Node) string(ht int, indent string) string {
	s := strings.Builder{}
	children := h.children

	if ht == 0 {
		for j := 0; j < h.m; j++ {
			s.WriteString(indent + children[j].key + " " + fmt.Sprintf("%v", children[j].value) + "\n")
		}
	} else {
		for j := 0; j < h.m; j++ {
			s.WriteString(indent + "(" + children[j].key + " " + fmt.Sprintf("%v", ht) + ")\n")
			s.WriteString(children[j].next.string(ht-1, indent + "     "))
		}
	}
	return s.String()
}
