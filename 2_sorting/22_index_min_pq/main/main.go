package main

import (
	idxMinPQ "github.com/lee-hen/Algorithms/2_sorting/22_index_min_pq"
	"github.com/lee-hen/Algorithms/util"

	"fmt"
	"math/rand"
	"time"
)

func main() {
	priorities := []float64{10.2, 65.1, 32.0, 21.3, 100.8, 85.5, 71.2, 3.33, -1.563}
	pq := idxMinPQ.NewIndexMinPQ(len(priorities))

	for i, priority := range priorities {
		pq.Insert(i, priority)
	}

	for _, idx := range pq.PQ() {
		fmt.Printf("idx:%d-priority: %f\n", idx, pq.PriorityOf(idx))
	}

	fmt.Println("--------------------------------------------------------")

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < len(priorities); i++ {
		if rand.Float64() < 0.5 {
			pq.DecreasePriority(i, priorities[i]-1)
		} else {
			pq.IncreasePriority(i, priorities[i]+1)
		}
	}

	for !pq.IsEmpty() {
		maxPrior := pq.MinPriority()
		idx := pq.DelMin()
		fmt.Printf("idx:%d-priority: %f\n", idx, maxPrior)
	}

	fmt.Println("--------------------------------------------------------")

	for i, priority := range priorities {
		pq.Insert(i, priority)
	}

	perm := make([]int, len(priorities))
	for i := range priorities {
		perm[i] = i
	}

	util.ShuffleIntSlice(perm)

	for _, idx := range perm {
		priority := pq.PriorityOf(idx)
		pq.Delete(idx)
		fmt.Printf("deleted: idx:%d-priority:%f\n", idx, priority)
		fmt.Printf("minPriority:%f\n", pq.MinPriority())
	}
}
