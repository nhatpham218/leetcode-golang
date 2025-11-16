package leetcode

// 1513. Number of Substrings With Only 1s
// https://leetcode.com/problems/number-of-substrings-with-only-1s/
func numSub(s string) int {
	count1 := 0
	ans := 0
	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			count1++
		}
		if s[i] == '0' || i == len(s)-1 {
			ans += count1 * (count1 + 1) / 2
			count1 = 0
		}
	}
	return ans % (1e9 + 7)
}
