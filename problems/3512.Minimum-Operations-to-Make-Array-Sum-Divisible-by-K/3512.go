package leetcode

// 3512. Minimum Operations to Make Array Sum Divisible by K
// https://leetcode.com/problems/minimum-operations-to-make-array-sum-divisible-by-k/
func minOperations(nums []int, k int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum % k
}
