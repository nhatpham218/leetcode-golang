package leetcode

import "sort"

// 3397. Maximum Number of Distinct Elements After Operations
// https://leetcode.com/problems/maximum-number-of-distinct-elements-after-operations/
func maxDistinctElements(nums []int, k int) int {
	sort.Ints(nums)
	lastPicked := -1000000000000000000 // -10^18
	distinctCount := 0

	for _, num := range nums {
		lowerBound := num - k
		upperBound := num + k
		if lastPicked < lowerBound {
			lastPicked = lowerBound
		} else {
			lastPicked++
		}
		if lastPicked <= upperBound {
			distinctCount++
		} else {
			lastPicked--
		}
	}

	return distinctCount
}
