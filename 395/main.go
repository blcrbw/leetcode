package main

import "sort"

// 395. Longest Substring with At Least K Repeating Characters
// Given a string s and an integer k, return the length of the longest substring of s such that the frequency of each
// character in this substring is greater than or equal to k.
// if no such substring exists, return 0.

type Item struct {
	ids   []int
	count int
	full  bool
}

func longestSubstring(s string, k int) int {
	l := len(s)
	if l < k {
		return 0
	}
	cnt := make(map[rune]*Item)
	res := 0
	skips := []int{-1}
	for i, r := range s {
		item, ok := cnt[r]
		if !ok {
			item = &Item{
				ids:   []int{},
				count: 0,
			}
			cnt[r] = item
		}
		item.ids = append(item.ids, i)
		item.count++
		if k <= item.count {
			item.full = true
		}
	}
	for _, item := range cnt {
		if !item.full {
			skips = append(skips, item.ids...)
		}
	}
	if len(skips) == 1 {
		return l
	}
	skips = append(skips, l)
	sort.Ints(skips)

	for i := 0; i < len(skips)-1; i++ {
		if i >= l-1 {
			break
		}
		res = max(res, longestSubstring(s[skips[i]+1:skips[i+1]], k))
	}

	return res
}
