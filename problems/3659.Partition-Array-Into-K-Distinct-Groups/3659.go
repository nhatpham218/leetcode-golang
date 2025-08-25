package leetcode

// 3659. Partition Array Into K-Distinct Groups
// You are given an integer array nums and an integer k.
// Your task is to determine whether it is possible to partition all elements of nums into one or more groups such that:
// Each group contains exactly k distinct elements.
// Each element in nums must be assigned to exactly one group.
// Return true if such a partition is possible, otherwise return false.
func partitionArray(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}

	max := len(nums) / k
	memo := make(map[int]int)
	for _, num := range nums {
		memo[num]++
		if memo[num] > max {
			return false
		}
	}

	return true
}
