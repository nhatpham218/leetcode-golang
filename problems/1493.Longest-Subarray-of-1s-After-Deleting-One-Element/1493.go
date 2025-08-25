package leetcode

func longestSubarray(nums []int) int {
	maxLen := 0
	left, zeroCount := 0, 0

	for right := 0; right < len(nums); right++ {
		if nums[right] == 0 {
			zeroCount++
		}
		for zeroCount > 1 {
			if nums[left] == 0 {
				zeroCount--
			}
			left++
		}
		if right-left > maxLen {
			maxLen = right - left
		}
	}
	return maxLen
}
