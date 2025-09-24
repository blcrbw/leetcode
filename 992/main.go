package main

// 992. Subarrays with K Different Integers
// Given an integer array nums and an integer k, return the number of good subarrays of nums.
// A good array is an array where the number of different integers in that array is exactly k.
//    For example, [1,2,3,1,2] has 3 different integers: 1, 2, and 3.
// A subarray is a contiguous part of an array.

func subarraysWithAtMostKDistinct(nums []int, k int) int {
	if k == 0 {
		return 0
	}

	hash := make(map[int]int)
	unique := 0
	result := 0
	left, right := 0, 0
	for right = 0; right < len(nums); right++ {
		hash[nums[right]]++
		if hash[nums[right]] == 1 {
			unique++
		}

		for unique > k {
			hash[nums[left]]--
			if hash[nums[left]] == 0 {
				unique--
			}
			left++
		}
		result += right - left + 1
	}
	return result
}

func subarraysWithKDistinct(nums []int, k int) int {
	return subarraysWithAtMostKDistinct(nums, k) - subarraysWithAtMostKDistinct(nums, k-1)
}

//func subarraysWithKDistinct(nums []int, k int) int {
//	cur := make(map[int]interface{})
//	l := len(nums)
//	if l < k {
//		return 0
//	}
//
//	res := 0
//	subStart, subEnd := 0, 0
//	for _, num := range nums {
//		cur[num] = nil
//		subEnd++
//		if len(cur) >= k {
//			break
//		}
//	}
//	if len(cur) < k {
//		return 0
//	}
//
//	for i := subEnd; i < l; i++ {
//		if _, ok := cur[nums[i]]; ok {
//			subEnd++
//		} else {
//			res += calcUnique(nums[subStart:subEnd], k)
//			subEnd++
//			cur = make(map[int]interface{})
//			for j := i; j >= 0; j-- {
//				cur[nums[j]] = nil
//				subStart = j
//				if len(cur) > k {
//					delete(cur, nums[j])
//					subStart = j + 1
//					break
//				}
//			}
//		}
//	}
//	res += calcUnique(nums[subStart:subEnd], k)
//	return res
//}
//
//func calcUnique(nums []int, k int) int {
//	l := len(nums)
//	var hash map[int]interface{}
//	// Account size == l by default.
//	res := 1
//	for size := k; size < l; size++ {
//		for i := 0; i <= l-size; i++ {
//			hash = make(map[int]interface{})
//			for _, val := range nums[i : size+i] {
//				hash[val] = nil
//			}
//			if len(hash) == k {
//				res++
//			}
//		}
//	}
//	return res
//}
