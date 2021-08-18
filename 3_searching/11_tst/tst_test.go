package tst

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// she sells sea shells by the sea shore

func TestCase1(t *testing.T) {
	tst := TST{}
	str := []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"}

	for i, s := range str {
		tst.Put(s, i)
	}

	t.Log("Keys(\"\"):")
	for _, key := range tst.Keys() {
		if key != "" {
			t.Log(key)
		}
	}

	require.Equal(t, "shells", tst.LongestPrefixOf("shellsort"))
	require.Equal(t, "", tst.LongestPrefixOf("quicksort"))

	t.Log("KeysWithPrefix(\"sh\"):")
	for _, s := range tst.KeysWithPrefix("sh") {
		if s != "" {
			t.Log(s)
		}
	}

	t.Log("KeysThatMatch(\".he.l.\"):")
	for _, s := range tst.KeysThatMatch(".he.l.") {
		if s != "" {
			t.Log(s)
		}
	}
}

