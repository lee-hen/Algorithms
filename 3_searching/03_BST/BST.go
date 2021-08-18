package BST

import (
	"fmt"
	"github.com/lee-hen/Algorithms/util"
	"log"
)

// Proposition C. Search hits in a BST built from N random keys require ~ 2 ln N (about 1.39 lg N) compares, on the average.
// Proof: The number of compares used for a search hit ending at a given node is 1 plus the depth.
// Adding the depths of all nodes, we get a quantity known as the internal path length of the tree.
// Thus, the desired quantity is 1 plus the average internal path length of the BST,
// which we can analyze with the same argument that we used for PROPOSITION K in SECTION 2.3: Let CN be the total internal path length of a BST built from inserting N randomly ordered distinct keys,
// so that the average cost of a search hit is 1 +CN / N. We have C0 = C1 = 0 and for N > 1 we can write a recurrence relationship
// that directly mirrors the recursive BST structure: CN = N − 1 + (C0 + CN−1) / N + (C1 + CN−2)/N + . . . (CN−1 + C0)/N The N − 1 term
// takes into account that the root contributes 1 to the path length of each of the other N − 1 nodes in the tree; the rest of the expression accounts
// for the subtrees, which are equally likely to be any of the N sizes. After rearranging terms, this recurrence is nearly identical to the one that we solved in SECTION 2.3 for quicksort,
// and we can derive the approximation CN ~ 2N ln N.

// Proposition D. Insertions and search misses in a BST built from N random keys require ~ 2 ln N (about 1.39 lg N) compares, on the average.
// Proof: Insertions and search misses take one more compare, on the average, than search hits. This fact is not difficult to establish by induction

// Proposition E. In a BST, all operations take time proportional to the height of the tree, in the worst case.
// Proof: All of these methods go down one or two paths in the tree. The length of any path is no more than the height, by definition.

type BST struct {
	root *Node // root of BST
}

// Node
// BST helper node data type
type Node struct {
	Key string // sorted by key
	// Value associated data
	// size number of nodes in subtree
	Value, size int

	Left, Right *Node // left and right subtrees
}

func NewBST() *BST {
	return &BST{}
}

func newNode(key string, value, size int) *Node {
	return &Node{Key: key, Value: value, size: size}
}

// IsEmpty
// Returns true if this symbol table is empty.
func (bst *BST) IsEmpty() bool {
	return size(bst.root) == 0
}

// Size
// Returns the number of key-value pairs in this symbol table.
func (bst *BST) Size() int {
	return size(bst.root)
}

func size(x *Node) int {
	if x == nil {
		return 0
	}
	return x.size
}

// Contains
// Does this symbol table contain the given key?
// return true if this symbol table contains key and
// return false otherwise
func (bst *BST) Contains(key string) bool {
	if key == "" {
		log.Fatalln("argument to contains() is null")
	}

	_, found := bst.Get(key)
	return found
}

// Get
// Returns the value associated with the given key.
// Return 0, false if the key is not in the symbol table
func (bst *BST) Get(key string) (int, bool) {
	if x := get(bst.root, key); x != nil {
		return x.Value, true
	}

	return 0, false
}

func get(x *Node, key string) *Node {
	if key == "" {
		log.Fatalln("calls get() with a null key")
	}

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

// Put
// Inserts the specified key-value pair into the symbol table, overwriting the old
// value with the new value if the symbol table already contains the specified key.
func (bst *BST) Put(key string, value int) {
	if key == "" {
		log.Fatalln("calls put() with a null key")
	}

	bst.root = put(bst.root, key, value)
}

func put(x *Node, key string, value int) *Node {
	if x == nil {
		return newNode(key, value, 1)
	}

	if key < x.Key {
		x.Left = put(x.Left, key, value)
	} else if key > x.Key {
		x.Right = put(x.Right, key, value)
	} else {
		x.Value = value
	}

	x.size = 1 + size(x.Left) + size(x.Right)

	return x
}

// DelMin
// Removes the smallest key and associated value from the symbol table.
func (bst *BST) DelMin() {
	if bst.IsEmpty() {
		log.Fatalln("Symbol table underflow")
	}

	bst.root = delMin(bst.root)
}

func delMin(x *Node) *Node {
	if x.Left == nil {
		return x.Right
	}

	x.Left = delMin(x.Left)
	x.size = size(x.Left) + size(x.Right) + 1

	return x
}

// DelMax
// Removes the largest key and associated value from the symbol table.
func (bst *BST) DelMax() {
	if bst.IsEmpty() {
		log.Fatalln("Symbol table underflow")
	}

	bst.root = delMax(bst.root)
}

func delMax(x *Node) *Node {
	if x.Right == nil {
		return x.Left
	}

	x.Right = delMax(x.Right)
	x.size = size(x.Left) + size(x.Right) + 1

	return x
}

// Del
// Removes the specified key and its associated value from this symbol table
// (if the key is in this symbol table).
func (bst *BST) Del(key string) {
	if key == "" {
		log.Fatalln("calls delete() with a null key")
	}

	bst.root = del(bst.root, key)
}

func del(x *Node, key string) *Node {
	if x == nil {
		return nil
	}

	if key < x.Key {
		x.Left = del(x.Left, key)
	} else if key > x.Key {
		x.Right = del(x.Right, key)
	} else {
		if x.Right == nil {
			return x.Left
		}
		if x.Left == nil {
			return x.Right
		}

		t := x
		x = min(t.Right)
		x.Right = delMin(t.Right)
		x.Left = t.Left
	}

	x.size = size(x.Left) + size(x.Right) + 1
	return x
}

// Min
// Returns the smallest key in the symbol table.
func (bst *BST) Min() string {
	if bst.IsEmpty() {
		log.Fatalln("calls min() with empty symbol table")
	}

	return min(bst.root).Key
}

func min(x *Node) *Node{
	if x.Left == nil {
		return x
	}

	return min(x.Left)
}

// Max
// Returns the largest key in the symbol table.
func (bst *BST) Max() string {
	if bst.IsEmpty() {
		log.Fatalln("calls max() with empty symbol table")
	}

	return max(bst.root).Key
}

func max(x *Node) *Node{
	if x.Right == nil {
		return x
	}

	return max(x.Right)
}

// Floor
// Returns the largest key in the symbol table less than or equal to key
func (bst *BST) Floor(key string) string {
	if key == "" {
		log.Fatalln("argument to floor() is null")
	}

	if bst.IsEmpty() {
		log.Fatalln("calls floor() with empty symbol table")
	}

	x := floor(bst.root, key)

	if x == nil {
		fmt.Printf("argument: %s to floor() is too small\n", key)
		return ""
	}

	return x.Key
}

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

	t := floor(x.Right, key)
	if t != nil {
		return t
	}

	return x
}

func (bst *BST) Floor2(key string) string {
	k := floor2(bst.root, key, "")
	if k == "" {
		fmt.Printf("argument: %s to floor2() is too small\n", key)
		return ""
	}

	return k
}

func floor2(x *Node, key, best string) string {
	if x == nil {
		return best
	}

	if key < x.Key {
		return floor2(x.Left, key, best)
	} else if key > x.Key {
		return floor2(x.Right, key, x.Key)
	}

	return x.Key
}

// Ceiling
// Returns the smallest key in the symbol table greater than or equal to key.
func (bst *BST) Ceiling(key string) string {
	if key == "" {
		log.Fatalln("argument to ceiling() is null")
	}

	if bst.IsEmpty() {
		log.Fatalln("calls ceiling() with empty symbol table")
	}

	x := ceiling(bst.root, key)

	if x == nil {
		fmt.Printf("argument: %s to ceiling() is too large\n", key)
		return ""
	}

	return x.Key
}

func ceiling(x *Node, key string) *Node {
	if x == nil {
		return nil
	}

	if key == x.Key {
		return x
	}

	if key < x.Key {
		t := ceiling(x.Left, key)
		if t != nil {
			return t
		}

		return x
	}

	return ceiling(x.Right, key)
}

// Select
// Return the key in the symbol table of a given rank.
// This key has the property that there are rank keys in
// the symbol table that are smaller. In other words, this key is the
// (rank+1)st smallest key in the symbol table.
func (bst *BST) Select(rank int) string {
	if rank < 0 || rank >= bst.Size() {
		log.Fatalf("argument to select() is invalid: : %d\n", rank)
	}

	return pick(bst.root, rank)
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
func (bst *BST) Rank(key string) int {
	if key == "" {
		log.Fatalln("argument to rank() is null")
	}

	return rank(key, bst.root)
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
func (bst *BST) Keys() []string {
	if bst.IsEmpty() {
		return []string{}
	}

	return bst.KeysBetween(bst.Min(), bst.Max())
}

// KeysBetween
// Returns all keys in the symbol table in the given range
func (bst *BST) KeysBetween(lo, hi string) []string {
	if lo == "" {
		log.Fatalln("first argument to keys() is null")
	}

	if hi == "" {
		log.Fatalln("second argument to keys() is null")
	}

	keys := make([]string, 0)
	bst.root.keys(&keys, lo, hi)

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
func (bst *BST) SizeBetween(lo, hi string) int {
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

// Height
// Returns the height of the BST (for debugging).
func (bst *BST) Height() int {
	return height(bst.root)
}

func height(x *Node) int {
	if x == nil {
		return -1
	}

	return 1 + util.Max(height(x.Left), height(x.Right))
}

// LevelOrder
// Returns the keys in the BST in level order (for debugging).
func (bst *BST) LevelOrder() []string {
	keys := make([]string, 0)

	queue := []*Node{bst.root}
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

func Check(bst *BST) bool {
	if !bst.isBST() {
		fmt.Println("Not in symmetric order")
	}

	if !bst.isSizeConsistent() {
		fmt.Println("Subtree counts not consistent")
	}

	if !bst.isRankConsistent() {
		fmt.Println("Ranks not consistent")
	}

	return bst.isBST() && bst.isSizeConsistent() && bst.isRankConsistent()
}

func (bst *BST) isBST() bool {
	return isBST(bst.root, "", "")
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

func (bst *BST) isSizeConsistent() bool {
	return isSizeConsistent(bst.root)
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

func (bst *BST) isRankConsistent() bool {
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
