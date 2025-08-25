package leetcode

import "fmt"

type point struct {
	x, y int
}

type rectangle struct {
	topLeft, bottomRight point
}

// 3197. Find the Minimum Area to Cover All Ones II
func minimumSum(grid [][]int) int {
	rows, cols := len(grid), len(grid[0])
	minTotal := rows * cols

	// try split 1 time by vertical line
	for x := 0; x < rows-1; x++ {
		rec1 := rectangle{point{0, 0}, point{x, cols - 1}}
		rec2 := rectangle{point{x + 1, 0}, point{rows - 1, cols - 1}}
		rec21, rec22 := split(grid, rec2)
		total := calculateBoundingArea(grid, rec1) + calculateBoundingArea(grid, rec21) + calculateBoundingArea(grid, rec22)
		if total < minTotal {
			minTotal = total
		}

		rec11, rec12 := split(grid, rec1)
		total = calculateBoundingArea(grid, rec11) + calculateBoundingArea(grid, rec12) + calculateBoundingArea(grid, rec2)
		if total < minTotal {
			minTotal = total
		}
	}
	// try split 1 time by horizontal line
	for y := 0; y < cols-1; y++ {
		rec1 := rectangle{point{0, 0}, point{rows - 1, y}}
		rec2 := rectangle{point{0, y + 1}, point{rows - 1, cols - 1}}
		rec21, rec22 := split(grid, rec2)
		total := calculateBoundingArea(grid, rec1) + calculateBoundingArea(grid, rec21) + calculateBoundingArea(grid, rec22)
		if total < minTotal {
			minTotal = total
		}

		rec11, rec12 := split(grid, rec1)
		total = calculateBoundingArea(grid, rec11) + calculateBoundingArea(grid, rec12) + calculateBoundingArea(grid, rec2)
		if total < minTotal {
			minTotal = total
		}
	}

	return minTotal
}

func split(grid [][]int, rec rectangle) (res1 rectangle, res2 rectangle) {
	minTotal := calculateBoundingArea(grid, rec)
	fmt.Println("minTotal", minTotal)
	fmt.Println("rec", rec)
	// split by horizontal line
	for y := rec.topLeft.y; y < rec.bottomRight.y; y++ {
		rec1 := rectangle{point{rec.topLeft.x, rec.topLeft.y}, point{rec.bottomRight.x, y}}
		rec2 := rectangle{point{rec.topLeft.x, y + 1}, point{rec.bottomRight.x, rec.bottomRight.y}}
		total := calculateBoundingArea(grid, rec1) + calculateBoundingArea(grid, rec2)
		if total <= minTotal && total > 0 {
			minTotal = total
			res1 = rec1
			res2 = rec2
		}
	}

	// split by vertical line
	for x := rec.topLeft.x; x < rec.bottomRight.x; x++ {
		rec1 := rectangle{point{rec.topLeft.x, rec.topLeft.y}, point{x, rec.bottomRight.y}}
		rec2 := rectangle{point{x + 1, rec.topLeft.y}, point{rec.bottomRight.x, rec.bottomRight.y}}
		total := calculateBoundingArea(grid, rec1) + calculateBoundingArea(grid, rec2)
		if total <= minTotal && total > 0 {
			minTotal = total
			res1 = rec1
			res2 = rec2

		}
	}
	fmt.Println("res1", res1, "res2", res2, "minTotal", minTotal)

	return res1, res2
}

// calculateBoundingArea calculates the minimum area of rectangle needed to cover all positions
func calculateBoundingArea(grid [][]int, rec rectangle) int {
	checked := false
	minX, minY, maxX, maxY := rec.bottomRight.x, rec.bottomRight.y, rec.topLeft.x, rec.topLeft.y
	for x := rec.topLeft.x; x <= rec.bottomRight.x; x++ {
		for y := rec.topLeft.y; y <= rec.bottomRight.y; y++ {
			if grid[x][y] == 1 {
				minX = min(minX, x)
				maxX = max(maxX, x)
				minY = min(minY, y)
				maxY = max(maxY, y)
				checked = true
			}
		}
	}

	if !checked {
		return 999999999999999999
	}

	return (maxX - minX + 1) * (maxY - minY + 1)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
