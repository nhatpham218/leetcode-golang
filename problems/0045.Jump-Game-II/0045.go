package leetcode

// 45. Jump Game II
// https://leetcode.com/problems/jump-game-ii/description/
func jump(nums []int) int {
	near, far, jumps := 0, 0, 0

	for far < len(nums)-1 {
		farthest := 0
		for i := near; i <= far; i++ {
			farthest = max(farthest, i+nums[i])
		}
		near = far + 1
		far = farthest
		jumps++
	}

	return jumps
}
