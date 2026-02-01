package leetcode

// 3010. Divide an Array Into Subarrays With Minimum Cost I
// https://leetcode.com/problems/divide-an-array-into-subarrays-with-minimum-cost-i/description/
func minimumCost(nums []int) int {
	ans := nums[0]

	min1, min2 := 100, 100
	for _, num := range nums[1:] {
		if num < min1 {
			if num < min2 {
				min1 = min2
				min2 = num
			} else {
				min1 = num
			}
		}
	}
	return ans + min1 + min2
}
