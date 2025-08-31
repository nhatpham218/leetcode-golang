# [37. Sudoku Solver](https://leetcode.com/problems/sudoku-solver/)

## Problem Description

Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

1. Each of the digits 1-9 must occur exactly once in each row.
2. Each of the digits 1-9 must occur exactly once in each column.
3. Each of the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.

The '.' character indicates empty cells.

## Example

**Input:**
```
board = [["5","3",".",".","7",".",".",".","."],
         ["6",".",".","1","9","5",".",".","."],
         [".","9","8",".",".",".",".","6","."],
         ["8",".",".",".","6",".",".",".","3"],
         ["4",".",".","8",".","3",".",".","1"],
         ["7",".",".",".","2",".",".",".","6"],
         [".","6",".",".",".",".","2","8","."],
         [".",".",".","4","1","9",".",".","5"],
         [".",".",".",".","8",".",".","7","9"]]
```

**Output:**
```
[["5","3","4","6","7","8","9","1","2"],
 ["6","7","2","1","9","5","3","4","8"],
 ["1","9","8","3","4","2","5","6","7"],
 ["8","5","9","7","6","1","4","2","3"],
 ["4","2","6","8","5","3","7","9","1"],
 ["7","1","3","9","2","4","8","5","6"],
 ["9","6","1","5","3","7","2","8","4"],
 ["2","8","7","4","1","9","6","3","5"],
 ["3","4","5","2","8","6","1","7","9"]]
```

## Constraints

- `board.length == 9`
- `board[i].length == 9`
- `board[i][j]` is a digit or '.'
- It is guaranteed that the input board has only one solution.

## Algorithm

The solution uses a **backtracking algorithm**:

1. **Find Empty Cell**: Scan the board to find the first empty cell (marked with '.').
2. **Try Digits**: For each digit from '1' to '9', check if placing it in the empty cell is valid.
3. **Validate Placement**: A digit is valid if it doesn't already exist in:
   - The same row
   - The same column  
   - The same 3x3 sub-box
4. **Recursive Solve**: If the digit is valid, place it and recursively solve the rest of the board.
5. **Backtrack**: If the recursive call fails, remove the digit and try the next one.
6. **Base Case**: If no empty cells remain, the puzzle is solved.

## Time Complexity

- **Time**: O(9^(n*n)) where n=9, but with pruning from validation, it's much faster in practice
- **Space**: O(n*n) for the recursion stack in the worst case

## Key Functions

- `solveSudoku(board [][]byte)`: Main function that modifies the board in-place
- `solve(board [][]byte) bool`: Recursive backtracking function
- `isValid(board [][]byte, row, col int, digit byte) bool`: Validates if a digit can be placed at a position

