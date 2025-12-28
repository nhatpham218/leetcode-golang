package leetcode

// 1351. Count Negative Numbers in a Sorted Matrix
// https://leetcode.com/problems/count-negative-numbers-in-a-sorted-matrix/description/
func countNegatives(grid [][]int) int {
	count := 0

	for i := len(grid) - 1; i >= 0; i-- {
		for j := len(grid[0]) - 1; j >= 0; j-- {
			if grid[i][j] < 0 {
				count++
			} else {
				break
			}
		}
	}

	return count
}
