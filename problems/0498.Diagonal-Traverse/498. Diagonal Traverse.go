package leetcode

func findDiagonalOrder(matrix [][]int) []int {
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	row, col, dir, i, x, y, d := len(matrix), len(matrix[0]), [2][2]int{
		{-1, 1},
		{1, -1},
	}, 0, 0, 0, 0
	total := row * col
	res := make([]int, total)
	for i < total {
		for x >= 0 && x < row && y >= 0 && y < col {
			res[i] = matrix[x][y]
			i++
			x += dir[d][0]
			y += dir[d][1]
		}
		d = (d + 1) % 2
		if x == row {
			x--
			y += 2
		}
		if y == col {
			y--
			x += 2
		}
		if x < 0 {
			x = 0
		}
		if y < 0 {
			y = 0
		}
	}
	return res
}
