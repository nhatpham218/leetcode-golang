package leetcode

// 3459. Length of Longest V-Shaped Diagonal Segment
func lenOfVDiagonal(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	// Direction vectors: {1,1}, {1,-1}, {-1,-1}, {-1,1}
	// Representing: down-right, down-left, up-left, up-right
	dirs := [][]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

	// Memoization table: [row][col][direction][hasMadeTurn]
	memo := make([][][][]int, m)
	for i := range memo {
		memo[i] = make([][][]int, n)
		for j := range memo[i] {
			memo[i][j] = make([][]int, 4)
			for k := range memo[i][j] {
				memo[i][j][k] = make([]int, 2)
				for l := range memo[i][j][k] {
					memo[i][j][k][l] = -1
				}
			}
		}
	}

	var dfs func(int, int, int, bool, int) int
	dfs = func(cx, cy, direction int, turn bool, target int) int {
		nx := cx + dirs[direction][0]
		ny := cy + dirs[direction][1]

		// Check boundaries and target value
		if nx < 0 || ny < 0 || nx >= m || ny >= n || grid[nx][ny] != target {
			return 0
		}

		turnIndex := 0
		if turn {
			turnIndex = 1
		}

		// Check memoization
		if memo[nx][ny][direction][turnIndex] != -1 {
			return memo[nx][ny][direction][turnIndex]
		}

		// Continue walking in the original direction
		maxStep := dfs(nx, ny, direction, turn, 2-target)

		// If we haven't made a turn yet, try making a clockwise 90-degree turn
		if turn {
			newDirection := (direction + 1) % 4
			maxStep = max(maxStep, dfs(nx, ny, newDirection, false, 2-target))
		}

		memo[nx][ny][direction][turnIndex] = maxStep + 1
		return maxStep + 1
	}

	res := 0
	// Try starting from every cell that has value 1
	for i := range m {
		for j := 0; j < n; j++ {
			if grid[i][j] == 1 {
				// Try all 4 diagonal directions
				for direction := 0; direction < 4; direction++ {
					res = max(res, dfs(i, j, direction, true, 2)+1)
				}
			}
		}
	}

	return res
}
