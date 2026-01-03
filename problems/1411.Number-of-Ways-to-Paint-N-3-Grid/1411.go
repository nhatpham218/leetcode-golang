package leetcode

// 1411. Number of Ways to Paint N × 3 Grid
// https://leetcode.com/problems/number-of-ways-to-paint-n-3-grid/description/
func numOfWays(n int) int {
	mod := 1000000007
	dp := make([][2]int, n+1)
	dp[1][0] = 6
	dp[1][1] = 6
	for i := 2; i <= n; i++ {
		dp[i][0] = (dp[i-1][0]*2 + dp[i-1][1]*2) % mod
		dp[i][1] = (dp[i-1][0]*2 + dp[i-1][1]*3) % mod
	}
	return (dp[n][0] + dp[n][1]) % mod
}
