package leetcode

// 3349. Adjacent Increasing Subarrays Detection I
// https://leetcode.com/problems/adjacent-increasing-subarrays-detection-i/
func hasIncreasingSubarrays(nums []int, k int) bool {
	n := len(nums)

	for i := 0; i <= n-2*k; i++ {
		if isStrictlyIncreasing(nums, i, i+k-1) {
			if isStrictlyIncreasing(nums, i+k, i+2*k-1) {
				return true
			}
		}
	}

	return false
}

func isStrictlyIncreasing(nums []int, start, end int) bool {
	for i := start; i < end; i++ {
		if nums[i] >= nums[i+1] {
			return false
		}
	}
	return true
}
