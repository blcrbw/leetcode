package main

import "sort"

// 611. Valid Triangle Number
// Given an integer array nums, return the number of triplets chosen from the array that can make triangles if we take
// them as side lengths of a triangle.

func triangleNumber(nums []int) int {
	// triangle condition is the longest side is longer than sum of to other sides
	// sort array to be able to iterate from higher number to lower one to be able to compare the longest side with the
	// others.
	sort.Ints(nums)
	l := len(nums)
	if l < 3 {
		return 0
	}
	res := 0

	for i := l - 1; i >= 2; i-- {
		// iterate through the longest side
		j := 0
		k := i - 1
		for j < k {
			// then iterate through the longest side of the rest to find the minimal side that meet sum of lengths condition.
			if nums[i] < nums[j]+nums[k] {
				// sum up the number of suitable short and mid side pairs
				res += k - j
				// then check the shorter mid side one.
				k--
			} else {
				// if the current side is too short for specific longest and mid pair, then check the longer one.
				j++
			}
		}
	}

	return res
}
