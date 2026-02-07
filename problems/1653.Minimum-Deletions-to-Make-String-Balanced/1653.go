package leetcode

// 1653. Minimum Deletions to Make String Balanced
// https://leetcode.com/problems/minimum-deletions-to-make-string-balanced/description/
func minimumDeletions(s string) int {
	leftB := 0
	rightA := 0
	for _, c := range s {
		if c == 'a' {
			rightA++
		}
	}
	min := rightA
	for _, c := range s {
		if c == 'a' {
			rightA--
		} else {
			leftB++
		}
		if leftB+rightA < min {
			min = leftB + rightA
		}
	}
	return min
}
