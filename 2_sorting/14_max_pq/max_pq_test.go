package max_pq

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func isMaxPQ(pq *MaxPQ) bool {
	return pq.isMaxHeap()
}

func TestCase1(t *testing.T) {
	var maxPQ = NewMaxPQ([]int{48, 12, 24, 7, 8, -5, 24, 391, 24, 56, 2, 6, 8, 41})
	maxPQ.Insert(76)
	require.Equal(t, true, isMaxPQ(maxPQ))
	require.Equal(t, 391, maxPQ.Max())
	require.Equal(t, 391, maxPQ.DelMax())
	require.Equal(t, true, isMaxPQ(maxPQ))
	require.Equal(t, 76, maxPQ.Max())
	require.Equal(t, 76, maxPQ.DelMax())
	require.Equal(t, true, isMaxPQ(maxPQ))
	require.Equal(t, 56, maxPQ.Max())
	maxPQ.Insert(87)
	require.Equal(t, true, isMaxPQ(maxPQ))
}

func TestCase2(t *testing.T) {
	var maxPQ = NewMaxPQ([]int{-4, 5, 10, 8, -10, -6, -4, -2, -5, 3, 5, -4, -5, -1, 1, 6, -7, -6, -7, 8})
	require.Equal(t, true, isMaxPQ(maxPQ))
}
