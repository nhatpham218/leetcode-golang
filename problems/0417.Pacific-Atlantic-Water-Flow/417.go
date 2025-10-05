package leetcode

// 417. Pacific Atlantic Water Flow
// https://leetcode.com/problems/pacific-atlantic-water-flow/

// DFS approach: Use depth-first search to find cells reachable from each ocean
func pacificAtlantic(heights [][]int) [][]int {
	if len(heights) == 0 || len(heights[0]) == 0 {
		return [][]int{}
	}

	m, n := len(heights), len(heights[0])

	// Arrays to track which cells can reach Pacific and Atlantic
	pacific := make([][]bool, m)
	atlantic := make([][]bool, m)
	for i := range pacific {
		pacific[i] = make([]bool, n)
		atlantic[i] = make([]bool, n)
	}

	// Directions: up, down, left, right
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// DFS function to mark reachable cells
	var dfs func([][]bool, int, int)
	dfs = func(visited [][]bool, r, c int) {
		if visited[r][c] {
			return
		}
		visited[r][c] = true

		for _, dir := range directions {
			nr, nc := r+dir[0], c+dir[1]
			if nr >= 0 && nr < m && nc >= 0 && nc < n &&
				heights[nr][nc] >= heights[r][c] {
				dfs(visited, nr, nc)
			}
		}
	}

	// Start DFS from Pacific edges (top and left)
	for i := 0; i < m; i++ {
		dfs(pacific, i, 0)
	}
	for j := 0; j < n; j++ {
		dfs(pacific, 0, j)
	}

	// Start DFS from Atlantic edges (bottom and right)
	for i := 0; i < m; i++ {
		dfs(atlantic, i, n-1)
	}
	for j := 0; j < n; j++ {
		dfs(atlantic, m-1, j)
	}

	// Find cells that can reach both oceans
	var result [][]int
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if pacific[i][j] && atlantic[i][j] {
				result = append(result, []int{i, j})
			}
		}
	}

	return result
}
