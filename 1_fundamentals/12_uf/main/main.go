package main

import (
	UF "github.com/lee-hen/Algorithms/1_fundamentals/12_uf"

	"fmt"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		fmt.Println(err)
	}
	uf := UF.NewUF(n)

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
	}
}

