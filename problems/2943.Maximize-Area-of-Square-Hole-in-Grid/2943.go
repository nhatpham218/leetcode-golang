package leetcode

import "sort"

// 2943. Maximize Area of Square Hole in Grid
// https://leetcode.com/problems/maximize-area-of-square-hole-in-grid/description/
func maximizeSquareHoleArea(n int, m int, hBars []int, vBars []int) int {
	maxGap := min(getMaxGap(hBars), getMaxGap(vBars))
	return maxGap * maxGap
}

func getMaxGap(bars []int) int {
	sort.Ints(bars)
	count, res := 2, 2
	for i := 1; i < len(bars); i++ {
		if bars[i-1]+1 == bars[i] {
			count++
		} else {
			count = 2
		}
		res = max(res, count)
	}
	return res
}
