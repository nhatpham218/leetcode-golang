package leetcode

// 3643. Flip Square Submatrix Vertically
// You are given an m x n integer matrix grid, and three integers x, y, and k.
// The integers x and y represent the row and column indices of the top-left corner of a square submatrix and the integer k represents the size (side length) of the square submatrix.
// Your task is to flip the submatrix by reversing the order of its rows vertically.
// Return the updated matrix.
func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	for i := range k / 2 {
		for j := range k {
			grid[x+i][y+j], grid[x+k-i-1][y+j] = grid[x+k-i-1][y+j], grid[x+i][y+j]
		}
	}
	return grid
}
