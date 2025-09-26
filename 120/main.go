package main

// 120. Triangle
// Given a triangle array, return the minimum path sum from top to bottom.
// For each step, you may move to an adjacent number of the row below. More formally, if you are on index i on the
// current row, you may move to either index i or index i + 1 on the next row.

func minimumTotal(triangle [][]int) int {
	levels := len(triangle)
	// skip lowest level
	for i := levels - 2; i >= 0; i-- {
		// for each element on the current level we have to sum up the triangle[i][j] value with the min of lower level
		// options triangle[i+1][j] and triangle[i+1][j+1]
		for j := 0; j <= i; j++ {
			if triangle[i+1][j] < triangle[i+1][j+1] {
				triangle[i][j] += triangle[i+1][j]
			} else {
				triangle[i][j] += triangle[i+1][j+1]
			}
		}
		// and then check upper level
	}

	// the top element has the minimum path sum.
	return triangle[0][0]
}
