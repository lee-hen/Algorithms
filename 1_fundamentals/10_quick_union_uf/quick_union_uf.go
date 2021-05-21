package main

import (
	"fmt"
	"log"
)

// 1.5.7 Develop classes QuickUnionUF and QuickFindUF that implement quick-union and quick-find, respectively.

type QuickUnionUF struct {
	count int
	parent []int
}

func NewQuickUnionUF(n int) *QuickUnionUF {
	parent := make([]int, n, n)
	for i := 0; i < n; i++ {
		parent[i] = i
	}
	return &QuickUnionUF {
		count: n,
		parent: parent,
	}
}

func (uf *QuickUnionUF) Count() int {
	return uf.count
}

func (uf *QuickUnionUF) Find(p int) int {
	uf.validate(p)
	for p != uf.parent[p] {
		p = uf.parent[p]
	}
	return p
}

func (uf *QuickUnionUF) validate (p int) {
	n := len(uf.parent)
	if p < 0 || p >= n {
		log.Fatalf("index %d is not between 0 and %d", p, n-1 )
	}
}

func (uf *QuickUnionUF) Connected(p, q int) bool {
	return uf.Find(p) == uf.Find(q)
}

func (uf *QuickUnionUF) Union(p, q int) {
	rootP := uf.Find(p)
	rootQ := uf.Find(q)

	if rootP == rootQ {
		return
	}
	uf.parent[rootP] = rootQ
	uf.count--
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
	}
	uf := NewQuickUnionUF(n)

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
