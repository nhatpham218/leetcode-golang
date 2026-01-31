package leetcode

// 66. Plus One
// https://leetcode.com/problems/plus-one/description/
func plusOne(digits []int) []int {
	remainder := 1
	for i := len(digits) - 1; i >= 0 && remainder == 1; i-- {
		digits[i] += 1
		digits[i] %= 10
		if digits[i] != 0 {
			remainder = 0
		}
	}
	if remainder == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}
