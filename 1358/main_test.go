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
			inp:  "abcabc",
			res:  10,
		},
		{
			name: "Example 2",
			inp:  "aaacb",
			res:  3,
		},
		{
			name: "Example 3",
			inp:  "abc",
			res:  1,
		},
		{
			name: "Example 4",
			inp:  "ab",
			res:  0,
		},
		{
			name: "Example 5",
			inp:  "",
			res:  0,
		},
		{
			name: "Example 6",
			inp:  "aaab",
			res:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, numberOfSubstrings(tt.inp), tt.name)
		})
	}
}
