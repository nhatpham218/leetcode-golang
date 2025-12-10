package leetcode

// 3577. Count the Number of Computer Unlocking Permutations
// https://leetcode.com/problems/count-the-number-of-computer-unlocking-permutations/description/
func countPermutations(complexity []int) int {
	const MOD = 1_000_000_007
	ans := 1
	for i := 1; i < len(complexity); i++ {
		if complexity[i] <= complexity[0] {
			return 0
		}
		ans = (ans * i) % MOD
	}
	return ans
}
