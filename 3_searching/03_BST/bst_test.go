package BST

import (
	"github.com/stretchr/testify/require"

	"strings"
	"testing"
)

func TestCase1(t *testing.T) {
	test := "S E A R C H E X A M P L E"
	keys := strings.Split(test, " ")

	st := NewBST()
	for i := 0; i < len(keys); i++ {
		st.Put(keys[i], i)
	}

	require.True(t, Check(st))
}
