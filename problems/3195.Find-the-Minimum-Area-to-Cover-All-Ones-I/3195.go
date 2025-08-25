package leetcode

import "math"

// 3195. Find the Minimum Area to Cover All Ones I
func minimumArea(grid [][]int) int {
	minX, minY, maxX, maxY := float64(len(grid)), float64(len(grid[0])), float64(0), float64(0)
	for x := 0; x < len(grid); x++ {
		for y := 0; y < len(grid[0]); y++ {
			if grid[x][y] == 1 {
				minX = math.Min(minX, float64(x))
				maxX = math.Max(maxX, float64(x))
				minY = math.Min(minY, float64(y))
				maxY = math.Max(maxY, float64(y))
			}
		}
	}
	return int((maxX - minX + 1) * (maxY - minY + 1))
}
