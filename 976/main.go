package main

import "sort"

// 976. Largest Perimeter Triangle
// Given an integer array nums, return the largest perimeter of a triangle with a non-zero area, formed from three of
// these lengths. If it is impossible to form any triangle of a non-zero area, return 0.

func largestPerimeter(nums []int) int {
	res := 0

	sort.Ints(nums)
	i := len(nums) - 1

	for i > 1 {
		if nums[i] > 2*nums[i-1] || nums[i] >= nums[i-1]+nums[i-2] {
			i--
		} else {
			res = nums[i] + nums[i-1] + nums[i-2]
			break
		}
	}

	return res
}
