package leetcode

import (
	"math"
	"sort"
)

// 3027. Find the Number of Ways to Place People II
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
		// Alice's position (upper left corner)
		ax, ay := points[i][0], points[i][1]
		maxY := -math.MaxInt64 // Use very small value as negative infinity

		for j := i + 1; j < n; j++ {
			// Bob's position (lower right corner)
			bx, by := points[j][0], points[j][1]

			// Check if Alice can be upper left and Bob lower right
			// Alice: upper left means ax <= bx and ay >= by
			if ax <= bx && ay >= by {
				// Check if this is a valid pair (no other points inside/on fence)
				// Since points are sorted by x, we only need to check y coordinate
				// We need by > maxY to ensure no point is between Alice and Bob in y-direction
				if by > maxY {
					result++
					maxY = by
				}
			}
		}
	}

	return result
}
