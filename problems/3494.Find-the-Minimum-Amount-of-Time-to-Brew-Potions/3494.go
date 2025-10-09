package leetcode

// 3494. Find the Minimum Amount of Time to Brew Potions
// https://leetcode.com/problems/find-the-minimum-amount-of-time-to-brew-potions/
func minTime(skill []int, mana []int) int64 {
	n := len(skill)
	// f[i] tracks the earliest free time for wizard i (after their last potion)
	f := make([]int64, n)

	// Process each potion
	for _, x := range mana {
		// Start from wizard 0's earliest free time
		now := f[0]

		// Calculate when each wizard can work on this potion
		for i := 1; i < n; i++ {
			// Wizard i starts either when wizard i-1 finishes this potion,
			// or when wizard i is free, whichever is later
			now = max(now+int64(skill[i-1])*int64(x), f[i])
		}

		// Last wizard finishes at this time
		f[n-1] = now + int64(skill[n-1])*int64(x)

		// Update all other f values in reverse order
		for i := n - 2; i >= 0; i-- {
			f[i] = f[i+1] - int64(skill[i+1])*int64(x)
		}
	}

	return f[n-1]
}
