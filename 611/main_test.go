package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSlidingMax(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "Example 1",
			nums: []int{2, 2, 3, 4},
			res:  3,
		},
		{
			name: "Example 2",
			nums: []int{2, 3, 4, 4},
			res:  4,
		},
		{
			name: "Example 2",
			nums: []int{2, 3, 4, 4, 7, 9, 18, 19, 30},
			res:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, triangleNumber(tt.nums), tt.name)
		})
	}
}
