package leetcode

// 3461. Check If Digits Are Equal in String After Operations I
// https://leetcode.com/problems/check-if-digits-are-equal-in-string-after-operations-i/
func hasSameDigits(s string) bool {
	b := []byte(s)
	for i := len(b) - 1; i >= 2; i-- {
		for j := 0; j < i; j++ {
			b[j] = (b[j]+b[j+1])%10 + '0'
		}
	}
	return b[0] == b[1]
}
