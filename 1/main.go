package main

// 1. Two Sum
// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

func twoSum(nums []int, target int) []int {
	targetMinusInput := make(map[int]int)
	for idx, inp := range nums {
		targetMinusInput[target-inp] = idx
	}
	for idx, inp := range nums {
		targetIdx, ok := targetMinusInput[inp]
		if ok && targetIdx != idx {
			return []int{idx, targetIdx}
		}
	}
	return []int{}
}
