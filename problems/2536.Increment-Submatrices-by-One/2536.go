package leetcode

// 2536. Increment Submatrices by One
// https://leetcode.com/problems/increment-submatrices-by-one/
func rangeAddQueries(n int, queries [][]int) [][]int {
	mat := make([][]int, n)
	for i := range mat {
		mat[i] = make([]int, n)
	}

	for _, query := range queries {
		row1, col1 := query[0], query[1]
		row2, col2 := query[2], query[3]
		for i := row1; i <= row2; i++ {
			mat[i][col1]++
			if col2+1 < n {
				mat[i][col2+1]--
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 1; j < n; j++ {
			mat[i][j] += mat[i][j-1]
		}
	}

	return mat
}
