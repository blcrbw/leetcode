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
			num:     []int{2, 2, -1, 3, -2, 2, 1, 1, 1, 0, -1},
			wantRes: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.wantRes, minimumPairRemoval(tt.num), tt.name)
		})
	}
}
