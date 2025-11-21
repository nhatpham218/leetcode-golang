package leetcode

// 1930. Unique Length-3 Palindromic Subsequences
// https://leetcode.com/problems/unique-length-3-palindromic-subsequences/
func countPalindromicSubsequence(s string) int {
	firstPos := make([]int, 26)
	lastPos := make([]int, 26)

	for i := range firstPos {
		firstPos[i] = -1
		lastPos[i] = -1
	}

	for i := 0; i < len(s); i++ {
		charIdx := int(s[i] - 'a')
		if firstPos[charIdx] == -1 {
			firstPos[charIdx] = i
		}
		lastPos[charIdx] = i
	}

	count := 0
	for i := 0; i < 26; i++ {
		if firstPos[i] == -1 {
			continue
		}

		between := make(map[byte]bool)
		for j := firstPos[i] + 1; j < lastPos[i]; j++ {
			between[s[j]] = true
		}

		count += len(between)
	}

	return count
}
