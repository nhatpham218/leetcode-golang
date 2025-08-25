# 3212. Count Submatrices With Equal Frequency of X and Y

## Problem Description

Given a 2D character matrix `grid`, where `grid[i][j]` is either 'X', 'Y', or '.', return the number of submatrices that contain:

- `grid[0][0]`
- an equal frequency of 'X' and 'Y'.
- at least one 'X'.

## Examples

### Example 1:
```
Input: grid = [["X","Y","."],["Y",".","."]]
Output: 3
```

### Example 2:
```
Input: grid = [["X","X"],["X","Y"]]
Output: 0
```
**Explanation:** No submatrix has an equal frequency of 'X' and 'Y'.

### Example 3:
```
Input: grid = [[".","."],[".","."]]
Output: 0
```
**Explanation:** No submatrix has at least one 'X'.

## Constraints

- `1 <= grid.length, grid[i].length <= 1000`
- `grid[i][j]` is either 'X', 'Y', or '.'.

## Solution

The solution function `numberOfSubmatrices` is left blank for you to implement.
