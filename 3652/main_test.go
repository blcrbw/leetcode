package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumPairRemoval(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		strategy []int
		k        int
		res      int64
	}{
		{
			name:     "Example 129",
			prices:   []int{4, 7, 13},
			strategy: []int{-1, -1, 0},
			k:        2,
			res:      9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, maxProfit(tt.prices, tt.strategy, tt.k), tt.name)
		})
	}
}
