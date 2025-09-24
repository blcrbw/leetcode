package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRU(t *testing.T) {
	tests := []struct {
		name string
		cmds []string
		args [][]int
		res  []int
	}{
		{
			name: "Example 1",
			cmds: []string{"LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"},
			args: [][]int{[]int{2}, []int{1, 1}, []int{2, 2}, []int{1}, []int{3, 3}, []int{2}, []int{4, 4}, []int{1}, []int{3}, []int{4}},
			res:  []int{0, 0, 0, 1, 0, -1, 0, -1, 3, 4},
		},
		{
			name: "Example 24",
			cmds: []string{"LRUCache", "get", "get", "put", "get", "put", "put", "put", "put", "get", "put"},
			args: [][]int{[]int{1}, []int{6}, []int{8}, []int{12, 1}, []int{2}, []int{15, 11}, []int{5, 2}, []int{1, 15}, []int{4, 2}, []int{5}, []int{15, 15}},
			res:  []int{0, -1, -1, 0, -1, 0, 0, 0, 0, -1, 0},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, test(tt.cmds, tt.args, tt.res), tt.name)
		})
	}
}
