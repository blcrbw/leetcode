package main

import (
	"slices"
	"sort"
)

// 2009. Minimum Number of Operations to Make Array Continuous
// You are given an integer array nums. In one operation, you can replace any element in nums with any integer.
// nums is considered continuous if both of the following conditions are fulfilled:
//    All elements in nums are unique.
//    The difference between the maximum element and the minimum element in nums equals nums.length - 1.
// For example, nums = [4, 2, 5, 3] is continuous, but nums = [1, 2, 3, 5, 6] is not continuous.
// Return the minimum number of operations to make nums continuous.

func minOperations(nums []int) int {
	// target of continuous numbers
	l := len(nums)

	// remove duplicates
	hash := make(map[int]interface{})
	for _, n := range nums {
		hash[n] = nil
	}
	sorted := make([]int, 0)
	for n, _ := range hash {
		sorted = append(sorted, n)
	}

	// sort array
	slices.Sort(sorted)

	// walk through elements and find the max available distance between current element and value+len final one.
	maxContinuous := 0
	for i, v := range sorted {
		final := sort.SearchInts(sorted[i:], v+l)
		if final > maxContinuous {
			maxContinuous = final
		}
	}

	return l - maxContinuous
}
