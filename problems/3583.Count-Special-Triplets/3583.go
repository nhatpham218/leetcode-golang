package leetcode

// 3583. Count Special Triplets
// https://leetcode.com/problems/count-special-triplets/description/
func specialTriplets(nums []int) int {
	const MOD = 1_000_000_007

	freqPrev := make(map[int]int)
	freqNext := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		freqNext[nums[i]]++
	}
	ans := 0

	for i := 0; i < len(nums); i++ {
		freqNext[nums[i]]--

		need := nums[i] * 2
		ans = (ans + freqPrev[need]*freqNext[need]) % MOD
		freqPrev[nums[i]]++
	}
	return ans
}
