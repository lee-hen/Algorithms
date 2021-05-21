package main

import (
	"fmt"
	"log"
)

type WeightedQuickUnionUF struct {
	count int
	parent, size []int
}

func NewWeightedQuickUnionUF(n int) *WeightedQuickUnionUF {
	parent := make([]int, n, n)
	size := make([]int, n, n)
	for i := 0; i < n; i++ {
		parent[i] = i
		size[i] = 1
	}
	return &WeightedQuickUnionUF {
		count: n,
		parent: parent,
		size: size,
	}
}

func (uf *WeightedQuickUnionUF) Count() int {
	return uf.count
}

func (uf *WeightedQuickUnionUF) Find(p int) int {
	uf.validate(p)
	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p
}

func (uf *WeightedQuickUnionUF) validate (p int) {
	n := len(uf.parent)
	if p < 0 || p >= n {
		log.Fatalf("index %d is not between 0 and %d", p, n-1 )
	}
}

func (uf *WeightedQuickUnionUF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *WeightedQuickUnionUF) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)

	if rootP == rootQ {
		return
	}

	if uf.size[rootP] < uf.size[rootQ] {
		uf.parent[rootP] = rootQ
		uf.size[rootQ] += uf.size[rootP]
	} else {
		uf.parent[rootQ] = rootP
		uf.size[rootP] += uf.size[rootQ]
	}
	uf.count--
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
	}
	uf := NewWeightedQuickUnionUF(n)

	for {
		var p, q int
		_, err = fmt.Scan(&p)
		if err != nil {
			fmt.Println(err)
		}
		_, err = fmt.Scan(&q)
		if err != nil {
			fmt.Println(err)
		}

		if uf.Find(p) == uf.Find(q) {
			continue
		}
		uf.Union(p, q)
		fmt.Println(p, q, uf.Count())
		fmt.Println(uf.parent)
	}
}
