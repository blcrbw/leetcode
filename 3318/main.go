package main

import "sort"

// 3318. Find X-Sum of All K-Long Subarrays I
// You are given an array nums of n integers and two integers k and x.
//
// The x-sum of an array is calculated by the following procedure:
//
//    Count the occurrences of all elements in the array.
//    Keep only the occurrences of the top x most frequent elements. If two elements have the same number of occurrences,
//        the element with the bigger value is considered more frequent.
//    Calculate the sum of the resulting array.
//
// Note that if an array has less than x distinct elements, its x-sum is the sum of the array.
//
// Return an integer array answer of length n - k + 1 where answer[i] is the x-sum of the subarray nums[i..i + k - 1].

func findXSum(nums []int, k int, x int) []int {
	l := len(nums)
	res := make([]int, 0, l-k+1)

	counts := make(map[int]int)

	// Count nums in initial subarray.
	for i := 0; i < k; i++ {
		_, ok := counts[nums[i]]
		if !ok {
			counts[nums[i]] = 1
		} else {
			counts[nums[i]]++
		}
	}

	// Cals first sum
	res = append(res, xSum(counts, x))

	// Remove first item and add next one to move sliding window
	for i := k; i < l; i++ {
		counts[nums[i-k]]--
		_, ok := counts[nums[i]]
		if !ok {
			counts[nums[i]] = 1
		} else {
			counts[nums[i]]++
		}
		// Calc current sum and append to results array.
		res = append(res, xSum(counts, x))
	}

	return res
}

func xSum(counts map[int]int, x int) int {
	slice := make([]int, 0, x)

	for n, _ := range counts {
		slice = append(slice, n)
	}

	// Sort subarray numbers by count and values.
	sort.Slice(slice, func(i, j int) bool {
		if counts[slice[i]] == counts[slice[j]] {
			return slice[i] > slice[j]
		}

		return counts[slice[i]] > counts[slice[j]]
	})

	res := 0

	// Sum up x most frequent numbers.
	for i := range min(x, len(slice)) {
		res += slice[i] * counts[slice[i]]
	}

	return res
}
