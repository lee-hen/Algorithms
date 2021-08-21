package main

import (
	suffix "github.com/lee-hen/Algorithms/5_context_or_beyond/03_suffix_array"

	"fmt"
)

// Lrs
// Returns the longest repeated substring of the specified string.
func Lrs(text string) string{
	sa := suffix.SuffixArray(text)
	lrs := ""
	for i := 1; i < len(text); i++ {
		length := sa.Lcp(i)
		if length > len(lrs) {
			lrs = text[sa.Index(i): sa.Index(i) + length]
		}
	}

	return lrs
}

func main() {
	text := "it was the best of times it was the worst of times " +
		"it was the age of wisdom it was the age of foolishness " +
		"it was the epoch of belief it was the epoch of incredulity " +
		"it was the season of light it was the season of darkness " +
		"it was the spring of hope it was the winter of despair "
	fmt.Println("'" + Lrs(text) + "'")

	text = "Oh! jolly is the gale," +
		"And a joker is the whale," +
		"A' flourishin' his tail,-" +
		"Such a funny, sporty, gamy, jesty, joky, hoky-poky " +
		"lad, is the Ocean, oh! " +
		"The scud all a flyin'," +
		"That's his flip only foamin';" +
		"When he stirs in the spicin',-" +
		"Such a funny, sporty, gamy, jesty, joky, hoky-poky " +
		"lad, is the Ocean, oh! " +
		"Thunder splits the ships," +
		"But he only smacks his lips," +
		"A tastin' of this flip,-" +
		"Such a funny, sporty, gamy, jesty, joky, hoky-poky " +
		"lad, is the Ocean, oh! "
	fmt.Println("'" + Lrs(text) + "'")

	text = "aaaaaaaaa"
	fmt.Println("'" + Lrs(text) + "'")

	text = "abcdefg"
	fmt.Println("'" + Lrs(text) + "'")
}
