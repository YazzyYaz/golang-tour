package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	stringSlice := strings.Fields(s)
	for _, value := range stringSlice {
		m[value] += 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
