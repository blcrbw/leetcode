package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticationManager(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		res  int
	}{
		{
			name: "Example 1",
			nums: []int{4, 2, 5, 3},
			res:  0,
		},
		{
			name: "Example 2",
			nums: []int{1, 2, 3, 5, 6},
			res:  1,
		},
		{
			name: "Example 3",
			nums: []int{1, 10, 100, 1000},
			res:  3,
		},
		{
			name: "Example 26",
			nums: []int{8, 5, 9, 9, 8, 4},
			res:  2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, minOperations(tt.nums), tt.name)
		})
	}
}
