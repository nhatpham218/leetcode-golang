package leetcode

// 2483. Minimum Penalty for a Shop
// https://leetcode.com/problems/minimum-penalty-for-a-shop/description/
func bestClosingTime(customers string) int {
	score, minScore, ans := 0, 0, 0
	for i, c := range customers {
		if c == 'N' {
			score++
		} else {
			score--
		}
		if score < minScore {
			minScore = score
			ans = i + 1
		}
	}
	return ans
}
