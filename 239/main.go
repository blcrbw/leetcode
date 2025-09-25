package main

import (
	"container/heap"
)

// 239. Sliding Window Maximum
// You are given an array of integers nums, there is a sliding window of size k which is moving from the very left of
// the array to the very right. You can only see the k numbers in the window. Each time the sliding window moves right
// by one position.
// Return the max sliding window.

type Item struct {
	value     int
	next      *Item
	heapIndex int
}

type maxHeap struct {
	data []*Item
}

func (h *maxHeap) Len() int {
	return len(h.data)
}
func (h *maxHeap) Less(i, j int) bool {
	return h.data[i].value > h.data[j].value
}
func (h *maxHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.data[i].heapIndex, h.data[j].heapIndex = i, j
}
func (h *maxHeap) Push(x interface{}) {
	i := x.(*Item)
	i.heapIndex = h.Len()
	h.data = append(h.data, i)
}
func (h *maxHeap) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

func maxSlidingWindow(nums []int, k int) []int {
	maxHp := &maxHeap{}
	heap.Init(maxHp)
	l := len(nums)
	res := make([]int, l-k+1)
	var prev *Item
	var last *Item

	for i := 0; i < k; i++ {
		item := &Item{
			value: nums[i],
		}
		heap.Push(maxHp, item)
		if last == nil {
			last = item
		}
		if prev != nil {
			prev.next = item
		}
		prev = item
	}

	for i := k; i < l; i++ {
		res[i-k] = maxHp.data[0].value
		heap.Remove(maxHp, last.heapIndex)
		item := &Item{
			value: nums[i],
		}
		heap.Push(maxHp, item)
		prev.next = item
		prev = item
		last = last.next
	}

	res[l-k] = maxHp.data[0].value

	return res
}
