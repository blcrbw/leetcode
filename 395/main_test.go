package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLFU(t *testing.T) {
	tests := []struct {
		name string
		inp  string
		k    int
		res  int
	}{
		{
			name: "Example 1",
			inp:  "aaabb",
			k:    3,
			res:  3,
		},
		{
			name: "Example 2",
			inp:  "ababbc",
			k:    2,
			res:  5,
		},
		{
			name: "Example 3",
			inp:  "ababbcccc",
			k:    4,
			res:  4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, longestSubstring(tt.inp, tt.k), tt.name)
		})
	}
}
