package avl_tree_st

import (
	"github.com/lee-hen/Algorithms/util"

	"fmt"
	"log"
)

type AVLTree struct {
	root *Node // The root node.
}

// NewAVLTree
// Initializes an empty symbol table.
func NewAVLTree() *AVLTree {
	return &AVLTree{}
}

// Node
// An inner node of the AVL tree.
type Node struct {
	Key string // key
	// the associated value
	// height of the subtree
	// number of nodes in subtree
	Value, height, size int

	// left subtree
	// right subtree
	Left, Right *Node
}

func newNode(key string, value, height, size int) *Node {
	return &Node{Key: key, Value: value, height: height, size: size}
}

// IsEmpty
// Checks if the symbol table is empty.
func (tree *AVLTree) IsEmpty() bool {
	return tree.root == nil
}

// Size
// Returns the number of key-value pairs in this symbol table.
func (tree *AVLTree) Size() int {
	return size(tree.root)
}

// Returns the number of nodes in the subtree.
func size(x *Node) int {
	if x == nil {
		return 0
	}
	return x.size
}


// Height
// Returns the height of the internal AVL tree. It is assumed that the
// height of an empty tree is -1 and the height of a tree with just one node
// is 0.
func (tree *AVLTree) Height() int {
	return height(tree.root)
}

// Returns the height of the subtree.
func height(x *Node) int {
	if x == nil {
		return -1
	}
	return x.height
}


// Get
// Returns the value associated with the given key.
func (tree *AVLTree) Get(key string) (int, bool) {
	if key == "" {
		log.Fatalln("argument to get() is null")
	}

	if x := get(tree.root, key); x != nil {
		return  x.Value, true
	}

	return 0, false
}

// Returns value associated with the given key in the subtree or nil if no such key.
func get(x *Node, key string) *Node {
	if x == nil {
		return nil
	}

	if key < x.Key {
		return get(x.Left, key)
	} else if key > x.Key {
		return get(x.Right, key)
	} else {
		return x
	}
}

// Contains
// Checks if the symbol table contains the given key.
func (tree *AVLTree) Contains(key string) bool {
	_, found := tree.Get(key)
	return found
}

// Put
// Inserts the specified key-value pair into the symbol table, overwriting
// the old value with the new value if the symbol table already contains the
// specified key.
func (tree *AVLTree) Put(key string, value int) {
	if key == "" {
		log.Fatalln("first argument to put() is null")
	}

	tree.root = put(tree.root, key, value)
}

// Inserts the key-value pair in the subtree. It overrides the old value
// with the new value.
func put(x *Node, key string, value int) *Node {
	if x == nil {
		return newNode(key, value, 0, 1)
	}

	if key < x.Key {
		x.Left = put(x.Left, key, value)
	} else if key > x.Key {
		x.Right = put(x.Right, key, value)
	} else {
		x.Value = value
		return x
	}

	x.size = 1 + size(x.Left) + size(x.Right)
	x.height = 1 + util.Max(height(x.Left), height(x.Right))

	return balance(x)
}

// Restores the AVL tree property of the subtree.
func balance(x *Node) *Node {
	if balanceFactor(x) < -1 {
		if balanceFactor(x.Right) > 0 {
			x.Right = x.Right.rotateRight()
		}
		x = x.rotateLeft()
	} else if balanceFactor(x) > 1 {
		if balanceFactor(x.Left) < 0 {
			x.Left = x.Left.rotateLeft()
		}
		x = x.rotateRight()
	}

	return x
}

// Returns the balance factor of the subtree. The balance factor is defined
// as the difference in height of the left subtree and right subtree, in
// this order. Therefore, a subtree with a balance factor of -1, 0 or 1 has
// the AVL property since the heights of the two child subtrees differ by at
// most one.
func balanceFactor(x *Node) int {
	return height(x.Left) - height(x.Right)
}


// Rotates the given subtree to the right.
func (x *Node) rotateRight() *Node {
	y := x.Left
	x.Left = y.Right
	y.Right = x
	y.size = x.size // set origin root node's size
	x.size = 1 + size(x.Left) + size(x.Right)
	x.height = 1 + util.Max(height(x.Left), height(x.Right))
	y.height = 1 + util.Max(height(y.Left), height(y.Right))
	return y
}

// Rotates the given subtree to the left.
func (x *Node) rotateLeft() *Node {
	y := x.Right
	x.Right = y.Left
	y.Left = x
	y.size = x.size
	x.size = 1 + size(x.Left) + size(x.Right)
	x.height = 1 + util.Max(height(x.Left), height(x.Right))
	y.height = 1 + util.Max(height(y.Left), height(y.Right))
	return y
}

// Del
// Removes the specified key and its associated value from the symbol table
// (if the key is in the symbol table).
func (tree *AVLTree) Del(key string) {
	if key == "" {
		log.Fatalln("argument to delete() is null")
	}

	if !tree.Contains(key) {
		return
	}

	tree.root = del(tree.root, key)
}

// Removes the specified key and its associated value from the given
// subtree.
func del(x *Node, key string) *Node {
	if key < x.Key {
		x.Left = del(x.Left, key)
	} else if key > x.Key {
		x.Right = del(x.Right, key)
	} else {
		if x.Left == nil {
			return x.Right
		} else if x.Right == nil {
			return x.Left
		} else {
			y := x
			x = min(y.Right)
			x.Right = delMin(y.Right)
			x.Left = y.Left
		}
	}

	x.size = 1 + size(x.Left) + size(x.Right)
	x.height = 1 + util.Max(height(x.Left), height(x.Right))
	return balance(x)
}


// DelMin
// Removes the smallest key and associated value from the symbol table.
func (tree *AVLTree) DelMin() {
	if tree.IsEmpty() {
		log.Fatalln("called deleteMin() with empty symbol table")
	}

	tree.root = delMin(tree.root)
}

// Removes the smallest key and associated value from the given subtree.
func delMin(x *Node) *Node {
	if x.Left == nil {
		return x.Right
	}

	x.Left = delMin(x.Left)

	x.size = 1 + size(x.Left) + size(x.Right)
	x.height = 1 + util.Max(height(x.Left), height(x.Right))
	return balance(x)
}

// DelMax
// Removes the largest key and associated value from the symbol table.
func (tree *AVLTree) DelMax() {
	if tree.IsEmpty() {
		log.Fatalln("called deleteMax() with empty symbol table")
	}
	tree.root = delMax(tree.root)
}

// Removes the largest key and associated value from the given subtree.
func delMax(x *Node) *Node {
	if x.Right == nil {
		return x.Left
	}

	x.Right = delMax(x.Right)

	x.size = 1 + size(x.Left) + size(x.Right)
	x.height = 1 + util.Max(height(x.Left), height(x.Right))
	return balance(x)
}

// Ordered symbol table methods.

// Min
// Returns the smallest key in the symbol table.
func (tree *AVLTree) Min() string {
	if tree.IsEmpty() {
		log.Fatalln("calls min() with empty symbol table")
	}

	return min(tree.root).Key
}

// Returns the node with the smallest key in the subtree.
func min(x *Node) *Node {
	if x.Left == nil {
		return x
	}

	return min(x.Left)
}

// Max
// Returns the largest key in the symbol table.
func (tree *AVLTree) Max() string {
	if tree.IsEmpty() {
		log.Fatalln("calls max() with empty symbol table")
	}

	return max(tree.root).Key
}

// Returns the node with the largest key in the subtree.
func max(x *Node) *Node{
	if x.Right == nil {
		return x
	}

	return max(x.Right)
}

// Floor
// Returns the largest key in the symbol table less than or equal to key
func (tree *AVLTree) Floor(key string) string {
	if key == "" {
		log.Fatalln("argument to floor() is null")
	}

	if tree.IsEmpty() {
		log.Fatalln("calls floor() with empty symbol table")
	}

	x := floor(tree.root, key)

	if x == nil {
		fmt.Printf("argument: %s to floor() is too small\n", key)
		return ""
	}

	return x.Key
}

// Returns the node in the subtree with the largest key less than or equal
// to the given key.
func floor(x *Node, key string) *Node {
	if x == nil {
		return nil
	}

	if key == x.Key {
		return x
	}

	if key < x.Key {
		return floor(x.Left, key)
	}

	y := floor(x.Right, key)
	if y != nil {
		return y
	}

	return x
}

// Ceiling
// Returns the smallest key in the symbol table greater than or equal to key.
func (tree *AVLTree) Ceiling(key string) string {
	if key == "" {
		log.Fatalln("argument to ceiling() is null")
	}

	if tree.IsEmpty() {
		log.Fatalln("calls ceiling() with empty symbol table")
	}

	x := ceiling(tree.root, key)

	if x == nil {
		fmt.Printf("argument: %s to ceiling() is too large\n", key)
		return ""
	}

	return x.Key
}

// Returns the node in the subtree with the smallest key greater than or
// equal to the given key.
func ceiling(x *Node, key string) *Node {
	if x == nil {
		return nil
	}

	if key == x.Key {
		return x
	}

	if key > x.Key {
		return ceiling(x.Right, key)
	}

	y := ceiling(x.Left, key)
	if y != nil {
		return y
	}

	return x
}

// Select
// Return the key in the symbol table of a given rank.
// This key has the property that there are rank keys in
// the symbol table that are smaller. In other words, this key is the
// (rank+1)st smallest key in the symbol table.
func (tree *AVLTree) Select(rank int) string {
	if rank < 0 || rank >= tree.Size() {
		log.Fatalf("argument to select() is invalid: : %d\n", rank)
	}

	return pick(tree.root, rank)
}

func pick(x *Node, rank int) string {
	if x == nil {
		return ""
	}
	leftSize := size(x.Left)
	if leftSize > rank {
		return pick(x.Left, rank)
	} else if leftSize < rank {
		return pick(x.Right, rank-leftSize-1)
	} else {
		return x.Key
	}
}

// Rank
// Return the number of keys in the symbol table strictly less than key.
func (tree *AVLTree) Rank(key string) int {
	if key == "" {
		log.Fatalln("argument to rank() is null")
	}

	return rank(key, tree.root)
}

func rank(key string, x *Node) int {
	if x == nil {
		return 0
	}

	if key < x.Key {
		return rank(key, x.Left)
	} else if key > x.Key {
		return 1 + size(x.Left) + rank(key, x.Right)
	} else {
		return size(x.Left)
	}
}

// Keys
// Returns all keys in the symbol table as an Iterable.
func (tree *AVLTree) Keys() []string {
	return tree.KeysInOrder()
}

// KeysInOrder
// Returns all keys in the symbol table following an in-order traversal.
func (tree *AVLTree) KeysInOrder() []string {
	keys := make([]string, 0)
	keysInOrder(tree.root, &keys)
	return keys
}

// Adds the keys in the subtree to queue following an in-order traversal.
func keysInOrder(x *Node, keys *[]string) {
	if x == nil {
		return
	}
	keysInOrder(x.Left, keys)
	*keys = append(*keys, x.Key)
	keysInOrder(x.Right, keys)
}

// KeysLevelOrder
// Returns all keys in the symbol table following a level-order traversal.
func (tree *AVLTree) KeysLevelOrder() []string {
	keys := make([]string, 0)

	queue := []*Node{tree.root}
	for len(queue) > 0 {
		var x *Node
		x, queue = queue[0], queue[1:]

		if x == nil {
			continue
		}

		keys = append(keys, x.Key)

		queue = append(queue, x.Left)
		queue = append(queue, x.Right)
	}

	return keys
}

// KeysBetween
// Returns all keys in the symbol table in the given range
func (tree *AVLTree) KeysBetween(lo, hi string) []string {
	if lo == "" {
		log.Fatalln("first argument to keys() is null")
	}

	if hi == "" {
		log.Fatalln("second argument to keys() is null")
	}

	keys := make([]string, 0)
	tree.root.keys(&keys, lo, hi)

	return keys
}

func (x *Node) keys(keys *[]string, lo, hi string) {
	if x == nil {
		return
	}

	if lo < x.Key {
		x.Left.keys(keys, lo, hi)
	}

	if lo <= x.Key && hi >= x.Key {
		*keys = append(*keys, x.Key)
	}

	if hi > x.Key {
		x.Right.keys(keys, lo, hi)
	}
}

// SizeBetween
// Returns the number of keys in the symbol table in the given range.
func (tree *AVLTree) SizeBetween(lo, hi string) int {
	if lo == "" {
		log.Fatalln("first argument to size() is null")
	}
	if hi == "" {
		log.Fatalln("second argument to size() is null")
	}

	if lo > hi {
		return 0
	}

	if tree.Contains(hi) {
		return tree.Rank(hi) - tree.Rank(lo) + 1
	} else {
		return tree.Rank(hi) - tree.Rank(lo)
	}
}

func Check(tree *AVLTree) bool {
	if !tree.isBST() {
		fmt.Println("Symmetric order not consistent")
	}

	if !tree.isAVL() {
		fmt.Println("AVL property not consistent")
	}

	if !tree.isSizeConsistent() {
		fmt.Println("Subtree counts not consistent")
	}

	if !tree.isRankConsistent() {
		fmt.Println("Ranks not consistent")
	}

	return tree.isBST() && tree.isAVL() && tree.isSizeConsistent() && tree.isRankConsistent()
}

func (tree *AVLTree) isAVL() bool {
	return isAVL(tree.root)
}

func isAVL(x *Node) bool {
	if x == nil {
		return true
	}
	bf := balanceFactor(x)
	if bf > 1 || bf < -1 {
		return false
	}
	return isAVL(x.Left) && isAVL(x.Right)
}

func (tree *AVLTree) isBST() bool {
	return isBST(tree.root, "", "")
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

func (tree *AVLTree) isSizeConsistent() bool {
	return isSizeConsistent(tree.root)
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

func (tree *AVLTree) isRankConsistent() bool {
	for i := 0; i < tree.Size(); i++ {
		if i != tree.Rank(tree.Select(i)) {
			return false
		}
	}

	for _, key := range tree.Keys() {
		if key != tree.Select(tree.Rank(key)) {
			return false
		}
	}

	return true
}
