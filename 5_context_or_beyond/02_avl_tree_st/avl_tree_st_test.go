package avl_tree_st

import (
	"github.com/stretchr/testify/require"

	"strings"
	"testing"
)

func TestCase1(t *testing.T) {
	test := "S E A R C H E X A M P L E"
	keys := strings.Split(test, " ")

	st := NewAVLTree()
	for i := 0; i < len(keys); i++ {
		st.Put(keys[i], i)
	}

	require.True(t, Check(st))
}
