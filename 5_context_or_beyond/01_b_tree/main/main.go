package main

import (
	btree "github.com/lee-hen/Algorithms/5_context_or_beyond/01_b_tree"

	"fmt"
)

func main() {
	st := btree.NewBTree()
	st.Put("www.cs.princeton.edu", "128.112.136.12")
	st.Put("www.cs.princeton.edu", "128.112.136.11")
	st.Put("www.princeton.edu",    "128.112.128.15")
	st.Put("www.yale.edu",         "130.132.143.21")
	st.Put("www.simpsons.com",     "209.052.165.60")
	st.Put("www.apple.com",        "17.112.152.32")
	st.Put("www.amazon.com",       "207.171.182.16")
	st.Put("www.ebay.com",         "66.135.192.87")
	st.Put("www.cnn.com",          "64.236.16.20")
	st.Put("www.google.com",       "216.239.41.99")
	st.Put("www.nytimes.com",      "199.239.136.200")
	st.Put("www.microsoft.com",    "207.126.99.140")
	st.Put("www.dell.com",         "143.166.224.230")
	st.Put("www.slashdot.org",     "66.35.250.151")
	st.Put("www.espn.com",         "199.181.135.201")
	st.Put("www.weather.com",      "63.111.66.11")
	st.Put("www.yahoo.com",        "216.109.118.65")


	fmt.Println("cs.princeton.edu: ", st.Get("www.cs.princeton.edu"))
	fmt.Println("hardvardsucks.com:", st.Get("www.harvardsucks.com"))
	fmt.Println("simpsons.com:     ", st.Get("www.simpsons.com"))
	fmt.Println("apple.com:        ", st.Get("www.apple.com"))
	fmt.Println("ebay.com:         ", st.Get("www.ebay.com"))
	fmt.Println("dell.com:         ", st.Get("www.dell.com"))
	fmt.Println()

	fmt.Println("size:   ", st.Size())
	fmt.Println("height:  ", st.Height())
	fmt.Println(st)
	fmt.Println()
}
