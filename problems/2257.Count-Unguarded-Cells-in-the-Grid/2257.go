package leetcode

// 2257. Count Unguarded Cells in the Grid
// https://leetcode.com/problems/count-unguarded-cells-in-the-grid/
func countUnguarded(m int, n int, guards [][]int, walls [][]int) int {
	grid := make([][]int, m)
	for i := range grid {
		grid[i] = make([]int, n)
	}

	for _, wall := range walls {
		grid[wall[0]][wall[1]] = 2
	}

	for _, guard := range guards {
		grid[guard[0]][guard[1]] = 3
	}

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, guard := range guards {
		for _, dir := range dirs {
			r, c := guard[0]+dir[0], guard[1]+dir[1]
			for r >= 0 && r < m && c >= 0 && c < n {
				if grid[r][c] == 2 || grid[r][c] == 3 {
					break
				}
				grid[r][c] = 1
				r += dir[0]
				c += dir[1]
			}
		}
	}

	count := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 0 {
				count++
			}
		}
	}

	return count
}
