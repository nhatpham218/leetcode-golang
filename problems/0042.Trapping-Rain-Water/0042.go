package leetcode

// 0042. Trapping Rain Water
// https://leetcode.com/problems/trapping-rain-water/description/
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	left, right := 0, n-1
	leftMax, rightMax := 0, 0
	water := 0
	for left < right {
		if height[left] < height[right] {
			if height[left] >= leftMax {
				leftMax = height[left]
			} else {
				water += leftMax - height[left]
			}
			left++
		} else {
			if height[right] >= rightMax {
				rightMax = height[right]
			} else {
				water += rightMax - height[right]
			}
			right--
		}
	}
	return water
}
