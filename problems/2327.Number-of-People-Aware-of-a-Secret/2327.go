package leetcode

func peopleAwareOfSecret(n int, delay int, forget int) int {
	const MOD = 1000000007

	// dp[i][j] = number of people who have known the secret for exactly j+1 days at day i
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, forget)
	}

	// Day 1: one person discovers the secret (knows for 1 day)
	dp[1][0] = 1

	for i := 2; i <= n; i++ {
		// People who knew for j days yesterday now know for j+1 days
		for j := 1; j < forget; j++ {
			dp[i][j] = dp[i-1][j-1]
		}

		// New people learn the secret from those who can share
		// Can share from day delay to day forget-1 (inclusive)
		for j := delay - 1; j < forget-1; j++ {
			dp[i][0] = (dp[i][0] + dp[i-1][j]) % MOD
		}
	}

	// Count all people who know the secret at day n
	total := 0
	for j := 0; j < forget; j++ {
		total = (total + dp[n][j]) % MOD
	}

	return total
}
