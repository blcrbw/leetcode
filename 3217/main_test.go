package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLFU(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		head []int
		res  []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3},
			head: []int{1, 2, 3, 4, 5},
			res:  []int{4, 5},
		},
		{
			name: "Example 2",
			nums: []int{1},
			head: []int{1, 2, 1, 2, 1, 2},
			res:  []int{2, 2, 2},
		},
		{
			name: "Example 3",
			nums: []int{5},
			head: []int{1, 2, 3, 4},
			res:  []int{1, 2, 3, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, wrapper(tt.nums, tt.head), tt.name)
		})
	}
}

func wrapper(nums []int, head []int) []int {
	h := &ListNode{Val: head[0]}
	res := make([]int, 0, len(nums))
	prev := h

	for i := 1; i < len(head); i++ {
		n := &ListNode{Val: head[i]}
		prev.Next = n
		prev = n
	}

	m := modifiedList(nums, h)

	for {
		res = append(res, m.Val)
		if m.Next != nil {
			m = m.Next
		} else {
			break
		}
	}

	return res
}
