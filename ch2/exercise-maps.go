package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	ss := strings.Fields(s)
	for _, v := range ss {
		m[v]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
