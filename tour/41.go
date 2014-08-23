package main

import (
	"strings"

	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	words := strings.Fields(s)
	for i := 0; i < len(words); i++ {
		ret[words[i]] += 1
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
