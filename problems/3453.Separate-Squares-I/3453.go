package leetcode

import "math"

// 3453. Separate Squares I
// https://leetcode.com/problems/separate-squares-i/description/
func separateSquares(squares [][]int) float64 {
	totalArea := 0
	minY := math.MaxInt64
	maxY := math.MinInt64
	for _, square := range squares {
		totalArea += square[2] * square[2]
		minY = min(minY, square[1])
		maxY = max(maxY, square[1]+square[2])
	}

	// Helper function to calculate area below a given y line
	calcBottomArea := func(y float64) float64 {
		area := 0.0
		for _, square := range squares {
			bottom := float64(square[1])
			top := float64(square[1] + square[2])
			side := float64(square[2])
			if y <= bottom {
				// Line is below the square, no area below
				area += 0
			} else if y >= top {
				// Line is above the square, entire square is below
				area += side * side
			} else {
				// Line cuts through the square
				area += (y - bottom) * side
			}
		}
		return area
	}

	targetArea := float64(totalArea) / 2
	lo := float64(minY)
	hi := float64(maxY)

	for hi-lo > 1e-5 {
		mid := (lo + hi) / 2
		bottomArea := calcBottomArea(mid)
		if bottomArea < targetArea {
			lo = mid
		} else {
			hi = mid
		}
	}

	return lo
}
