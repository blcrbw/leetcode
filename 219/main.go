package main

// 219. Contains Duplicate II
// Given an integer array nums and an integer k, return true if there are two distinct indices i and j in the array such
// that nums[i] == nums[j] and abs(i - j) <= k.

func containsNearbyDuplicate(nums []int, k int) bool {
	// hash stores number: lastIndex pairs.
	hash := make(map[int]int)

	for i, num := range nums {
		if last, ok := hash[num]; ok {
			if i-last <= k {
				return true
			}
		}
		hash[num] = i
	}

	return false
}
