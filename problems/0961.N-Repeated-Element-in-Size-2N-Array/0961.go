package leetcode

// 961. N-Repeated Element in Size 2N Array
// https://leetcode.com/problems/n-repeated-element-in-size-2n-array/description/
func repeatedNTimes(nums []int) int {
	vis := make(map[int]bool)
	for _, num := range nums {
		if vis[num] {
			return num
		}
		vis[num] = true
	}
	return 0
}
