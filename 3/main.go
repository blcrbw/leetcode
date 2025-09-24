package main

// 594. Longest Harmonious Subsequence
// We define a harmonious array as an array where the difference between its maximum value and its minimum value is exactly 1.
// Given an integer array nums, return the length of its longest harmonious
// among all its possible subsequences.

type Char struct {
	val   rune
	index int
	prev  *Char
	next  *Char
}

func lengthOfLongestSubstring(s string) int {
	hash := make(map[rune]*Char)

	maxLn := 0
	var first, last *Char
	var prev *Char

	for idx, run := range s {
		char, ok := hash[run]
		if !ok || char.index < first.index {
			// New char in substring
			char = &Char{
				val:   run,
				index: idx,
			}
			if prev != nil {
				prev.next = char
				char.prev = prev
			}
			hash[run] = char
			last = char
			if first == nil {
				first = char
			}
		} else {
			if last != nil && first != nil {
				ln := last.index - first.index + 1
				if ln > maxLn {
					maxLn = ln
				}
			}

			first = char.next
			if first == nil {
				first = char
			}
			last.next = char
			char.prev = last
			last = char
			first.prev = nil
			char.index = idx
			char.next = nil
		}
		prev = char
	}
	if last != nil && first != nil {
		ln := last.index - first.index + 1
		if ln > maxLn {
			maxLn = ln
		}
	}

	return maxLn
}
