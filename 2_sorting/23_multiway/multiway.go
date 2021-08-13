package main

import (
	idxMinPQ "github.com/lee-hen/Algorithms/2_sorting/22_index_min_pq"

	"fmt"
)

// The IndexMinPQ client Multiway on page 322 solves the multiway merge problem: it merges together several sorted input
// streams into one sorted output stream. This problem arises in many applications: the streams might be the output of scientific instruments
// (sorted by time), lists of information from the web such as music or movies (sorted by title or artist name),
// commercial transactions (sorted by account number or time), or whatever. If you have the space,
// you might just read them all into an array and sort them, but with a priority queue, you can read input streams and
// put them in sorted order on the output no matter how long they are.
func merge(files [][]byte) string {
	sorted := make([]byte, 0)

	pq := idxMinPQ.NewIndexMinPQ(len(files))
	fileIndices := make(map[int]int)

	for idx := range files {
		pq.Insert(idx, float64(files[idx][0]))
		fileIndices[idx] = 0
	}

	for !pq.IsEmpty() {
		filename := byte(pq.MinPriority())
		idx := pq.DelMin()
		sorted = append(sorted, filename)

		if fileIndices[idx] == len(files[idx])-1 {
			continue
		}

		fileIndices[idx]++
		pq.Insert(idx, float64(files[idx][fileIndices[idx]]))
	}

	return string(sorted)
}

func main() {
	files := make([][]byte, 0)
	files = append(files, []byte{'A', 'B', 'C', 'F', 'G', 'I', 'I', 'Z'})
	files = append(files, []byte{'B', 'D', 'H', 'P', 'Q', 'Q'})
	files = append(files, []byte{'A', 'B', 'E', 'F', 'J', 'N'})
	fmt.Println("--------")
	fmt.Println(merge(files))
}
