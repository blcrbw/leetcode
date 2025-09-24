package main

// 1358. Number of Substrings Containing All Three Characters
// Given a string s consisting only of characters a, b and c.
// Return the number of substrings containing at least one occurrence of all these characters a, b and c.

func numberOfSubstrings(s string) int {
	runes := []rune(s)
	left := 0
	hash := map[rune]int{
		'a': 0,
		'b': 0,
		'c': 0,
	}
	res := 0
	for right, r := range s {
		hash[r]++
		for hash['a'] > 0 && hash['b'] > 0 && hash['c'] > 0 {
			res += len(s) - right
			hash[runes[left]]--
			left++
		}
	}

	return res
}
