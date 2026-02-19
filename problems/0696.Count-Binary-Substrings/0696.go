package leetcode

// 0696. Count Binary Substrings
// https://leetcode.com/problems/count-binary-substrings/description/
func countBinarySubstrings(s string) int {
	count := 0
	prevCount := 0
	currentCount := 0
	for i := 0; i < len(s); i++ {
		if s[i] != s[i-1] {
			count += min(prevCount, currentCount)
			prevCount = currentCount
			currentCount = 1
		} else {
			currentCount++
		}
	}
	return count + min(prevCount, currentCount)
}
