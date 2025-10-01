package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumPairRemoval(t *testing.T) {
	tests := []struct {
		name string
		b    int
		e    int
		res  int
	}{
		{
			name: "Example 1",
			b:    9,
			e:    3,
			res:  13,
		},
		{
			name: "Example 2",
			b:    15,
			e:    4,
			res:  19,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, numWaterBottles(tt.b, tt.e), tt.name)
		})
	}
}
