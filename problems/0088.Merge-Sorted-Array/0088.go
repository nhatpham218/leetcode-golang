package leetcode

// 88. Merge Sorted Array
// https://leetcode.com/problems/merge-sorted-array/description/
func merge(nums1 []int, m int, nums2 []int, n int) {
	copiedNums1 := make([]int, m)
	for i := range m {
		copiedNums1[i] = nums1[i]
	}

	i, i1, i2 := 0, 0, 0
	for i < m+n {
		if i1 < m && ((i2 < n && copiedNums1[i1] <= nums2[i2]) || i2 == n) {
			nums1[i] = copiedNums1[i1]
			i1++
		} else if i2 < n {
			nums1[i] = nums2[i2]
			i2++
		}
		i++
	}
}
