package leetcode

// 3507. Minimum Pair Removal to Sort Array I
// https://leetcode.com/problems/minimum-pair-removal-to-sort-array-i/description/
func minimumPairRemoval(nums []int) int {
	arr := append([]int{}, nums...)
	ops := 0

	for !isSorted(arr) {
		minSum := int(1e9)
		idx := 0
		for i := 0; i+1 < len(arr); i++ {
			if arr[i]+arr[i+1] < minSum {
				minSum = arr[i] + arr[i+1]
				idx = i
			}
		}
		arr[idx] = minSum
		arr = append(arr[:idx+1], arr[idx+2:]...)
		ops++
	}
	return ops
}

func isSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}
