package leetcode

import "sort"

// 3186. Maximum Total Damage With Spell Casting
// https://leetcode.com/problems/maximum-total-damage-with-spell-casting/
func maximumTotalDamage(power []int) int64 {
	// Group and sum spells by power value
	powerCount := make(map[int]int64)
	for _, p := range power {
		powerCount[p] += int64(p)
	}

	// Sort unique power values
	uniquePowers := make([]int, 0, len(powerCount))
	for p := range powerCount {
		uniquePowers = append(uniquePowers, p)
	}
	sort.Ints(uniquePowers)

	n := len(uniquePowers)
	if n == 0 {
		return 0
	}

	// dp[i] = max damage considering first i unique powers
	dp := make([]int64, n+1)

	for i := 1; i <= n; i++ {
		currentPower := uniquePowers[i-1]
		currentDamage := powerCount[currentPower]

		// Skip current power
		dp[i] = dp[i-1]

		// Take current power (find last valid index)
		j := i - 1
		for j > 0 && uniquePowers[j-1] >= currentPower-2 {
			j--
		}

		dp[i] = max(dp[i], dp[j]+currentDamage)
	}

	return dp[n]
}
