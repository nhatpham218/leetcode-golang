package leetcode

import "sort"

// 1984. Minimum Difference Between Highest and Lowest of K Scores
// https://leetcode.com/problems/minimum-difference-between-highest-and-lowest-of-k-scores/description/
func minimumDifference(nums []int, k int) int {
	sort.Ints(nums)
	minDiff := int(1e9)
	for i := 0; i <= len(nums)-k; i++ {
		minDiff = min(minDiff, nums[i+k-1]-nums[i])
	}
	return minDiff
}
