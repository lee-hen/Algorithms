package main

import (
	"fmt"
	"github.com/lee-hen/Algorithms/util"
	"log"
	"math"
)

type SparseVector struct {
	d int // dimension
	st map[int]float64 // the vector, represented by index-value pairs
}

// NewSparseVector
// Initializes a d-dimensional zero vector.
func NewSparseVector(dimension int) *SparseVector {
	return &SparseVector{
		d: dimension,
		st: make(map[int]float64),
	}
}

// Put
// Sets the ith coordinate of this vector to the specified value.
func Put(this *SparseVector, i int, value float64) {
	if i < 0 || i >= this.d {
		log.Fatalln("Illegal index")
	}
	if value == 0.0 {
		delete(this.st, i)
	} else {
		this.st[i] = value
	}
}

// Get
// Returns the ith coordinate of this vector.
func Get(this *SparseVector, i int) float64 {
	if i < 0 || i >= this.d {
		log.Fatalln("Illegal index")
	}

	if value, ok := this.st[i]; ok {
		return value
	}

	return 0.0
}

// Nnz
// Returns the number of nonzero entries in this vector.
func Nnz(this *SparseVector) int {
	return len(this.st)
}

// Dimension
// Returns the dimension of this vector.
func Dimension(this *SparseVector) int {
	return this.d
}

// Dot
// Returns the inner product of this vector with the specified vector
func Dot(this, that *SparseVector) float64 {
	if this.d != that.d {
		log.Fatalln("Illegal index")
	}
	var sum = 0.0

	// iterate over the vector with the fewest nonzeros
	if Nnz(this) <= Nnz(that) {
		for i := range this.st {
			if _, ok := that.st[i]; ok {
				sum += Get(this, i) * Get(that, i)
			}
		}
	} else {
		for i := range that.st {
			if _, ok := this.st[i]; ok {
				sum += Get(this, i) * Get(that, i)
			}
		}
	}

	return sum
}

// Magnitude
// Returns the Magnitude of this vector.
// This is also known as the L2 norm or the Euclidean norm.
func Magnitude(this *SparseVector) float64 {
	return math.Sqrt(Dot(this, this))
}

// Scale
// Returns the scalar-vector product of this vector with the specified scalar.
func Scale(this *SparseVector, alpha float64) *SparseVector {
	c := NewSparseVector(this.d)
	for i := range this.st {
		Put(c, i, alpha * Get(this, i))
	}
	return c
}

// Plus
// Returns the sum of this vector and the specified vector.
func Plus(this, that *SparseVector) *SparseVector {
	if this.d != that.d {
		log.Fatalln("Vector lengths disagree")
	}
	c := NewSparseVector(this.d)
	for i := range this.st {
		Put(c, i, Get(this, i))
	}

	for i := range that.st {
		Put(c, i, Get(that, i) + Get(c, i))
	}

	return c
}

func main(){
	a := NewSparseVector(10)
	b := NewSparseVector(10)

	Put(a,3, 0.50)
	Put(a,9, 0.75)
	Put(a,6, 0.11)
	Put(a,6, 0.00)
	Put(b,3, 0.60)
	Put(b,4, 0.90)

	fmt.Println("a =", a)
	fmt.Println("b =", b)

	fmt.Println("a dot b =", Dot(a, b))
	fmt.Println("a + b =", Plus(a, b))


	for i := 'A'; i <= 'Z'; i++ {
		hashCode := util.String(string(i))
		fmt.Println(hashCode* 0x7fffffff %10)
	}
}
