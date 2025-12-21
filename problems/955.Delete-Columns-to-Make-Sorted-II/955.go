package leetcode

// 955. Delete Columns to Make Sorted II
// https://leetcode.com/problems/delete-columns-to-make-sorted-ii/description/
func minDeletionSize(strs []string) int {
	n := len(strs)
	if n <= 1 {
		return 0
	}

	m := len(strs[0])
	resolved := make([]bool, n-1)
	unresolved := n - 1
	deletions := 0

	for col := 0; col < m && unresolved > 0; col++ {
		needDelete := false
		for row := 0; row < n-1; row++ {
			if !resolved[row] && strs[row][col] > strs[row+1][col] {
				needDelete = true
				break
			}
		}

		if needDelete {
			deletions++
			continue
		}

		for row := 0; row < n-1; row++ {
			if !resolved[row] && strs[row][col] < strs[row+1][col] {
				resolved[row] = true
				unresolved--
			}
		}
	}

	return deletions
}
