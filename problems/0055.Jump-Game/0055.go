package leetcode

// 55. Jump Game
// https://leetcode.com/problems/jump-game/description/
func canJump(nums []int) bool {
	n := len(nums)
	queue := []int{0}
	visited := make([]bool, n)
	visited[0] = true
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr+nums[curr] >= n-1 {
			return true
		}
		for i := 1; i <= nums[curr]; i++ {
			next := curr + i
			if next < n && !visited[next] {
				queue = append(queue, next)
				visited[next] = true
			}
		}
	}
	return false
}
