package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := map[string]int{}
	for _, f := range strings.Fields(s) {
		count := m[f]
		m[f] = count + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
