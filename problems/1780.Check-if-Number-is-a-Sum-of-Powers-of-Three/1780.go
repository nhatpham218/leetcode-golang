package main

// 1780. Check if Number is a Sum of Powers of Three
// https://leetcode.com/problems/check-if-number-is-a-sum-of-powers-of-three/
//
// Given an integer n, return true if it is possible to represent n as the sum of distinct powers of three.
// Otherwise, return false.
//
// An integer y is a power of three if there exists an integer x such that y == 3^x.
//
// Example 1:
// Input: n = 12
// Output: true
// Explanation: 12 = 3^1 + 3^2
//
// Example 2:
// Input: n = 91
// Output: true
// Explanation: 91 = 3^0 + 3^2 + 3^4
//
// Example 3:
// Input: n = 21
// Output: false

func checkPowersOfThree(n int) bool {
	// Dynamic programming approach
	// dp[i] represents whether number i can be represented as sum of distinct powers of 3

	// Generate powers of 3 up to n
	powers3 := []int{1}
	for power := 3; power <= n; power *= 3 {
		powers3 = append(powers3, power)
	}

	// Initialize dp array
	dp := make([]bool, n+1)
	dp[0] = true // Base case: 0 can be represented as empty sum

	// For each power of 3, try to use it to build numbers
	for _, power := range powers3 {
		// Iterate backwards to avoid using the same power multiple times
		for i := n; i >= power; i-- {
			if dp[i-power] {
				dp[i] = true
			}
		}
	}

	return dp[n]
}
