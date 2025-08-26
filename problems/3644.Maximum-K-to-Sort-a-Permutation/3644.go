package leetcode

import (
	"math"

	"github.com/samber/lo"
)

// 3644. Maximum K to Sort a Permutation
// You are given an integer array nums of length n, where nums is a permutation of the numbers in the range [0..n - 1].
// You may swap elements at indices i and j only if nums[i] AND nums[j] == k, where AND denotes the bitwise AND operation and k is a non-negative integer.
// Return the maximum value of k such that the array can be sorted in non-decreasing order using any number of such swaps. If nums is already sorted, return 0.
func sortPermutation(nums []int) int {
	mask := math.MaxInt
	for i, number := range nums {
		if number != i {
			mask = mask & number
		}
	}
	return lo.Ternary(mask == math.MaxInt, 0, mask)
}
