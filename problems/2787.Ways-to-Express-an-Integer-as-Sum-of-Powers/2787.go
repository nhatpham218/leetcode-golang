package leetcode

// 2787. Ways to Express an Integer as Sum of Powers
// https://leetcode.com/problems/ways-to-express-an-integer-as-sum-of-powers/

func numberOfWays(n int, x int) int {
	mod := 1000000007

	// dp[i] represents the number of ways to express i as sum of powers of x
	dp := make([]int, n+1)
	dp[0] = 1 // Base case: empty sum

	// Try each possible power of x
	for power := 1; power <= n; power++ {
		val := 1
		for i := 0; i < x; i++ {
			val *= power
			if val > n {
				break
			}
		}

		// If this power value exceeds n, we can stop
		if val > n {
			break
		}

		// Update dp array: for each value from n down to val
		for i := n; i >= val; i-- {
			dp[i] = (dp[i] + dp[i-val]) % mod
		}
	}

	return dp[n]
}
