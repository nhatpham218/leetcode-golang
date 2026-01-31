package leetcode

// 960. Delete Columns to Make Sorted III
// https://leetcode.com/problems/delete-columns-to-make-sorted-iii/
func minDeletionSize(strs []string) int {
	cols := len(strs[0])
	maxLen := make([]int, cols)
	for i := range maxLen {
		maxLen[i] = 1
	}
	for i := cols - 2; i >= 0; i-- {
	search:
		for j := i + 1; j < cols; j++ {
			for _, row := range strs {
				if row[i] > row[j] {
					continue search
				}
			}
			maxLen[i] = max(maxLen[i], 1+maxLen[j])
		}
	}
	kept := 0
	for _, l := range maxLen {
		kept = max(kept, l)
	}
	return cols - kept
}
