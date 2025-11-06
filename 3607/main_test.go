package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLFU(t *testing.T) {
	tests := []struct {
		name    string
		c       int
		conn    [][]int
		queries [][]int
		res     []int
	}{
		{
			name:    "Example 1",
			c:       5,
			conn:    [][]int{[]int{1, 2}, []int{2, 3}, []int{3, 4}, []int{4, 5}},
			queries: [][]int{[]int{1, 3}, []int{2, 1}, []int{1, 1}, []int{2, 2}, []int{1, 2}},
			res:     []int{3, 2, 3},
		},
		{
			name:    "Example 2",
			c:       3,
			conn:    [][]int{},
			queries: [][]int{[]int{1, 1}, []int{2, 1}, []int{1, 1}},
			res:     []int{1, -1},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, processQueries(tt.c, tt.conn, tt.queries), tt.name)
		})
	}
}
