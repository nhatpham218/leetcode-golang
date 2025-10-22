package leetcode

import "sort"

// 3347. Maximum Frequency of an Element After Performing Operations II
// https://leetcode.com/problems/maximum-frequency-of-an-element-after-performing-operations-ii/
func maxFrequency(nums []int, k int, numOperations int) int {
	n := len(nums)
	sort.Ints(nums)

	maxFreq := 0
	left := 0
	right := 0

	// Phase 1: Consider existing values as targets
	for i := 0; i < n; i++ {
		currVal := nums[i]
		leftBound := max(1, currVal-k)
		rightBound := min(nums[n-1], currVal+k)

		// Count frequency of current value
		freq := 1
		nextIdx := i + 1
		for nextIdx < n && nums[nextIdx] == currVal {
			freq++
			nextIdx++
		}
		i = nextIdx - 1

		// Move left pointer to first element >= leftBound
		for left < n && nums[left] < leftBound {
			left++
		}

		// Move right pointer to last element <= rightBound
		right = max(right, i)
		for right+1 < n && nums[right+1] <= rightBound {
			right++
		}

		// Calculate max frequency for this target
		// freq: already at target, totalInRange-freq: need conversion
		totalInRange := right - left + 1
		maxFreq = max(maxFreq, freq+min(totalInRange-freq, numOperations))
	}

	// Phase 2: Consider virtual targets (overlapping intervals)
	// Elements within 2k range can converge to a value between them
	for left, right = 0, 0; right < n; right++ {
		currVal := nums[right]
		leftBound := max(1, currVal-2*k)
		for left < right && nums[left] < leftBound {
			left++
		}
		maxFreq = max(maxFreq, min(right-left+1, numOperations))
	}

	return maxFreq
}
