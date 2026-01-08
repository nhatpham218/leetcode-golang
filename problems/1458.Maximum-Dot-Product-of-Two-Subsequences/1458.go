package leetcode

import "math"

// 1458. Maximum Dot Product of Two Subsequences
// https://leetcode.com/problems/maximum-dot-product-of-two-subsequences/description/
func maxDotProduct(a []int, b []int) int {
	n, m := len(a), len(b)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
		for j := range dp[i] {
			dp[i][j] = math.MinInt32
		}
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			take := a[i-1]*b[j-1] + max(0, dp[i-1][j-1])
			dp[i][j] = max(take, max(dp[i-1][j], dp[i][j-1]))
		}
	}
	return dp[n][m]
}
