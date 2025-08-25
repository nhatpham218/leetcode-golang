package leetcode

func numberOfSubmatrices(grid [][]byte) int {
	rows := len(grid) + 1
	cols := len(grid[0]) + 1

	countX := make([][]int, rows)
	countY := make([][]int, rows)
	countX[0] = make([]int, cols)
	countY[0] = make([]int, cols)
	res := 0
	for i := 1; i < rows; i++ {

		countX[i] = make([]int, cols)
		countY[i] = make([]int, cols)
		// fmt.Println(countX[i], countY[i])

		for j := 1; j < cols; j++ {
			addX := shouldAdd(grid[i-1][j-1], 'X')
			addY := shouldAdd(grid[i-1][j-1], 'Y')
			countX[i][j] = countX[i-1][j] + countX[i][j-1] - countX[i-1][j-1] + int(addX)
			countY[i][j] = countY[i-1][j] + countY[i][j-1] - countY[i-1][j-1] + int(addY)
		}
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if countX[i][j] == countY[i][j] && countX[i][j] > 0 {
				res += 1
			}
		}
	}

	return res
}

func shouldAdd(char byte, checkChar byte) int {
	if char == checkChar {
		return 1
	}
	return 0
}
