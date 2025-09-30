package leetcode

// 2221. Find Triangular Sum of an Array
// https://leetcode.com/problems/find-triangular-sum-of-an-array/

func triangularSum(nums []int) int {
	n := len(nums)

	// Use in-place computation to save space
	for i := n - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			nums[j] = (nums[j] + nums[j+1]) % 10
		}
	}

	return nums[0]
}
