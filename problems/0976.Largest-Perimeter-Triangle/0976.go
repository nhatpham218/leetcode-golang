package leetcode

import "sort"

// 976. Largest Perimeter Triangle
// https://leetcode.com/problems/largest-perimeter-triangle
func largestPerimeter(nums []int) int {
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] > nums[j]
	})

	for i := 0; i < len(nums)-2; i++ {
		a, b, c := nums[i], nums[i+1], nums[i+2]
		if b+c > a {
			return a + b + c
		}
	}

	return 0
}
