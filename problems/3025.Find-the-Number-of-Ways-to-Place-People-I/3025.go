package leetcode

import "sort"

// 3025. Find the Number of Ways to Place People I
func numberOfPairs(points [][]int) int {
	// Sort points by x ascending, then by y descending
	sort.Slice(points, func(i, j int) bool {
		if points[i][0] == points[j][0] {
			return points[i][1] > points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	n := len(points)
	result := 0

	for i := 0; i < n; i++ {
		top := points[i][1]
		bot := -1 // Use -1 as negative infinity since y >= 0

		for j := i + 1; j < n; j++ {
			y := points[j][1]
			if bot < y && y <= top {
				result++
				bot = y
				if bot == top {
					break
				}
			}
		}
	}

	return result
}
