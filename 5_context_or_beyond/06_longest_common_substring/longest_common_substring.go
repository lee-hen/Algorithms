package main

import (
	suffix "github.com/lee-hen/Algorithms/5_context_or_beyond/04_suffix_array_x"
	"github.com/lee-hen/Algorithms/util"
	"os"

	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

// return the longest common prefix of suffix s[p..] and suffix t[q..]
func lcp(s, t string, p, q int) string {
	n := util.Min(len(s)-p, len(t)-q)
	for i := 0; i < n; i++ {
		if util.CharAt(s, p+i) != util.CharAt(t, q+i) {
			return s[p:p+i]
		}
	}

	return s[p:p+n]
}

// compare suffix s[p..] and suffix t[q..]
func compare(s, t string, p, q int) int {
	n := util.Min(len(s)-p, len(t)-q)
	for i := 0; i < n; i++ {
		if util.CharAt(s, p+i) != util.CharAt(t, q+i) {
			return util.CharAt(s, p+i) - util.CharAt(t, q+i)
		}
	}

	if len(s)-p < len(t)-q {
		return -1
	} else if len(s)-p > len(t)-q {
		return +1
	} else {
		return 0
	}
}

// Lcs
// Returns the longest common string of the two specified strings.
func Lcs(s, t string) string {
	suffix1 := suffix.SuffixArrayX(s)
	suffix2 := suffix.SuffixArrayX(t)

	// find longest common substring by "merging" sorted suffixes
	lcs := ""
	var i, j int
	for i < len(s) && j < len(t) {
		p := suffix1.Index(i)
		q := suffix2.Index(j)

		x := lcp(s, t, p, q)

		if len(x) > len(lcs) {
			lcs = x
		}

		if cmp := compare(s, t, p, q); cmp < 0 {
			i++
		} else if cmp >= 0 {
			j++
		}
	}

	return lcs
}

func main() {
	var err error
	var content1, content2 []byte
	pwd, _ := os.Getwd()
	content1, err = ioutil.ReadFile(pwd + "/data/mobydick.txt")
	if err != nil {
		log.Fatal(err)
	}
	content2, err = ioutil.ReadFile(pwd + "/data/tale.txt")
	if err != nil {
		log.Fatal(err)
	}
	space := regexp.MustCompile(`\r?\n?\s+`)


	text1 := strings.Trim(string(content1), " ")
	text1 = space.ReplaceAllString(text1, " ")

	text2 := strings.Trim(string(content2), " ")
	text2 = space.ReplaceAllString(text2, " ")

	fmt.Println("'" + Lcs(text1, text2) + "'")
}
