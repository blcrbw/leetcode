package main

// 79. Word Search
// Given an m x n grid of characters board and a string word, return true if word exists in the grid.
// The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or
// vertically neighboring. The same letter cell may not be used more than once.

func exist(board [][]byte, word string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if board[i][j] == word[0] && dfs(board, i, j, 0, word) {
				return true
			}
		}
	}
	return false
}

func dfs(board [][]byte, i int, j int, countFound int, word string) bool {
	if countFound == len(word) {
		return true
	}

	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[countFound] {
		return false
	}
	tmp := board[i][j]
	board[i][j] = '#'
	res := dfs(board, i-1, j, countFound+1, word) || dfs(board, i+1, j, countFound+1, word) || dfs(board, i, j-1, countFound+1, word) || dfs(board, i, j+1, countFound+1, word)
	board[i][j] = tmp
	return res
}
