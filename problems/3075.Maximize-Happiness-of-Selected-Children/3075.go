package leetcode

import "sort"

// 3075. Maximize Happiness of Selected Children
// https://leetcode.com/problems/maximize-happiness-of-selected-children/description/
func maximumHappinessSum(happiness []int, k int) int64 {
	sort.Ints(happiness)
	totalHappiness := 0
	for i := 0; i < k; i++ {
		happiness[len(happiness)-i-1] -= i
		if happiness[len(happiness)-i-1] > 0 {
			totalHappiness += happiness[len(happiness)-i-1]
		}
	}
	return int64(totalHappiness)
}
