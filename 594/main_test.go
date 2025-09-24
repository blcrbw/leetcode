package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairs(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "Example 1",
			nums: []int{1, 3, 2, 2, 5, 2, 3, 7},
			res:  5,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 4},
			res:  2,
		},
		{
			name: "Example 3",
			nums: []int{1, 1, 1, 1},
			res:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, findLHS(tt.nums), tt.name)
		})
	}
}
