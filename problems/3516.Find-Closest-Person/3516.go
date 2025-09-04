package leetcode

import "math"

func findClosest(x int, y int, z int) int {
	dist1 := math.Abs(float64(x - z))
	dist2 := math.Abs(float64(y - z))

	if dist1 < dist2 {
		return 1
	} else if dist2 < dist1 {
		return 2
	}
	return 0
}
