package leetcode

// 2154. Keep Multiplying Found Values by Two
// https://leetcode.com/problems/keep-multiplying-found-values-by-two/
func findFinalValue(nums []int, original int) int {
	mapNums := make(map[int]bool)
	for _, num := range nums {
		mapNums[num] = true
	}
	for mapNums[original] {
		original *= 2
	}
	return original
}
