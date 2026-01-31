package leetcode

// 944. Delete Columns to Make Sorted
// https://leetcode.com/problems/delete-columns-to-make-sorted/
func minDeletionSize(strs []string) int {
	length := len(strs[0])
	ans := 0
	for i := 0; i < length; i++ {
		for j := 1; j < len(strs); j++ {
			if strs[j][i] < strs[j-1][i] {
				ans++
				break
			}
		}
	}

	return ans
}
