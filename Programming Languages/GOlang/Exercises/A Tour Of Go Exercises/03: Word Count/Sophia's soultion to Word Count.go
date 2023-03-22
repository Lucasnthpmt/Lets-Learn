// This solution doesn't run in the ide, only in the "A tour of go" playground
package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	var words []string
	words = strings.Fields(s)
	m := make(map[string]int)
	for _, v := range words {
		if m[v] == 0 {
			m[v] = 1
		} else {
			m[v] += 1
		}
	}

	return m
}

func main() {
	wc.Test(WordCount)
}
