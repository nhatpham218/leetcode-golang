package leetcode

// 3713. Longest Balanced Substring I
// https://leetcode.com/problems/longest-balanced-substring-i/description/
func longestBalanced(s string) int {
	n := len(s)
	best := 1
	var freq [26]int
	for left := 0; left < n; left++ {
		freq = [26]int{}
		distinct, maxFreq, atMax := 0, 0, 0
		for right := left; right < n; right++ {
			idx := s[right] - 'a'
			freq[idx]++
			cnt := freq[idx]
			if cnt == 1 {
				distinct++
			}
			if cnt > maxFreq {
				maxFreq, atMax = cnt, 1
			} else if cnt == maxFreq {
				atMax++
			}
			if distinct == atMax {
				best = max(best, right-left+1)
			}
		}
	}
	return best
}
