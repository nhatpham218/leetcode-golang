package leetcode

// 11. Container With Most Water
// https://leetcode.com/problems/container-with-most-water/
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	maxWater := 0

	for left < right {
		width := right - left
		currHeight := min(height[left], height[right])
		area := width * currHeight

		if area > maxWater {
			maxWater = area
		}

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}
