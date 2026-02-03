package leetcode

// 3637. Trionic Array I
// https://leetcode.com/problems/trionic-array-i/description/
func isTrionic(nums []int) bool {
	p, q, i := 0, 0, 1
	for i < len(nums) && nums[i] > nums[i-1] {
		i++
	}
	p = i - 1
	for i < len(nums) && nums[i] < nums[i-1] {
		i++
	}
	q = i - 1
	for i < len(nums) && nums[i] > nums[i-1] {
		i++
	}

	return p > 0 && q > p && q < len(nums) && i == len(nums)
}
