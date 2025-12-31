package leetcode

// 1970. Last Day Where You Can Still Cross
// https://leetcode.com/problems/last-day-where-you-can-still-cross/description/
func latestDayToCross(row int, col int, cells [][]int) int {
	canCrossOnDay := func(day int) bool {
		grid := make([][]int, row)
		for i := range grid {
			grid[i] = make([]int, col)
		}
		for i := 0; i < day; i++ {
			grid[cells[i][0]-1][cells[i][1]-1] = 1
		}
		return canCross(grid)
	}

	lo, hi := 1, len(cells)
	for lo < hi {
		mid := (lo + hi + 1) / 2
		if canCrossOnDay(mid) {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	return lo
}

func canCross(grid [][]int) bool {
	cols := len(grid[0])
	for col := 0; col < cols; col++ {
		if grid[0][col] == 0 && dfs(grid, 0, col) {
			return true
		}
	}
	return false
}

func dfs(grid [][]int, r, c int) bool {
	if r == len(grid)-1 && grid[r][c] == 0 {
		return true
	}

	if r < 0 || r >= len(grid) || c < 0 || c >= len(grid[0]) || grid[r][c] > 0 {
		return false
	}

	grid[r][c] = 2
	check := dfs(grid, r+1, c) || dfs(grid, r-1, c) || dfs(grid, r, c+1) || dfs(grid, r, c-1)
	grid[r][c] = 0
	return check
}
