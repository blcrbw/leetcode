package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairs(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		res  []float64
	}{
		{
			name: "Example 1",
			nums: []int{1, 3, -1, -3, 5, 3, 6, 7},
			k:    3,
			res:  []float64{1, -1, -1, 3, 5, 6},
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4, 2, 3, 1, 4, 2},
			k:    3,
			res:  []float64{2, 3, 3, 3, 2, 3, 2},
		},
		{
			name: "Example 3",
			nums: []int{7, 9, 3, 8, 0, 2, 4, 8, 3, 9},
			k:    1,
			res:  []float64{7, 9, 3, 8, 0, 2, 4, 8, 3, 9},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, medianSlidingWindow(tt.nums, tt.k), tt.name)
		})
	}
}
