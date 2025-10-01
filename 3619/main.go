package main

// 3619. Count Islands With Total Value Divisible by K
// You are given an m x n matrix grid and a positive integer k. An island is a group of positive integers (representing
// land) that are 4-directionally connected (horizontally or vertically).
// The total value of an island is the sum of the values of all cells in the island.
// Return the number of islands with a total value divisible by k.

func countIslands(grid [][]int, k int) int {
	res := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] > 0 {
				sum := calcSum(i, j, grid)
				if divisible(sum, k) {
					res++
				}
			}
		}
	}

	return res
}

func calcSum(i, j int, grid [][]int) int {
	if grid[i][j] == 0 {
		return 0
	}
	sum := grid[i][j]
	grid[i][j] = 0

	if i > 0 {
		sum += calcSum(i-1, j, grid)
	}
	if i < (len(grid) - 1) {
		sum += calcSum(i+1, j, grid)
	}
	if j > 0 {
		sum += calcSum(i, j-1, grid)
	}
	if j < (len(grid[0]) - 1) {
		sum += calcSum(i, j+1, grid)
	}
	return sum
}

func divisible(x, k int) bool {
	return ((x / k) * k) == x
}
