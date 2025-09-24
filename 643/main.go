package main

// 594. Longest Harmonious Subsequence
// We define a harmonious array as an array where the difference between its maximum value and its minimum value is exactly 1.
// Given an integer array nums, return the length of its longest harmonious
// among all its possible subsequences.

func findMaxAverage(nums []int, k int) float64 {
	var maxSum, sum int
	for i := 0; i < len(nums)-k+1; i++ {
		iter := k - 1
		sum = 0
		for iter >= 0 {
			sum += nums[i+iter]
			iter--
		}
		if i == 0 {
			maxSum = sum
		} else {
			maxSum = max(sum, maxSum)
		}
	}

	return float64(maxSum) / float64(k)
}
