package main

import "sort"

// 15. 3Sum
// Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
// Notice that the solution set must not contain duplicate triplets.

func threeSum(nums []int) [][]int {
	res := make([][]int, 0)
	l := len(nums)
	sort.Ints(nums)

	for i := 0; i < l-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		target := -nums[i]
		left, right := i+1, l-1
		for left < right {
			sum := nums[left] + nums[right]

			if sum == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return res
}
