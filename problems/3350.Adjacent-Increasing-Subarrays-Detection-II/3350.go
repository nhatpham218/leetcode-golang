package leetcode

// 3350. Adjacent Increasing Subarrays Detection II
// https://leetcode.com/problems/adjacent-increasing-subarrays-detection-ii/
func maxIncreasingSubarrays(nums []int) int {
	n := len(nums)
	maxK := 1
	prev := 0
	curr := 1

	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			curr++
		} else {
			maxK = max(maxK, min(prev, curr))
			maxK = max(maxK, curr/2)
			prev = curr
			curr = 1
		}
	}

	maxK = max(maxK, min(prev, curr))
	maxK = max(maxK, curr/2)

	return maxK
}
