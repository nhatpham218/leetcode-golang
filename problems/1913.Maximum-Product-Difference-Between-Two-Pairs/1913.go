package leetcode

import "sort"

// 1913. Maximum Product Difference Between Two Pairs
// https://leetcode.com/problems/maximum-product-difference-between-two-pairs/description/
func maxProductDifference(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)-1]*nums[len(nums)-2] - nums[0]*nums[1]
}
