package leetcode

// 1895. Largest Magic Square
// https://leetcode.com/problems/largest-magic-square/description/
func largestMagicSquare(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	prefixSumRow := make([][]int, m)
	prefixSumCol := make([][]int, n)
	for i := 0; i < m; i++ {
		prefixSumRow[i] = make([]int, n)
		prefixSumRow[i][0] = grid[i][0]
		for j := 1; j < n; j++ {
			prefixSumRow[i][j] = prefixSumRow[i][j-1] + grid[i][j]
		}
	}
	for j := 0; j < n; j++ {
		prefixSumCol[j] = make([]int, m)
		prefixSumCol[j][0] = grid[0][j]
		for i := 1; i < m; i++ {
			prefixSumCol[j][i] = prefixSumCol[j][i-1] + grid[i][j]
		}
	}

	for k := min(m, n); k > 1; k-- {
		for i := 0; i <= m-k; i++ {
			for j := 0; j <= n-k; j++ {
				if isMagicSquare(grid, i, j, k, prefixSumRow, prefixSumCol) {
					return k
				}
			}
		}
	}
	return 1
}

func rowSum(prefixSumRow [][]int, row, colStart, colEnd int) int {
	if colStart == 0 {
		return prefixSumRow[row][colEnd]
	}
	return prefixSumRow[row][colEnd] - prefixSumRow[row][colStart-1]
}

func colSum(prefixSumCol [][]int, col, rowStart, rowEnd int) int {
	if rowStart == 0 {
		return prefixSumCol[col][rowEnd]
	}
	return prefixSumCol[col][rowEnd] - prefixSumCol[col][rowStart-1]
}

func isMagicSquare(grid [][]int, i, j, k int, prefixSumRow [][]int, prefixSumCol [][]int) bool {
	sum := rowSum(prefixSumRow, i, j, j+k-1)

	// Check all row sums
	for x := i; x < i+k; x++ {
		if rowSum(prefixSumRow, x, j, j+k-1) != sum {
			return false
		}
	}

	// Check all column sums
	for y := j; y < j+k; y++ {
		if colSum(prefixSumCol, y, i, i+k-1) != sum {
			return false
		}
	}

	// Check main diagonal (top-left to bottom-right)
	diag1 := 0
	for d := 0; d < k; d++ {
		diag1 += grid[i+d][j+d]
	}
	if diag1 != sum {
		return false
	}

	// Check anti-diagonal (top-right to bottom-left)
	diag2 := 0
	for d := 0; d < k; d++ {
		diag2 += grid[i+d][j+k-1-d]
	}
	if diag2 != sum {
		return false
	}

	return true
}
