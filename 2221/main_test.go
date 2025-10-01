package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumPairRemoval(t *testing.T) {
	tests := []struct {
		name  string
		input []int
		res   int
	}{
		{
			name:  "Example 1",
			input: []int{1, 2, 3, 4, 5},
			res:   8,
		},
		{
			name:  "Example 2",
			input: []int{5},
			res:   5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, triangularSum(tt.input), tt.name)
		})
	}
}
