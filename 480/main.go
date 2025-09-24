package main

import (
	"container/heap"
)

// 480. Sliding Window Median
// The median is the middle value in an ordered integer list. If the size of the list is even, there is no middle value.
// So the median is the mean of the two middle values.
//    For examples, if arr = [2,3,4], the median is 3.
//    For examples, if arr = [1,2,3,4], the median is (2 + 3) / 2 = 2.5.
// You are given an integer array nums and an integer k. There is a sliding window of size k which is moving from the
// very left of the array to the very right. You can only see the k numbers in the window. Each time the sliding window
// moves right by one position.
// Return the median array for each window in the original array. Answers within 10-5 of the actual value will be accepted.

type Item struct {
	value     int
	next      *Item
	heapIndex int
	isLowHeap bool
}

type minHeap struct {
	data []*Item
}

func (h *minHeap) Len() int {
	return len(h.data)
}
func (h *minHeap) Less(i, j int) bool {
	return h.data[i].value < h.data[j].value
}
func (h *minHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
	h.data[i].heapIndex, h.data[j].heapIndex = i, j
}
func (h *minHeap) Push(x interface{}) {
	i := x.(*Item)
	i.heapIndex = h.Len()
	i.isLowHeap = false
	h.data = append(h.data, i)
}
func (h *minHeap) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
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
	i.isLowHeap = true
	h.data = append(h.data, i)
}
func (h *maxHeap) Pop() interface{} {
	old := h.data
	n := len(old)
	x := old[n-1]
	h.data = old[0 : n-1]
	return x
}

type Manipulator struct {
	lowHeap  *maxHeap
	highHeap *minHeap
}

func NewManipulator() *Manipulator {
	lowHeap := &maxHeap{}
	highHeap := &minHeap{}
	heap.Init(lowHeap)
	heap.Init(highHeap)
	return &Manipulator{lowHeap, highHeap}
}

func (m *Manipulator) Add(n int) *Item {
	i := &Item{value: n}
	if m.lowHeap.Len() == 0 || m.lowHeap.data[0].value >= n {
		heap.Push(m.lowHeap, i)
	} else {
		heap.Push(m.highHeap, i)
	}

	m.balance()
	return i
}

func (m *Manipulator) RemoveItem(i *Item) {
	if i.isLowHeap {
		heap.Remove(m.lowHeap, i.heapIndex)
	} else {
		heap.Remove(m.highHeap, i.heapIndex)
	}
	m.balance()
}

func (m *Manipulator) balance() {
	if m.lowHeap.Len() > m.highHeap.Len()+1 {
		heap.Push(m.highHeap, heap.Pop(m.lowHeap))
	} else if m.lowHeap.Len() < m.highHeap.Len() {
		heap.Push(m.lowHeap, heap.Pop(m.highHeap))
	}
}

func (m *Manipulator) GetMedian() float64 {
	if m.lowHeap.Len() == 0 {
		return 0
	} else if m.highHeap.Len() == m.lowHeap.Len() {
		return float64(m.lowHeap.data[0].value+m.highHeap.data[0].value) / 2.0
	} else {
		return float64(m.lowHeap.data[0].value)
	}
}

func medianSlidingWindow(nums []int, k int) []float64 {
	m := NewManipulator()
	l := len(nums)
	res := make([]float64, l-k+1)
	var prev *Item
	var last *Item

	for i := 0; i < k; i++ {
		item := m.Add(nums[i])
		if last == nil {
			last = item
		}
		if prev != nil {
			prev.next = item
		}
		prev = item
	}

	for i := k; i < l; i++ {
		res[i-k] = m.GetMedian()
		m.RemoveItem(last)
		item := m.Add(nums[i])
		prev.next = item
		prev = item
		last = last.next
	}

	res[l-k] = m.GetMedian()

	return res
}
