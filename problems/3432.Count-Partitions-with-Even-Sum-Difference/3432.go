package leetcode

// 3432. Count Partitions with Even Sum Difference
// https://leetcode.com/problems/count-partitions-with-even-sum-difference/description/
func countPartitions(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return 0
	}

	return len(nums) - 1
}
