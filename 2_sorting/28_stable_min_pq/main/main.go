package main

import (
	stableMinPQ "github.com/lee-hen/Algorithms/2_sorting/28_stable_min_pq"

	"fmt"
	"strings"
)

func main() {
	text := "it was the best of times it was the worst of times it was the " +
			 "age of wisdom it was the age of foolishness it was the epoch " +
			 "belief it was the epoch of incredulity it was the season of light " +
			 "it was the season of darkness it was the spring of hope it was the " +
			 "winter of despair"

	str := strings.Split(text, " ")

	tuples := make([]stableMinPQ.Tuple, 0)
	for i := range str {
		tuples = append(tuples, stableMinPQ.Tuple{
			Id: i,
			Name: str[i],
		})
	}

	pq := stableMinPQ.NewStableMinPQ(tuples)
	for !pq.IsEmpty() {
		fmt.Println(pq.DelMin())
	}
}