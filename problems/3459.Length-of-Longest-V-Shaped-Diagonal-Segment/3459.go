package leetcode

// 3459. Length of Longest V-Shaped Diagonal Segment
func lenOfVDiagonal(grid [][]int) int {
	// Four diagonal directions: down-right, down-left, up-left, up-right
	dirs := [4][2]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

	rows, cols := len(grid), len(grid[0])

	// Memoization: dpCache[row][col][direction][turn_used]
	dpCache := make([][][4][2]int, rows)
	for row := range dpCache {
		dpCache[row] = make([][4][2]int, cols)
		for col := range dpCache[row] {
			for dirIdx := range dpCache[row][col] {
				for turnUsed := range dpCache[row][col][dirIdx] {
					dpCache[row][col][dirIdx][turnUsed] = -1
				}
			}
		}
	}

	var dfs func(currRow, currCol, dirIdx int, canTurn bool, expected int) int
	dfs = func(currRow, currCol, dirIdx int, canTurn bool, expected int) int {
		nextRow := currRow + dirs[dirIdx][0]
		nextCol := currCol + dirs[dirIdx][1]

		// Boundary and value check
		if nextRow < 0 || nextCol < 0 || nextRow >= rows || nextCol >= cols || grid[nextRow][nextCol] != expected {
			return 0
		}

		turnIdx := 0
		if canTurn {
			turnIdx = 1
		}

		// Check memoization
		if dpCache[nextRow][nextCol][dirIdx][turnIdx] != -1 {
			return dpCache[nextRow][nextCol][dirIdx][turnIdx]
		}

		// Continue in same direction
		maxLen := dfs(nextRow, nextCol, dirIdx, canTurn, 2-expected)

		// Try clockwise turn if available
		if canTurn {
			clockwiseDir := (dirIdx + 1) % 4
			maxLen = max(maxLen, dfs(nextRow, nextCol, clockwiseDir, false, 2-expected))
		}

		dpCache[nextRow][nextCol][dirIdx][turnIdx] = maxLen + 1
		return maxLen + 1
	}

	maxResult := 0
	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == 1 {
				for dirIdx := 0; dirIdx < 4; dirIdx++ {
					segmentLen := dfs(row, col, dirIdx, true, 2) + 1
					maxResult = max(maxResult, segmentLen)
				}
			}
		}
	}
	return maxResult
}
