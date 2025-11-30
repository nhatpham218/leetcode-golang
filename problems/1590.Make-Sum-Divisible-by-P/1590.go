package leetcode

// 1590. Make Sum Divisible by P
// https://leetcode.com/problems/make-sum-divisible-by-p/
func minSubarray(nums []int, p int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}
	targetMod := totalSum % p
	if targetMod == 0 {
		return 0
	}

	modIndex := map[int]int{0: -1}
	prefixMod := 0
	minLen := len(nums)

	for i, num := range nums {
		prefixMod = (prefixMod + num) % p
		searchMod := (prefixMod - targetMod + p) % p
		if prevIdx, ok := modIndex[searchMod]; ok {
			if i-prevIdx < minLen {
				minLen = i - prevIdx
			}
		}
		modIndex[prefixMod] = i
	}

	if minLen == len(nums) {
		return -1
	}
	return minLen
}
