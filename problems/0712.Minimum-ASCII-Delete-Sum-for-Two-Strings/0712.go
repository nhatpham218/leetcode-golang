package leetcode

// 712. Minimum ASCII Delete Sum for Two Strings
// https://leetcode.com/problems/minimum-ascii-delete-sum-for-two-strings/description/

// Space-optimized DP solution
// dp[i][j] = minimum ASCII delete sum to make s1[0:i] and s2[0:j] equal
// Time: O(m*n), Space: O(n)
func minimumDeleteSum(s1 string, s2 string) int {
	// Base case: dp[0][j] = sum of ASCII values of s2[0:j]
	// (delete all characters from s2 when s1 is empty)
	prev_row := make([]int, len(s2)+1)
	for j := 1; j <= len(s2); j++ {
		prev_row[j] = prev_row[j-1] + int(s2[j-1])
	}

	for i := 1; i <= len(s1); i++ {
		// Base case: dp[i][0] = sum of ASCII values of s1[0:i]
		// (delete all characters from s1 when s2 is empty)
		curr_row := []int{prev_row[0] + int(s1[i-1])}
		for j := 1; j <= len(s2); j++ {
			if s1[i-1] == s2[j-1] {
				// Characters match, no deletion needed
				curr_row = append(curr_row, prev_row[j-1])
			} else {
				// Take minimum of: delete s1[i-1] or delete s2[j-1]
				curr_row = append(curr_row, min(prev_row[j]+int(s1[i-1]), curr_row[j-1]+int(s2[j-1])))
			}
		}
		prev_row = curr_row
	}

	return prev_row[len(prev_row)-1]
}
