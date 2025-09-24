package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name    string
		num     []int
		wantRes [][]int
	}{
		{
			name:    "Example 1",
			num:     []int{-1, 0, 1, 2, -1, -4},
			wantRes: [][]int{[]int{-1, -1, 2}, []int{-1, 0, 1}},
		},
		{
			name:    "Example 2",
			num:     []int{0, 1, 1},
			wantRes: [][]int{},
		},
		{
			name:    "Example 3",
			num:     []int{0, 0, 0},
			wantRes: [][]int{[]int{0, 0, 0}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.wantRes, threeSum(tt.num), tt.name)
		})
	}
}
