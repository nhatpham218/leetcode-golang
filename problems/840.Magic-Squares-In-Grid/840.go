package leetcode

// 840. Magic Squares In Grid
// https://leetcode.com/problems/magic-squares-in-grid/description/
func numMagicSquaresInside(grid [][]int) int {
	count := 0
	for i := 0; i < len(grid)-2; i++ {
		for j := 0; j < len(grid[0])-2; j++ {
			if grid[i+1][j+1] == 5 && isMagicSquare(grid, i, j) {
				count++
			}
		}
	}
	return count
}

func isMagicSquare(grid [][]int, r, c int) bool {
	nums := make([]int, 9)
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			v := grid[r+i][c+j]
			if v >= 1 && v <= 9 {
				nums[v-1]++
			} else {
				return false
			}
		}
	}
	for _, num := range nums {
		if num != 1 {
			return false
		}
	}

	sum := grid[r][c] + grid[r][c+1] + grid[r][c+2]
	for i := 0; i < 3; i++ {
		if grid[r+i][c]+grid[r+i][c+1]+grid[r+i][c+2] != sum {
			return false
		}
		if grid[r][c+i]+grid[r+1][c+i]+grid[r+2][c+i] != sum {
			return false
		}
	}
	if grid[r][c]+grid[r+1][c+1]+grid[r+2][c+2] != sum {
		return false
	}
	if grid[r][c+2]+grid[r+1][c+1]+grid[r+2][c] != sum {
		return false
	}
	return true
}
