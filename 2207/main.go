package main

import (
	"fmt"
)

func main() {
	text := "fwymvreuftzgrcrxczjacqovduqaiig"
	pattern := "yy"
	fmt.Println(maximumSubsequenceCount(text, pattern))
}

func maximumSubsequenceCount(text string, pattern string) int64 {
	var p0s, p1s, p0Counts, p1Counts int64
	runes := []rune(text)
	p0, p1 := rune(pattern[0]), rune(pattern[1])
	l := len(text)
	for _, t := range runes {
		if t == p0 {
			p0s++
		}
		if t == p1 {
			p1s++
		}
	}
	p0m, p1m := p0s, p1s
	for i := 0; i < l; i++ {
		if runes[i] == p1 {
			p1s--
		}
		if runes[l-1-i] == p0 {
			p0s--
		}
		if runes[i] == p0 {
			p0Counts += p1s
		}
		if runes[l-1-i] == p1 {
			p1Counts += p0s
		}
	}
	return Max(p0Counts+p1m, p1Counts+p0m)
}

func Max(i, j int64) int64 {
	if i > j {
		return i
	}
	return j
}
