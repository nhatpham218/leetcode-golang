package leetcode

// 80. Remove Duplicates from Sorted Array II
// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii/description/
func removeDuplicates(nums []int) int {
	copiedNums := make([]int, len(nums))
	copy(copiedNums, nums)
	k := 1
	if len(nums) < 2 {
		return len(nums)
	}
	for i := 2; i < len(nums); i++ {
		if copiedNums[i] != copiedNums[i-2] {
			nums[k] = copiedNums[i]
			k++
		}
	}
	return k
}
