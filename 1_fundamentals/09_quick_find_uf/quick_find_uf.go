package main

import (
	"fmt"
	"log"
)

// 1.5.7 Develop classes QuickUnionUF and QuickFindUF that implement quick-union and quick-find, respectively.

type QuickFindUF struct {
	count int
	id []int
}

func NewQuickFindUF(n int) *QuickFindUF {
	id := make([]int, n, n)
	for i := 0; i < n; i++ {
		id[i] = i
	}
	return &QuickFindUF {
		count: n,
		id: id,
	}
}

func (uf *QuickFindUF) Count() int {
	return uf.count
}

func (uf *QuickFindUF) Find(p int) int {
	uf.validate(p)
	return uf.id[p]
}

func (uf *QuickFindUF) validate (p int) {
	n := len(uf.id)
	if p < 0 || p >= n {
		log.Fatalf("index %d is not between 0 and %d", p, n-1 )
	}
}

func (uf *QuickFindUF) Connected(p, q int) bool {
	uf.validate(p)
	uf.validate(q)
	return uf.id[p] == uf.id[q]
}

func (uf *QuickFindUF) Union(p, q int) {
	uf.validate(p)
	uf.validate(q)
	pID := uf.id[p]
	qID := uf.id[q]

	if pID == qID {
		return
	}

	for i := 0; i < len(uf.id); i++ {
		if uf.id[i] == pID {
			uf.id[i] = qID
		}
	}
	uf.count--
}

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
	}
	uf := NewQuickFindUF(n)

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
		fmt.Println(uf.id)
	}
}
