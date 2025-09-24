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
		res  bool
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 1},
			k:    3,
			res:  true,
		},
		{
			name: "Example 2",
			nums: []int{1, 0, 1, 1},
			k:    1,
			res:  true,
		},
		{
			name: "Example 3",
			nums: []int{1, 2, 3, 1, 2, 3},
			k:    2,
			res:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, containsNearbyDuplicate(tt.nums, tt.k), tt.name)
		})
	}
}
