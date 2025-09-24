package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLFU(t *testing.T) {
	tests := []struct {
		name string
		args []int
		res  []int
	}{
		{
			name: "Example 1",
			args: []int{1, 2, 3, 4, 5, 6, -1, -1, -1, 7, 8, 9, 10, -1, -1, 11, 12},
			res:  []int{1, 2, 3, 7, 8, 11, 12, 9, 10, 4, 5, 6},
		},
		{
			name: "Example 2",
			args: []int{1, -1, 2, -1, 3},
			res:  []int{1, 2, 3},
		},
		{
			name: "Example 3",
			args: []int{1, 2, -1, 3},
			res:  []int{1, 3, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, test(tt.args), tt.name)
		})
	}
}
