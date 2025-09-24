package main

import "fmt"

func main() {
	//fmt.Println(minLengthAfterRemovals([]int{1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 4, 5, 6, 6, 6, 6, 6, 7, 8, 9}))
	//fmt.Println(minLengthAfterRemovals([]int{2, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4}))
	//fmt.Println(minLengthAfterRemovals([]int{0, 0, 0, 0, 0, 0, 0, 0, 2, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4}))
	fmt.Println(minLengthAfterRemovals([]int{1}))
}

func minLengthAfterRemovals(nums []int) int {
	var curVal, firstInd, maxNum int

	for cur, val := range nums {
		if curVal != val {
			curVal = val
			if maxNum < cur-firstInd {
				maxNum = cur - firstInd
			}
			firstInd = cur
		}
	}

	if firstInd < len(nums)-1 && maxNum < len(nums)-1-firstInd {
		// get max nums in the end of list
		maxNum = len(nums) - firstInd
	}

	if maxNum <= len(nums)/2 {
		return len(nums) % 2
	}
	return maxNum - (len(nums) - maxNum)
}
