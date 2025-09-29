package leetcode

// 1039. Minimum Score Triangulation of Polygon
// https://leetcode.com/problems/minimum-score-triangulation-of-polygon
func minScoreTriangulation(values []int) int {
	n := len(values)
	if n < 3 {
		return 0
	}

	memo := make([][]int, n)
	for i := range memo {
		memo[i] = make([]int, n)
		for j := range memo[i] {
			memo[i][j] = -1
		}
	}

	var dp func(int, int) int
	dp = func(i, j int) int {
		if j-i < 2 {
			return 0
		}

		if memo[i][j] != -1 {
			return memo[i][j]
		}

		res := 1 << 31
		for k := i + 1; k < j; k++ {
			score := values[i]*values[j]*values[k] + dp(i, k) + dp(k, j)
			if score < res {
				res = score
			}
		}

		memo[i][j] = res
		return res
	}

	return dp(0, n-1)
}
