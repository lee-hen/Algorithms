package binary_search_st

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestCase1(t *testing.T) {
	test := "S E A R C H E X A M P L E"
	keys := strings.Split(test, " ")

	st := NewBinarySearchST(len(keys))
	for i := 0; i < len(keys); i++ {
		st.Put(keys[i], i)
	}

	require.True(t, IsSorted(st))
	require.True(t, RankCheck(st))
}
