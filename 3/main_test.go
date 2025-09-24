package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairs(t *testing.T) {
	tests := []struct {
		name string
		inp  string
		res  int
	}{
		{
			name: "Example 1",
			inp:  "abcabcbb",
			res:  3,
		},
		{
			name: "Example 2",
			inp:  "bbbbb",
			res:  1,
		},
		{
			name: "Example 3",
			inp:  "pwwkew",
			res:  3,
		},
		{
			name: "Example 4",
			inp:  " ",
			res:  1,
		},
		{
			name: "Example 94",
			inp:  "cdd",
			res:  2,
		},
		{
			name: "Example 117",
			inp:  "nfpdmpi",
			res:  5,
		},
		{
			name: "Example 132",
			inp:  "uqinntq",
			res:  4,
		},
		{
			name: "Example 604",
			inp:  "asjrgapa",
			res:  6,
		},
		{
			name: "Example 922",
			inp:  "umvejcuuk",
			res:  6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, lengthOfLongestSubstring(tt.inp), tt.name)
		})
	}
}
