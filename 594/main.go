package main

// 594. Longest Harmonious Subsequence
// We define a harmonious array as an array where the difference between its maximum value and its minimum value is exactly 1.
// Given an integer array nums, return the length of its longest harmonious
// among all its possible subsequences.

func findContiniousLHS(nums []int) int {
	const diff = 1
	mn, mx := nums[0], nums[0]
	ln := 1
	maxLn := 1
	for _, num := range nums {
		if mn == mx {
			if num > mx && num == mx+diff {
				ln++
				mx = num
			} else if num < mn && num == mn-diff {
				ln++
				mn = num
			} else if num == mn {
				ln++
			} else {
				maxLn = max(ln, maxLn)
				ln = 1
			}
		} else {
			if num > mx {
				maxLn = max(ln, maxLn)
				ln = 1
			} else if num < mn {
				maxLn = max(ln, maxLn)
				ln = 1
			} else {
				ln++
			}
		}

	}
	maxLn = max(ln, maxLn)
	return maxLn
}

func findLHS(nums []int) int {
	const diff = 1
	maxLn := 0
	hash := make(map[int]int)
	for _, num := range nums {
		if _, ok := hash[num]; ok {
			hash[num]++
		} else {
			hash[num] = 1
		}
	}
	for num, ln := range hash {
		if nextLn, ok := hash[num+diff]; ok {
			maxLn = max(ln+nextLn, maxLn)
		}
	}

	return maxLn
}

func main() {
	findContiniousLHS([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15})
}
