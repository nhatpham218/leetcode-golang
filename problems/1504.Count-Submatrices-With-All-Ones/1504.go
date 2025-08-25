package leetcode

import "math"

func numSubmat(mat [][]int) int {
	rows := len(mat)
	cols := len(mat[0])

	res := 0

	for i := 1; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if mat[i][j] >= 1 {
				mat[i][j] = mat[i-1][j] + 1
			}
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			minHeight := float64(mat[i][j])
			for k := j; k >= 0; k-- {
				minHeight = math.Min(minHeight, float64(mat[i][k]))
				res += int(minHeight)
			}
		}
	}

	return res
}
