package leetcode

// 3147. Taking Maximum Energy From the Mystic Dungeon
// https://leetcode.com/problems/taking-maximum-energy-from-the-mystic-dungeon/
func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	dp := make([]int, n)
	maxEnergy := -1 << 30

	for i := n - 1; i >= 0; i-- {
		if i+k >= n {
			dp[i] = energy[i]
		} else {
			dp[i] = energy[i] + dp[i+k]
		}
		maxEnergy = max(maxEnergy, dp[i])
	}

	return maxEnergy
}
