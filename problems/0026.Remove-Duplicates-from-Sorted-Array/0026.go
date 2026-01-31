package leetcode

// 26. Remove Duplicates from Sorted Array
// https://leetcode.com/problems/remove-duplicates-from-sorted-array/description/
func removeDuplicates(nums []int) int {
	k := 0
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[i-2] {
			k++
			nums[k] = nums[i]
		}
	}
	return k + 1
}
