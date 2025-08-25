package leetcode

// 3660. Jump Game IX
// You are given an integer array nums.
// From any index i, you can jump to another index j under the following rules:
// Jump to index j where j > i is allowed only if nums[j] < nums[i].
// Jump to index j where j < i is allowed only if nums[j] > nums[i].
// For each index i, find the maximum value in nums that can be reached by following any sequence of valid jumps starting at i.
// Return an array ans where ans[i] is the maximum value reachable starting from index i.
func maxValue(nums []int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	// Find cuts using prefix max and suffix min
	// A cut happens at position i if prefixMax[i] <= suffixMin[i+1]
	// This means all values to the left are <= all values to the right

	prefixMax := make([]int, n)
	suffixMin := make([]int, n)

	// Build prefix maximums
	prefixMax[0] = nums[0]
	for i := 1; i < n; i++ {
		prefixMax[i] = max(prefixMax[i-1], nums[i])
	}

	// Build suffix minimums
	suffixMin[n-1] = nums[n-1]
	for i := n - 2; i >= 0; i-- {
		suffixMin[i] = min(suffixMin[i+1], nums[i])
	}

	// Find connected components
	ans := make([]int, n)
	start := 0

	for i := 0; i < n; i++ {
		// Check if there's a cut after position i
		if i == n-1 || prefixMax[i] <= suffixMin[i+1] {
			// Found end of current component at position i
			// Find maximum value in component [start, i]
			maxVal := nums[start]
			for j := start; j <= i; j++ {
				maxVal = max(maxVal, nums[j])
			}

			// Set answer for all positions in this component
			for j := start; j <= i; j++ {
				ans[j] = maxVal
			}

			start = i + 1
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
