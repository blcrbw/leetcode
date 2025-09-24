package main

import "container/heap"

// 3510. Minimum Pair Removal to Sort Array II
// Given an array nums, you can perform the following operation any number of times:
//
//	Select the adjacent pair with the minimum sum in nums. If multiple such pairs exist, choose the leftmost one.
//	Replace the pair with their sum.
//
// Return the minimum number of operations needed to make the array non-decreasing.
// An array is said to be non-decreasing if each element is greater than or equal to its previous element (if it exists).
type Item struct {
	sum           int
	left          int
	right         int
	previous      *Item
	next          *Item
	externalIndex int
	heapIndex     int
}

type Heap []*Item

func (h Heap) Len() int {
	return len(h)
}

func (h Heap) Less(i, j int) bool {
	if h[i].sum == h[j].sum {
		return h[i].externalIndex < h[j].externalIndex
	}
	return h[i].sum < h[j].sum
}

func (h Heap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
	h[i].heapIndex = i
	h[j].heapIndex = j
}

func (h *Heap) Push(x any) {
	item := x.(*Item)
	item.heapIndex = len(*h)
	*h = append(*h, item)
}

func (h *Heap) Pop() any {
	old := *h
	n := len(old)
	item := old[n-1]
	*h = old[:n-1]
	return item
}

func minimumPairRemoval(nums []int) int {
	if len(nums) < 2 {
		return 0
	}
	var minHeap Heap
	heap.Init(&minHeap)
	var prevItem *Item = nil

	var toSort, result int

	for i := range len(nums) - 1 {
		if nums[i] > nums[i+1] {
			toSort++
		}

		currentItem := &Item{sum: nums[i] + nums[i+1], left: nums[i], right: nums[i+1], previous: prevItem, next: nil, externalIndex: i, heapIndex: i}
		if prevItem != nil {
			prevItem.next = currentItem
		}
		prevItem = currentItem
		heap.Push(&minHeap, currentItem)
	}

	for toSort > 0 {
		result++
		minSumPair := heap.Pop(&minHeap).(*Item)
		if minSumPair.left > minSumPair.right {
			toSort--
		}
		if left := minSumPair.previous; left != nil {
			wasSorted := left.left <= left.right

			left.sum += minSumPair.right
			left.right += minSumPair.right
			left.next = minSumPair.next
			isSorted := left.left <= left.right
			if wasSorted && !isSorted {
				toSort++
			} else if !wasSorted && isSorted {
				toSort--
			}
			heap.Fix(&minHeap, left.heapIndex)
		}
		if right := minSumPair.next; right != nil {
			wasSorted := right.left <= right.right

			right.sum += minSumPair.left
			right.left += minSumPair.left
			right.previous = minSumPair.previous
			isSorted := right.left <= right.right
			if wasSorted && !isSorted {
				toSort++
			} else if !wasSorted && isSorted {
				toSort--
			}
			heap.Fix(&minHeap, right.heapIndex)
		}
	}

	return result
}
