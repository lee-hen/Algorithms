package main

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

// 1.5.17 Random connections. Develop a UF client ErdosRenyi that takes an integer value N from the command line,
// generates random pairs of integers between 0 and N-1, calling connected() to determine
// if they are connected and then union() if not (as in our development client),
// looping until all sites are connected, and printing the number of connections generated.
// Package your program as a static method count() that takes N as argument and
// returns the number of connections and a main() that takes N from the command line, calls count(),
// and prints the returned value.
func count(n int) int {
	edges := 0
	uf := NewUF(n)
	rand.Seed(time.Now().UnixNano())

	for uf.Count() > 1 {
		i :=  rand.Intn(n)
		j := rand.Intn(n)
		uf.Union(i, j)
		edges++
	}
	return edges
}

func main() {
	var n, trials int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
	}
	_, err = fmt.Scan(&trials)
	if err != nil {
		fmt.Println(err)
	}
	edges := make([]int, trials)

	for t := 0; t < trials; t++ {
		edges[t] = count(n)
	}
	// fmt.Println(edges)

	fmt.Printf("1/2 n ln n = %f", 0.5*float64(n)*math.Log(float64(n)))
	fmt.Println("")
	var sum, mean, sd float64
	for i := 0; i < len(edges); i++ {
		sum += float64(edges[i])
	}
	mean = sum/float64(len(edges))
	fmt.Printf("mean       = %f", mean)
	fmt.Println("")
	for j := 0; j < len(edges); j++ {
		sd += math.Pow(float64(edges[j]) - mean, 2)
	}
	sd = math.Sqrt(sd/float64(len(edges)))
	fmt.Printf("stddev     = %f", sd)
	fmt.Println("")
}

type UF struct {
	count int
	rank []byte
	parent []int
}

func NewUF(n int) *UF {
	parent := make([]int, n, n)
	rank := make([]byte, n, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		rank[i] = 0
	}
	return &UF {
		count: n,
		parent: parent,
		rank: rank,
	}
}

func (uf *UF) Find(p int) int {
	uf.validate(p)
	for p != uf.parent[p] {
		uf.parent[p] = uf.parent[uf.parent[p]]  // path compression by halving
		p = uf.parent[p]
	}
	return p
}

func (uf *UF) Count() int {
	return uf.count
}

func (uf *UF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *UF) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)

	if rootP == rootQ {
		return
	}

	if uf.rank[rootP] < uf.rank[rootQ] {
		uf.parent[rootP] = rootQ
	} else if uf.rank[rootP] > uf.rank[rootQ] {
		uf.parent[rootQ] = rootP
	} else {
		uf.parent[rootQ] = rootP
		uf.rank[rootP]++
	}
	uf.count--
}

func (uf *UF) validate (p int) {
	n := len(uf.parent)
	if p < 0 || p >= n {
		log.Fatalf("index %d is not between 0 and %d", p, n-1 )
	}
}
