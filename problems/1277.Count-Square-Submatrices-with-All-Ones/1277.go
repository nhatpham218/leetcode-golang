package leetcode

import (
	"fmt"
	"math"
)

func countSquares(matrix [][]int) int {
	rows := len(matrix)
	cols := len(matrix[0])

	dp := make([][]int, rows+1)
	for i := range dp {
		dp[i] = make([]int, cols+1)
	}
	res := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == 1 {
				dp[i+1][j+1] = min(dp[i][j+1], dp[i+1][j], dp[i][j]) + 1
				res += dp[i+1][j+1]
			}
		}
	}

	for i := 0; i <= rows; i++ {
		fmt.Println(dp[i])
	}

	return res
}

func min(a, b, c int) int {
	return int(math.Min(float64(a), math.Min(float64(b), float64(c))))
}
