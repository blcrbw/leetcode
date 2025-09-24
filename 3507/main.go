package main

// 3510. Minimum Pair Removal to Sort Array I
// Given an array nums, you can perform the following operation any number of times:
//    Select the adjacent pair with the minimum sum in nums. If multiple such pairs exist, choose the leftmost one.
//    Replace the pair with their sum.
//
// Return the minimum number of operations needed to make the array non-decreasing.
// An array is said to be non-decreasing if each element is greater than or equal to its previous element (if it exists).

type Item struct {
	value int
	next  *Item
}

func minimumPairRemoval(nums []int) int {
	result := 0
	var breakItem *Item
	var prev *Item
	for i, num := range nums {
		item := Item{value: num}
		if i > 0 {
			prev.next = &item
		} else {
			breakItem = &item
		}
		prev = &item
	}

	for {
		if changed := minPairRemove(breakItem); !changed {
			break
		}
		result++
	}

	return result
}

func minPairRemove(item *Item) bool {
	if item.next == nil {
		return false
	}
	needReplace := false
	minPair := item.next.value + item.value
	replaceItem := item
	for item.next != nil {
		if item.value > item.next.value {
			needReplace = true
		}
		pair := item.value + item.next.value
		if pair < minPair {
			minPair = pair
			replaceItem = item
		}
		item = item.next
	}
	if needReplace {
		replaceItem.next = replaceItem.next.next
		replaceItem.value = minPair
		return true
	}
	return false
}
