package leetcode

// 36. Valid Sudoku
// Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:
// Each row must contain the digits 1-9 without repetition.
// Each column must contain the digits 1-9 without repetition.
// Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.
func isValidSudoku(board [][]byte) bool {
	// check rows
	for i := range board {
		row := make(map[byte]bool)
		for j := range board[i] {
			if board[i][j] != '.' {
				if _, ok := row[board[i][j]]; ok {
					return false
				}
				row[board[i][j]] = true
			}
		}
	}

	// check columns
	for j := range board[0] {
		col := make(map[byte]bool)
		for i := range board {
			if board[i][j] != '.' {
				if _, ok := col[board[i][j]]; ok {
					return false
				}
				col[board[i][j]] = true
			}
		}
	}

	// check 3x3 sub-boxes
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			subBox := make(map[byte]bool)
			for k := range 3 {
				for l := range 3 {
					if board[i+k][j+l] != '.' {
						if _, ok := subBox[board[i+k][j+l]]; ok {
							return false
						}
						subBox[board[i+k][j+l]] = true
					}
				}
			}
		}
	}
	return true
}
