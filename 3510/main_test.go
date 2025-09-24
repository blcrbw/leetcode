package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumPairRemoval(t *testing.T) {
	tests := []struct {
		name    string
		num     []int
		wantRes int
	}{
		{
			name:    "Example 1",
			num:     []int{5, 2, 3, 1},
			wantRes: 2,
		},
		{
			name:    "Example 2",
			num:     []int{1, 2, 2},
			wantRes: 0,
		},
		{
			name:    "Example 3",
			num:     []int{-2, 1, 2, -1, -1, -2, -2, -1, -1, 1, 1},
			wantRes: 10,
		},
		{
			name:    "Example 6",
			num:     []int{-7, -2, -4, 4, 8, -6, 0, 0, 4, 5, 1, -8},
			wantRes: 11,
		},
		{
			name:    "Example 7",
			num:     []int{0, -1, -1, -1, 1, 0, -1},
			wantRes: 5,
		},
		{
			name:    "Example 8",
			num:     []int{3, 6, 4, -6, 2, -4, 5, -7, -3, 6, 3, -4},
			wantRes: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.wantRes, minimumPairRemoval(tt.num), tt.name)
		})
	}
}
