package linear_probing_hash_st

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCase1(t *testing.T) {
	hashSt := NewHashST()
	for i := 32; i <= 127; i++ {
		hashSt.Put(string(rune(i)), i)
		require.True(t, Check(hashSt))
	}

	require.Equal(t, 96, hashSt.Size())

	for i := 32; i <= 127; i++ {
		key := string(rune(i))
		val, _ := hashSt.Get(key)
		require.Equal(t, i, val)
		hashSt.Delete(key)
		require.True(t, Check(hashSt))
	}

	require.Equal(t, 0, hashSt.Size())
}
