package leetcode

// 474. Ones and Zeroes
// https://leetcode.com/problems/ones-and-zeroes/
func findMaxForm(strs []string, m int, n int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for _, str := range strs {
		zeros, ones := 0, 0
		for _, c := range str {
			if c == '0' {
				zeros++
			} else {
				ones++
			}
		}

		for i := m; i >= zeros; i-- {
			for j := n; j >= ones; j-- {
				if dp[i-zeros][j-ones]+1 > dp[i][j] {
					dp[i][j] = dp[i-zeros][j-ones] + 1
				}
			}
		}
	}
	return dp[m][n]
}
