package leetcode

// 3228. Maximum Number of Operations to Move Ones to the End
// https://leetcode.com/problems/maximum-number-of-operations-to-move-ones-to-the-end/
func maxOperations(s string) int {
	ans := 0
	count1 := 0
	for i, char := range s {
		if char == '1' {
			count1++
		}
		if char == '0' && i > 0 && s[i-1] == '1' {
			ans += count1
		}
	}
	return ans
}
