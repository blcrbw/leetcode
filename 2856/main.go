package main

import "fmt"

func main() {
	// 0 - 1
	// 0, 0 - 1 + 2 = 3
	// 0, 0, 0 - 3 + 2 + 1
	//

	fmt.Println(zeroFilledSubarray([]int{1, 3, 0, 0, 2, 0, 0, 4}))
	fmt.Println(zeroFilledSubarray([]int{0, 0, 0, 2, 0, 0}))
	fmt.Println(zeroFilledSubarray([]int{10, 10, 10, 2, 0, 10}))
	fmt.Println(zeroFilledSubarray([]int{10, 10, 10, 2, 10, 10}))
	fmt.Println(zeroFilledSubarray([]int{}))
}

func zeroFilledSubarray(nums []int) int64 {
	var result int64
	var numZeros int64 = 0

	for _, i := range nums {
		if i == 0 {
			numZeros++
			result += numZeros
		} else {
			numZeros = 0
		}
	}

	return result
}
