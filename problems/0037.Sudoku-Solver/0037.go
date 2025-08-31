package leetcode

// 37. Sudoku Solver
// Write a program to solve a Sudoku puzzle by filling the empty cells.
// A sudoku solution must satisfy all of the following rules:
// Each of the digits 1-9 must occur exactly once in each row.
// Each of the digits 1-9 must occur exactly once in each column.
// Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
func solveSudoku(board [][]byte) {
	solve(board)
}

func solve(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				for digit := byte('1'); digit <= '9'; digit++ {
					if isValid(board, i, j, digit) {
						board[i][j] = digit
						if solve(board) {
							return true
						}
						board[i][j] = '.'
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]byte, row, col int, digit byte) bool {
	// Check row
	for j := 0; j < 9; j++ {
		if board[row][j] == digit {
			return false
		}
	}

	// Check column
	for i := 0; i < 9; i++ {
		if board[i][col] == digit {
			return false
		}
	}

	// Check 3x3 box
	boxRow := (row / 3) * 3
	boxCol := (col / 3) * 3
	for i := boxRow; i < boxRow+3; i++ {
		for j := boxCol; j < boxCol+3; j++ {
			if board[i][j] == digit {
				return false
			}
		}
	}

	return true
}
