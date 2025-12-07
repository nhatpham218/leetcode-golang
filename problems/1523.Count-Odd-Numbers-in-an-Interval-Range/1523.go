package leetcode

// 1523. Count Odd Numbers in an Interval Range
// https://leetcode.com/problems/count-odd-numbers-in-an-interval-range/description/
func countOdds(low int, high int) int {
	if high%2 != 0 {
		high++
	}
	return (high - low + 1) / 2
}
