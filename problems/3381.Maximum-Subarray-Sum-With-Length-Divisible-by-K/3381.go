package leetcode

import "math"

// 3381. Maximum Subarray Sum With Length Divisible by K
// https://leetcode.com/problems/maximum-subarray-sum-with-length-divisible-by-k/
func maxSubarraySum(nums []int, k int) int64 {
	n := len(nums)
	prefixSum := int64(0)
	maxSum := int64(math.MinInt64)
	minByMod := make([]int64, k) //Min prefix sum by modulo k
	for i := 0; i < k; i++ {
		minByMod[i] = math.MaxInt64 / 2
	}
	minByMod[k-1] = 0

	for i := 0; i < n; i++ {
		prefixSum += int64(nums[i])
		if prefixSum-minByMod[i%k] > maxSum {
			maxSum = prefixSum - minByMod[i%k]
		}
		if prefixSum < minByMod[i%k] {
			minByMod[i%k] = prefixSum
		}
	}

	return maxSum
}
