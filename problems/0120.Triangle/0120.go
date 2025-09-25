package leetcode

// 120. Triangle
// https://leetcode.com/problems/triangle
func minimumTotal(triangle [][]int) int {
	n := len(triangle)
	if n == 0 {
		return 0
	}

	// dp[i] represents min path sum to reach current position
	dp := make([]int, n)

	// Initialize with last row
	for i := 0; i < n; i++ {
		dp[i] = triangle[n-1][i]
	}

	// Work bottom-up from second last row
	for row := n - 2; row >= 0; row-- {
		for col := 0; col <= row; col++ {
			// Min of two paths from next row + current value
			if dp[col] < dp[col+1] {
				dp[col] = triangle[row][col] + dp[col]
			} else {
				dp[col] = triangle[row][col] + dp[col+1]
			}
		}
	}

	return dp[0]
}
