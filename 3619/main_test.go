package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumPairRemoval(t *testing.T) {
	tests := []struct {
		name  string
		input [][]int
		k     int
		res   int
	}{
		{
			name:  "Example 1",
			input: [][]int{{0, 2, 1, 0, 0}, {0, 5, 0, 0, 5}, {0, 0, 1, 0, 0}, {0, 1, 4, 7, 0}, {0, 2, 0, 0, 8}},
			k:     5,
			res:   2,
		},
		{
			name:  "Example 2",
			input: [][]int{{3, 0, 3, 0}, {0, 3, 0, 3}, {3, 0, 3, 0}},
			k:     3,
			res:   6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, countIslands(tt.input, tt.k), tt.name)
		})
	}
}
