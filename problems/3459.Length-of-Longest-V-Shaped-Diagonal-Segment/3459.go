package leetcode

// MemoState represents the state for memoization
type MemoState struct {
	direction   int
	hasMadeTurn bool
}

// 3459. Length of Longest V-Shaped Diagonal Segment
func lenOfVDiagonal(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	m, n := len(grid), len(grid[0])
	// Direction vectors: {1,1}, {1,-1}, {-1,-1}, {-1,1}
	// Representing: down-right, down-left, up-left, up-right
	dirs := [][]int{{1, 1}, {1, -1}, {-1, -1}, {-1, 1}}

	// Memoization table: [row][col] -> map[MemoState]int
	memo := make([][]map[MemoState]int, m)
	for i := range memo {
		memo[i] = make([]map[MemoState]int, n)
		for j := range memo[i] {
			memo[i][j] = make(map[MemoState]int)
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

		state := MemoState{direction: direction, hasMadeTurn: !turn}

		// Check memoization
		if val, exists := memo[nx][ny][state]; exists {
			return val
		}

		// Continue walking in the original direction
		maxStep := dfs(nx, ny, direction, turn, 2-target)

		// If we haven't made a turn yet, try making a clockwise 90-degree turn
		if turn {
			newDirection := (direction + 1) % 4
			maxStep = max(maxStep, dfs(nx, ny, newDirection, false, 2-target))
		}

		memo[nx][ny][state] = maxStep + 1
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
