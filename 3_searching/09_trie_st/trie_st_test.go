package trie_st

import (
	"github.com/stretchr/testify/require"
	"testing"
)

// she sells sea shells by the sea shore

func TestCase1(t *testing.T) {
	trie := TrieST{}
	str := []string{"she", "sells", "sea", "shells", "by", "the", "sea", "shore"}

	for i, s := range str {
		trie.Put(s, i)
	}

	t.Log("Keys(\"\"):")
	for _, key := range trie.Keys() {
		if key != "" {
			t.Log(key)
		}
	}

	require.Equal(t, "shells", trie.LongestPrefixOf("shellsort"))
	require.Equal(t, "", trie.LongestPrefixOf("quicksort"))

	t.Log("KeysWithPrefix(\"sh\"):")
	for _, s := range trie.KeysWithPrefix("sh") {
		if s != "" {
			t.Log(s)
		}
	}

	t.Log("KeysThatMatch(\".he.l.\"):")
	for _, s := range trie.KeysThatMatch(".he.l.") {
		if s != "" {
			t.Log(s)
		}
	}
}
