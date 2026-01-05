package leetcode

import "math"

// 1975. Maximum Matrix Sum
// https://leetcode.com/problems/maximum-matrix-sum/description/
func maxMatrixSum(matrix [][]int) int64 {
	countNegatives := 0
	minNonZero := int64(math.MaxInt64)
	totalSum := int64(0)
	have0 := false
	for _, row := range matrix {
		for _, num := range row {
			if num == 0 {
				have0 = true
			} else {
				minNonZero = min(minNonZero, abs(int64(num)))
				totalSum += abs(int64(num))
				if num < 0 {
					countNegatives++
				}
			}
		}
	}
	if countNegatives%2 == 0 || have0 {
		return totalSum
	}
	return totalSum - minNonZero*2
}

func abs(a int64) int64 {
	if a < 0 {
		return -a
	}
	return a
}
