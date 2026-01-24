package leetcode

import "sort"

// 1877. Minimize Maximum Pair Sum in Array
// https://leetcode.com/problems/minimize-maximum-pair-sum-in-array/description/
func minPairSum(nums []int) int {
	sort.Ints(nums)
	n := len(nums)
	maxSum := 0
	for i := 0; i < n/2; i++ {
		maxSum = max(maxSum, nums[i]+nums[n-i-1])
	}
	return maxSum
}
