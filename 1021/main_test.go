package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPairs(t *testing.T) {
	tests := []struct {
		name string
		s    string
		res  string
	}{
		{
			name: "Example 1",
			s:    "(()())(())",
			res:  "()()()",
		},
		{
			name: "Example 2",
			s:    "(()())(())(()(()))",
			res:  "()()()()(())",
		},
		{
			name: "Example 3",
			s:    "()()",
			res:  "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, removeOuterParentheses(tt.s), tt.name)
		})
	}
}
