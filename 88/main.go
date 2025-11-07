package main

// 88. Merge Sorted Array
// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n,
// representing the number of elements in nums1 and nums2 respectively.
//
// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
//
// The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
// To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged,
// and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

func merge(nums1 []int, m int, nums2 []int, n int) {
	for m > 0 || n > 0 {
		if n < 1 {
			nums1[m+n-1] = nums1[m-1]
			m--
		} else if m < 1 {
			nums1[m+n-1] = nums2[n-1]
			n--
		} else {
			if nums2[n-1] > nums1[m-1] {
				nums1[m+n-1] = nums2[n-1]
				n--
			} else {
				nums1[m+n-1] = nums1[m-1]
				m--
			}
		}
	}
}
