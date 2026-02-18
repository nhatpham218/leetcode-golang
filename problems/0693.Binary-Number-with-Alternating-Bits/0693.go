package leetcode

// 0693. Binary Number with Alternating Bits
// https://leetcode.com/problems/binary-number-with-alternating-bits/description/
func hasAlternatingBits(n int) bool {
	last := -1
	for n > 0 {
		last = n % 2
		n = n / 2
		if last == n%2 {
			return false
		}
		last = n % 2
	}
	return true
}
