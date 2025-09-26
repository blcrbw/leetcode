package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlidingMax(t *testing.T) {
	tests := []struct {
		name string
		nums [][]int
		res  int
	}{
		{
			name: "Example 1",
			nums: [][]int{[]int{2}, []int{3, 4}, []int{6, 5, 7}, []int{4, 1, 8, 3}},
			res:  11,
		},
		{
			name: "Example 2",
			nums: [][]int{[]int{-10}},
			res:  -10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, minimumTotal(tt.nums), tt.name)
		})
	}
}
