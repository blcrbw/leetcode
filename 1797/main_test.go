package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthenticationManager(t *testing.T) {
	tests := []struct {
		name string
		cmds []string
		args []testArgs
		res  []int
	}{
		{
			name: "Example 1",
			cmds: []string{"AuthenticationManager", "renew", "generate", "countUnexpiredTokens", "generate", "renew", "renew", "countUnexpiredTokens"},
			args: []testArgs{testArgs{currentTime: 5}, testArgs{"aaa", 1}, testArgs{"aaa", 2}, testArgs{currentTime: 6}, testArgs{"bbb", 7}, testArgs{"aaa", 8}, testArgs{"bbb", 10}, testArgs{currentTime: 15}},
			res:  []int{0, 0, 0, 1, 0, 0, 0, 0},
		},
		//{
		//	name: "Example 24",
		//	cmds: []string{"LRUCache", "get", "get", "put", "get", "put", "put", "put", "put", "get", "put"},
		//	args: [][]int{[]int{1}, []int{6}, []int{8}, []int{12, 1}, []int{2}, []int{15, 11}, []int{5, 2}, []int{1, 15}, []int{4, 2}, []int{5}, []int{15, 15}},
		//	res:  []int{0, -1, -1, 0, -1, 0, 0, 0, 0, -1, 0},
		//},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			assert.Equal(t, tt.res, test(tt.cmds, tt.args, tt.res), tt.name)
		})
	}
}
