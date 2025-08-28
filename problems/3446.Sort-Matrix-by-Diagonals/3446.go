package leetcode

import (
	"slices"
)

// 3446. Sort Matrix by Diagonals
func sortMatrix(grid [][]int) [][]int {
	n := len(grid)

	if n < 2 {
		return grid
	}

	// Sort bottom-left triangle diagonals (including main diagonal) in decreasing order
	for i := 0; i < n; i++ {
		temp := make([]int, n-i)
		for j := 0; j < n-i; j++ {
			temp[j] = grid[i+j][j]
		}
		slices.Sort(temp)
		slices.Reverse(temp)
		for j := 0; j < n-i; j++ {
			grid[i+j][j] = temp[j]
		}
	}

	// Sort top-right triangle diagonals in increasing order
	for i := 1; i < n; i++ {
		temp := make([]int, n-i)
		for j := 0; j < n-i; j++ {
			temp[j] = grid[j][i+j]
		}
		slices.Sort(temp)
		for j := 0; j < n-i; j++ {
			grid[j][i+j] = temp[j]
		}
	}

	return grid
}
