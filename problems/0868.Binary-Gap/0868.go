package leetcode

import (
	"math"
	"strconv"
)

// 0868. Binary Gap
// https://leetcode.com/problems/binary-gap/description/
func binaryGap(n int) int {
	binary := strconv.FormatInt(int64(n), 2)
	maxGap := 0
	currentGap := math.MinInt32
	for i := range binary {
		if binary[i] == '1' {
			maxGap = max(maxGap, currentGap)
			currentGap = 1
		} else {
			currentGap++
		}
	}
	return maxGap
}
