package red_black_bst

import (
	"github.com/lee-hen/Algorithms/util"

	"fmt"
	"log"
)

var root **Node
var RedBlackBst *RedBlackBST

type RedBlackBST struct {
	root *Node // root of BST
}

// Initializes an empty symbol table.
func init() {
	RedBlackBst = &RedBlackBST{}
	root = &RedBlackBst.root
}

const (
 	RED = true
	BLACK = false
)

// Node
// BST helper Node data type
type Node struct {
	Key string // key
	// Value associated data
	// subtree count
	Value, size int

	Color bool // color of parent link

	Left, Right *Node // left and right subtrees
}

func newNode(key string, color bool, value, size int) *Node {
	return &Node{Key: key, Color: color, Value: value, size: size}
}

// is node x red; false if x is null ?
func (h *Node) isRed() bool {
	if h == nil {
		return false
	}

	return h.Color == RED
}

// Returns the number of key-value pairs in this symbol table.
func size(x *Node) int {
	if x == nil {
		return 0
	}
	return x.size
}

// IsEmpty
// Is this symbol table empty?
// return true if this symbol table is empty and return false otherwise
func (bst *RedBlackBST) IsEmpty() bool {
	return *root == nil
}

// Size
// Returns the number of key-value pairs in this symbol table.
func (bst *RedBlackBST) Size() int {
	return size(*root)
}

// Standard BST search.

// Get
// Returns the value associated with the given key if the key is in the symbol table
// Return false, 0 if the key is not in the symbol table
func (bst *RedBlackBST) Get(key string) (bool, int) {
	if key == "" {
		log.Fatalln("argument to get() is null")
	}

	if x := get(*root, key); x != nil {
		return true, x.Value
	}

	return false, 0
}

// value associated with the given key in subtree rooted at x; null if no such key
func get(x *Node, key string) *Node {
	for x != nil {
		if key < x.Key {
			x = x.Left
		} else if key > x.Key {
			x = x.Right
		} else {
			return x
		}
	}

	return nil
}

// Contains
// Does this symbol table contain the given key?
// return true if this symbol table contains {@code key} and
// return false otherwise
func (bst *RedBlackBST) Contains(key string) bool {
	found, _ := bst.Get(key)
	return found
}

// Red-black tree insertion.

// Put
// Inserts the specified key-value pair into the symbol table, overwriting the old
// value with the new value if the symbol table already contains the specified key.
// Deletes the specified key (and its associated value) from this symbol table
// if the specified value is null.
func (bst *RedBlackBST) Put(key string, value int) {
	if key == "" {
		log.Fatalln("first argument to put() is null")
	}

	*root = put(*root, key, value)
	(*root).Color = BLACK
}

// insert the key-value pair in the subtree rooted at h
func put(h *Node, key string, value int) *Node {
	if h == nil {
		return newNode(key, RED, value, 1)
	}

	if key < h.Key {
		h.Left = put(h.Left, key, value)
	} else if key > h.Key {
		h.Right = put(h.Right, key, value)
	} else {
		h.Value = value
	}

	// fix-up any right-leaning links

	// 右->赤　左->黒 left rotate, !h.Left.isRed() not required
	if h.Right.isRed() && !h.Left.isRed()  {
		h = h.rotateLeft()
	}

	// 左->赤 左.左->赤 right rotate
	if h.Left.isRed() && h.Left.Left.isRed() {
		h = h.rotateRight()
	}

	// 右->赤　左->赤 flip color
	if h.Left.isRed() && h.Right.isRed() {
		flipColors(h)
	}

	h.size = 1 + size(h.Left) + size(h.Right)

	return h
}

// Red-black tree deletion.

// DelMin
// Removes the smallest key and associated value from the symbol table.
func (bst *RedBlackBST) DelMin() {
	if bst.IsEmpty() {
		log.Fatalln("BST underflow")
	}

	// if both children of root are black, set root to red
	if !(*root).Left.isRed() && !(*root).Right.isRed() {
		(*root).Color = RED
	}

	*root = delMin(*root)

	if !bst.IsEmpty() {
		(*root).Color = BLACK
	}
}

// delete the key-value pair with the minimum key rooted at h
func delMin(h *Node) *Node {
	// remove node on bottom level
	// (h must be RED by invariant)
	if h.Left == nil {
		return h.Right
	}

	// push red link down if necessary
	if !h.Left.isRed() && !h.Left.Left.isRed() {
		h = moveRedLeft(h)
	}

	// move down one level
	h.Left = delMin(h.Left)

	// fix right-leaning red links
	// and eliminate 4-nodes
	// on the way up
	return balance(h)
}

// DelMax
// Removes the largest key and associated value from the symbol table.
func (bst *RedBlackBST) DelMax() {
	if bst.IsEmpty() {
		log.Fatalln("BST underflow")
	}

	// if both children of root are black, set root to red
	if !(*root).Left.isRed() && !(*root).Right.isRed() {
		(*root).Color = RED
	}

	*root = delMax(*root)

	if !bst.IsEmpty() {
		(*root).Color = BLACK
	}
}

// delete the key-value pair with the maximum key rooted at h
// push reds down
// remove maximum
// fix right-leaning reds on the way up
func delMax(h *Node) *Node {
	// lean 3-nodes to the right
	if h.Left.isRed() {
		h = h.rotateRight()
	}

	// remove node on bottom level
	// (h must be RED by invariant)
	if h.Right == nil {
		return nil
	}

	// borrow from sibling if necessary
	if !h.Right.isRed() && !h.Right.Left.isRed() {
		h = moveRedRight(h)
	}

	// move down one level
	h.Right = delMax(h.Right)

	// fix right-leaning red links
	// and eliminate 4-nodes
	// on the way up
	return balance(h)
}

// Del
// Removes the specified key and its associated value from this symbol table
// (if the key is in this symbol table).
func (bst *RedBlackBST) Del(key string) {
	if key == "" {
		log.Fatalln("argument to delete() is null")
	}

	if !bst.Contains(key) {
		return
	}

	// if both children of root are black, set root to red
	if !(*root).Left.isRed() && !(*root).Right.isRed() {
		(*root).Color = RED
	}

	*root = del(*root, key)

	if !bst.IsEmpty() {
		(*root).Color = BLACK
	}
}

// delete the key-value pair with the given key rooted at h
func del(h *Node, key string) *Node {
	if key < h.Key { // LEFT
		// push red right if necessary move down (left)
		if !h.Left.isRed() && !h.Left.Left.isRed() {
			h = moveRedLeft(h)
		}
		h.Left = del(h.Left, key)
	} else {
		// the same as delete max start

		// rotate to push red right
		if h.Left.isRed() {
			h = h.rotateRight()
		}

		// EQUAL (at bottom)
		// delete node
		if key == h.Key && h.Right == nil {
			return nil
		}

		// push red right if necessary
		if !h.Right.isRed() && !h.Right.Left.isRed() {
			h = moveRedRight(h)
		}

		// the same as delete max end

		if key == h.Key {
			// replace current node with
			// successor key, value
			x := h.Right.min()
			h.Key = x.Key
			h.Value = x.Value

			//delete successor
			h.Right = delMin(h.Right)
		} else {
			// move down (right)
			h.Right = del(h.Right, key)
		}
	}

	// fix right-leaning red links
	// and eliminate 4-nodes
	// on the way up
	return balance(h)
}

// Red-black tree helper functions.

// left leaning
// make a left-leaning link lean to the right
func (h *Node) rotateRight() *Node {
	x := h.Left
	h.Left = x.Right
	x.Right = h
	x.Color = x.Right.Color // set origin root node's color at h
	x.Right.Color = RED
	x.size = h.size
	h.size = size(h.Left) + size(h.Right) + 1
	return x
}

// right leaning
// make a right-leaning link lean to the left
func (h *Node) rotateLeft() *Node {
	x := h.Right
	h.Right = x.Left
	x.Left = h
	x.Color = x.Left.Color // set origin root node's color at h
	x.Left.Color = RED
	x.size = h.size
	h.size = size(h.Left) + size(h.Right) + 1
	return x
}

// flip the colors of a node and its two children
func flipColors(h *Node) {
	h.Color = !h.Color
	h.Left.Color = !h.Left.Color
	h.Right.Color = !h.Right.Color
}

// Assuming that h is red and both h.left and h.left.left
// are black, make h.left or one of its children red.
func moveRedLeft(h *Node) *Node {
	flipColors(h)

	if h.Right.Left.isRed() {
		h.Right = h.Right.rotateRight()
		h = h.rotateLeft()
		flipColors(h)
	}
	return h
}

// Assuming that h is red and both h.right and h.right.left
// are black, make h.right or one of its children red.
func moveRedRight(h *Node) *Node {
	flipColors(h)

	// 2-3 node
	if h.Left.Left.isRed() {
		h = h.rotateRight()
		flipColors(h)
	}
	return h
}

// restore red-black tree invariant
func balance(h *Node) *Node {
	if h.Right.isRed() && !h.Left.isRed() {
		h = h.rotateLeft()
	}

	if h.Left.isRed() && h.Left.Left.isRed() {
		h = h.rotateRight()
	}

	if h.Left.isRed() && h.Right.isRed() {
		flipColors(h)
	}
	h.size = size(h.Right) + size(h.Left) + 1

	return h
}

// ---------------------------------------------------------------------------------------------------------------------

// Utility functions.

// Height
// Returns the height of the BST (for debugging).
func (bst *RedBlackBST) Height() int {
	return (*root).height()
}

func (h *Node) height() int {
	if h == nil {
		return -1
	}

	return 1 + util.Max(h.Left.height(), h.Right.height())
}

// Ordered symbol table methods.

// Min
// Returns the smallest key in the symbol table.
func (bst *RedBlackBST) Min() string {
	if bst.IsEmpty() {
		log.Fatalln("calls min() with empty symbol table")
	}

	return (*root).min().Key
}

// the smallest key in subtree rooted at x; null if no such key
func (h *Node) min() *Node{
	if h.Left == nil {
		return h
	}

	return h.Left.min()
}

// Max
// Returns the largest key in the symbol table.
func (bst *RedBlackBST) Max() string {
	if bst.IsEmpty() {
		log.Fatalln("calls max() with empty symbol table")
	}

	return (*root).max().Key
}

// the largest key in the subtree rooted at x; null if no such key
func (h *Node) max() *Node{
	if h.Right == nil {
		return h
	}

	return h.Right.max()
}

// Floor
// Returns the largest key in the symbol table less than or equal to key
func (bst *RedBlackBST) Floor(key string) string {
	if key == "" {
		log.Fatalln("argument to floor() is null")
	}

	if bst.IsEmpty() {
		log.Fatalln("calls floor() with empty symbol table")
	}

	x := (*root).floor(key)

	if x == nil {
		fmt.Printf("argument: %s to floor() is too small\n", key)
		return ""
	}

	return x.Key
}

// the largest key in the subtree rooted at x less than or equal to the given key
func (h *Node) floor(key string) *Node {
	if h == nil {
		return nil
	}

	if key == h.Key {
		return h
	}

	if key < h.Key {
		return h.Left.floor(key)
	}

	t := h.Right.floor(key)
	if t != nil {
		return t
	}

	return h
}

// Ceiling
// Returns the smallest key in the symbol table greater than or equal to key.
func (bst *RedBlackBST) Ceiling(key string) string {
	if key == "" {
		log.Fatalln("argument to ceiling() is null")
	}

	if bst.IsEmpty() {
		log.Fatalln("calls ceiling() with empty symbol table")
	}

	x := (*root).ceiling(key)

	if x == nil {
		fmt.Printf("argument: %s to ceiling() is too large\n", key)
		return ""
	}

	return x.Key
}

// the smallest key in the subtree rooted at x greater than or equal to the given key
func (h *Node) ceiling(key string) *Node {
	if h == nil {
		return nil
	}

	if key == h.Key {
		return h
	}

	if key < h.Key {
		t := h.Left.ceiling(key)
		if t != nil {
			return t
		}

		return h
	}

	return h.Right.ceiling(key)
}

// Select
// Return the key in the symbol table of a given rank.
// This key has the property that there are rank keys in
// the symbol table that are smaller. In other words, this key is the
// (rank+1)st smallest key in the symbol table.
func (bst *RedBlackBST) Select(rank int) string {
	if rank < 0 || rank >= bst.Size() {
		log.Fatalf("argument to select() is invalid: : %d\n", rank)
	}

	return (*root).pick(rank)
}

// Return key in BST rooted at x of given rank.
// Precondition: rank is in legal range.
func (h *Node) pick(rank int) string {
	if h == nil {
		return ""
	}
	leftSize := size(h.Left)
	if leftSize > rank {
		return h.Left.pick(rank)
	} else if leftSize < rank {
		return h.Right.pick(rank-leftSize-1)
	} else {
		return h.Key
	}
}

// Rank
// Return the number of keys in the symbol table strictly less than key.
func (bst *RedBlackBST) Rank(key string) int {
	if key == "" {
		log.Fatalln("argument to rank() is null")
	}

	return (*root).rank(key)
}

// number of keys less than key in the subtree rooted at h
func (h *Node) rank(key string) int {
	if h == nil {
		return 0
	}

	if key < h.Key {
		return h.Left.rank(key)
	} else if key > h.Key {
		return 1 + size(h.Left) + h.Right.rank(key)
	} else {
		return size(h.Left)
	}
}

// Range count and range search.

// Keys
// Returns all keys in the symbol table as an Iterable.
func (bst *RedBlackBST) Keys() []string {
	if bst.IsEmpty() {
		return []string{}
	}

	return bst.KeysBetween(bst.Min(), bst.Max())
}

// KeysBetween
// Returns all keys in the symbol table in the given range
func (bst *RedBlackBST) KeysBetween(lo, hi string) []string {
	if lo == "" {
		log.Fatalln("first argument to keys() is null")
	}

	if hi == "" {
		log.Fatalln("second argument to keys() is null")
	}

	keys := make([]string, 0)
	(*root).keys(&keys, lo, hi)

	return keys
}

// add the keys between lo and hi in the subtree rooted at x
// to the queue
func (h *Node) keys(keys *[]string, lo, hi string) {
	if h == nil {
		return
	}

	if lo < h.Key {
		h.Left.keys(keys, lo, hi)
	}

	if lo <= h.Key && hi >= h.Key {
		*keys = append(*keys, h.Key)
	}

	if hi > h.Key {
		h.Right.keys(keys, lo, hi)
	}
}

// SizeBetween
// Returns the number of keys in the symbol table in the given range.
func (bst *RedBlackBST) SizeBetween(lo, hi string) int {
	if lo == "" {
		log.Fatalln("first argument to size() is null")
	}
	if hi == "" {
		log.Fatalln("second argument to size() is null")
	}

	if lo > hi {
		return 0
	}

	if bst.Contains(hi) {
		return bst.Rank(hi) - bst.Rank(lo) + 1
	} else {
		return bst.Rank(hi) - bst.Rank(lo)
	}
}

// Check integrity of red-black tree data structure.

func Check(bst *RedBlackBST) bool {
	if !bst.isBST() {
		fmt.Println("Not in symmetric order")
	}

	if !bst.isSizeConsistent() {
		fmt.Println("Subtree counts not consistent")
	}

	if !bst.isRankConsistent() {
		fmt.Println("Ranks not consistent")
	}

	if !bst.is23(){
		fmt.Println("Not a 2-3 tree")
	}

	if !bst.isBalanced(){
		fmt.Println("Not balanced")
	}

	return bst.isBST() && bst.isSizeConsistent() && bst.isRankConsistent() && bst.is23() && bst.isBalanced()
}

func (bst *RedBlackBST) isBST() bool {
	return isBST(*root, "", "")
}

func isBST(x *Node, min, max string) bool {
	if x == nil {
		return true
	}

	if min != "" && x.Key <= min {
		return false
	}

	if max != "" && x.Key >= max {
		return false
	}

	return isBST(x.Left, min, x.Key) && isBST(x.Right, x.Key, max)
}

func (bst *RedBlackBST) isSizeConsistent() bool {
	return isSizeConsistent(*root)
}

func isSizeConsistent(x *Node) bool {
	if x == nil {
		return true
	}

	if x.size != size(x.Left) + size(x.Right) + 1 {
		return false
	}

	return isSizeConsistent(x.Left) && isSizeConsistent(x.Right)
}

func (bst *RedBlackBST) isRankConsistent() bool {
	for i := 0; i < bst.Size(); i++ {
		if i != bst.Rank(bst.Select(i)) {
			return false
		}
	}

	for _, key := range bst.Keys() {
		if key != bst.Select(bst.Rank(key)) {
			return false
		}
	}

	return true
}

// Does the tree have no red right links, and at most one (left)
// red links in a row on any path?
func (bst *RedBlackBST) is23() bool {
	return is23(*root)
}

func is23(x *Node) bool {
	if x == nil {
		return true
	}
	if x.Right.isRed() {
		return false
	}

	if x != *root && x.isRed() && x.Left.isRed() {
		return false
	}

	return is23(x.Left) && is23(x.Right)
}

// do all paths from root to leaf have same number of black edges?
func (bst *RedBlackBST) isBalanced() bool {
	var black int
	x := *root

	for x != nil {
		if !x.isRed() {
			black++
		}

		x = x.Left
	}

	return (*root).isBalanced(black)
}

// does every path from the root to a leaf have the given number of black links?
func (h *Node) isBalanced(black int) bool {
	if h == nil {
		return black == 0
	}

	if !h.isRed() {
		black--
	}

	return h.Left.isBalanced(black) && h.Right.isBalanced(black)
}
