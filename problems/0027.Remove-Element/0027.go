package leetcode

// 27. Remove Element
// https://leetcode.com/problems/remove-element/description/
func removeElement(nums []int, val int) int {
	k := 0
	for _, num := range nums {
		if num != val {
			nums[k] = num
			k++
		}
	}
	return k
}
