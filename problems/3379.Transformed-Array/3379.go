package leetcode

// 3379. Transformed Array
// https://leetcode.com/problems/transformed-array/description/
func constructTransformedArray(nums []int) []int {
	result := make([]int, len(nums))
	for i, v := range nums {
		result[i] = nums[getCircularIndex(i+v, len(nums))]
	}
	return result
}

func getCircularIndex(index int, length int) int {
	return ((index % length) + length) % length
}
