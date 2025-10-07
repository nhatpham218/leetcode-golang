package leetcode

// 778. Swim in Rising Water
// https://leetcode.com/problems/swim-in-rising-water/
func swimInWater(grid [][]int) int {
	n := len(grid)
	if n == 1 {
		return grid[0][0]
	}

	left, right := 0, n*n-1

	for left < right {
		mid := left + (right-left)/2
		if canReach(grid, mid) {
			right = mid
		} else {
			left = mid + 1
		}
	}

	return left
}

func canReach(grid [][]int, t int) bool {
	n := len(grid)
	visited := make([][]bool, n)
	for i := range visited {
		visited[i] = make([]bool, n)
	}

	return dfs(grid, visited, 0, 0, t)
}

func dfs(grid [][]int, visited [][]bool, r, c, t int) bool {
	n := len(grid)
	if r < 0 || r >= n || c < 0 || c >= n || visited[r][c] || grid[r][c] > t {
		return false
	}

	if r == n-1 && c == n-1 {
		return true
	}

	visited[r][c] = true

	dirs := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for _, dir := range dirs {
		if dfs(grid, visited, r+dir[0], c+dir[1], t) {
			return true
		}
	}

	return false
}
