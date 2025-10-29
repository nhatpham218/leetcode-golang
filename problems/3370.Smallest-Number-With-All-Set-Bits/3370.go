package leetcode

// 3370. Smallest Number With All Set Bits
// https://leetcode.com/problems/smallest-number-with-all-set-bits/
func smallestNumber(n int) int {
	result := 1
	for result < n {
		result = result*2 + 1
	}
	return result
}
