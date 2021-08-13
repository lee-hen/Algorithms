package min_pq

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func isMinPQ(pq *MinPQ) bool {
	return pq.isMinHeap()
}


func TestCase1(t *testing.T) {
	var minPQ = NewMinPQ([]int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41})
	minPQ.Insert(76)
	require.Equal(t, true, isMinPQ(minPQ))
	require.Equal(t, -5, minPQ.Min())
	require.Equal(t, -5, minPQ.DelMin())
	require.Equal(t, true, isMinPQ(minPQ))
	require.Equal(t, 2, minPQ.Min())
	require.Equal(t, 2, minPQ.DelMin())
	require.Equal(t, true, isMinPQ(minPQ))
	require.Equal(t, 6, minPQ.Min())
	minPQ.Insert(87)
	require.Equal(t, true, isMinPQ(minPQ))
}

func TestCase2(t *testing.T) {
	var minPQ = NewMinPQ([]int{-4, 5, 10, 8, -10, -6, -4, -2, -5, 3, 5, -4, -5, -1, 1, 6, -7, -6, -7, 8})
	require.Equal(t, true, isMinPQ(minPQ))
}
