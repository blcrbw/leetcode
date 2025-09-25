package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlidingMax(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		res  []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			res:  []int{3, 3, 5, 5, 6, 7},
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4, 2, 3, 1, 4, 2},
			k:    3,
			res:  []int{3, 4, 4, 4, 3, 4, 4},
		},
		{
			name: "Example 3",
			nums: []int{1},
			k:    1,
			res:  []int{1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, maxSlidingWindow(tt.nums, tt.k), tt.name)
		})
	}
}
