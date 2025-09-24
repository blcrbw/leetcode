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
		res  float64
	}{
		{
			name: "Example 1",
			nums: []int{1, 12, -5, -6, 50, 3},
			k:    4,
			res:  12.75,
		},
		{
			name: "Example 2",
			nums: []int{5},
			k:    1,
			res:  5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, findMaxAverage(tt.nums, tt.k), tt.name)
		})
	}
}
