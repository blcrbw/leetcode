package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name    string
		num     []int
		target  int
		res     []int
		wantRes []int
	}{
		{
			name:    "Example 1",
			num:     []int{2, 7, 11, 15},
			target:  9,
			wantRes: []int{0, 1},
		},
		{
			name:    "Example 2",
			num:     []int{3, 2, 4},
			target:  6,
			wantRes: []int{1, 2},
		},
		{
			name:    "Example 3",
			num:     []int{3, 3},
			target:  6,
			wantRes: []int{0, 1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.wantRes, twoSum(tt.num, tt.target), tt.name)
		})
	}
}
