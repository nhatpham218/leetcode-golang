# 36. Valid Sudoku

## Problem Description

Determine if a 9 x 9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

1. Each row must contain the digits 1-9 without repetition.
2. Each column must contain the digits 1-9 without repetition.
3. Each of the nine 3 x 3 sub-boxes of the grid must contain the digits 1-9 without repetition.

**Note:**
- A Sudoku board (partially filled) could be valid but is not necessarily solvable.
- Only the filled cells need to be validated according to the mentioned rules.

## Examples

### Example 1:
```
Input: board = 
[["5","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: true
```

### Example 2:
```
Input: board = 
[["8","3",".",".","7",".",".",".","."]
,["6",".",".","1","9","5",".",".","."]
,[".","9","8",".",".",".",".","6","."]
,["8",".",".",".","6",".",".",".","3"]
,["4",".",".","8",".","3",".",".","1"]
,["7",".",".",".","2",".",".",".","6"]
,[".","6",".",".",".",".","2","8","."]
,[".",".",".","4","1","9",".",".","5"]
,[".",".",".",".","8",".",".","7","9"]]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
```

## Constraints

- `board.length == 9`
- `board[i].length == 9`
- `board[i][j]` is a digit `1-9` or `'.'`.

## Algorithm Approach

To validate a Sudoku board, we need to check three conditions:

1. **Row validation**: Each row should not have duplicate digits (1-9)
2. **Column validation**: Each column should not have duplicate digits (1-9)  
3. **Sub-box validation**: Each 3x3 sub-box should not have duplicate digits (1-9)

**Key insights:**
- We only need to validate filled cells (ignore '.' cells)
- We can use hash sets to track seen digits in each row, column, and sub-box
- For 3x3 sub-boxes, we can calculate the box index as `(row/3)*3 + col/3`

**Approach:**
- Use three arrays of sets to track digits seen in rows, columns, and boxes
- Iterate through each cell once
- For each non-empty cell, check if the digit already exists in the corresponding row, column, or box
- If duplicate found, return false; otherwise add to sets
- If we complete the iteration without duplicates, return true

**Time Complexity:** O(1) - since the board size is fixed (9x9)  
**Space Complexity:** O(1) - constant space for the tracking sets
