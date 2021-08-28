package directed_edge

import (
	"fmt"
	"log"
)

type Edge struct {
	v, w   int
	weight float64
}

func NewEdge(v, w int, weight float64) *Edge {
	if v < 0 {
		log.Fatalln("vertex index must be a non-negative integer")
	}
	if w < 0 {
		log.Fatalln("vertex index must be a non-negative integer")
	}
	return &Edge{
		v, w, weight,
	}
}

func (e *Edge) Weight() float64 {
	return e.weight
}

func (e *Edge) From() int {
	return e.v
}

func (e *Edge) To() int {
	return e.w
}

func (e *Edge) CompareTo(other *Edge) int {
	if e.weight == other.weight {
		return 0
	}

	if e.weight-other.weight > 0.0 {
		return 1
	}

	return -1
}

func (e *Edge) String() string {
	return fmt.Sprintf("%d->%d %.5f", e.v, e.w, e.weight)
}
