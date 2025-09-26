package leetcode

import "sort"

// 611. Valid Triangle Number
// https://leetcode.com/problems/valid-triangle-number
func triangleNumber(nums []int) int {
	if len(nums) < 3 {
		return 0
	}

	sort.Ints(nums)
	count := 0

	// For each potential largest side
	for i := 2; i < len(nums); i++ {
		left, right := 0, i-1

		// Use two pointers to find valid triangles
		for left < right {
			if nums[left]+nums[right] > nums[i] {
				// All pairs from left to right-1 with right form valid triangles
				count += right - left
				right--
			} else {
				left++
			}
		}
	}

	return count
}
