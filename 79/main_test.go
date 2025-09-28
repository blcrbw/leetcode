package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		num  [][]byte
		word string
		res  bool
	}{
		{
			name: "Example 1",
			num:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word: "ABCCED",
			res:  true,
		},
		{
			name: "Example 2",
			num:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word: "SEE",
			res:  true,
		},
		{
			name: "Example 3",
			num:  [][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}},
			word: "ABCB",
			res:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, exist(tt.num, tt.word), tt.name)
		})
	}
}
