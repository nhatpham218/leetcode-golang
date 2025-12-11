package leetcode

// 3531. Count Covered Buildings
// https://leetcode.com/problems/count-covered-buildings/description/
func countCoveredBuildings(n int, buildings [][]int) int {
	maxXInRow := make([]int, n+1)
	minXInRow := make([]int, n+1)
	maxYInCol := make([]int, n+1)
	minYInCol := make([]int, n+1)

	for i := range minXInRow {
		minXInRow[i] = n + 1
		minYInCol[i] = n + 1
	}

	for _, building := range buildings {
		x, y := building[0], building[1]
		maxXInRow[y] = max(maxXInRow[y], x)
		minXInRow[y] = min(minXInRow[y], x)
		maxYInCol[x] = max(maxYInCol[x], y)
		minYInCol[x] = min(minYInCol[x], y)
	}

	coveredCount := 0
	for _, building := range buildings {
		x, y := building[0], building[1]
		if x > minXInRow[y] && x < maxXInRow[y] &&
			y > minYInCol[x] && y < maxYInCol[x] {
			coveredCount++
		}
	}

	return coveredCount
}
