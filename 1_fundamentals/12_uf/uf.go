package main

import (
	"fmt"
	"log"
)

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

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
	}
	uf := NewUF(n)

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
