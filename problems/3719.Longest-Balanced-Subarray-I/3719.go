package leetcode

// 3719. Longest Balanced Subarray I
// https://leetcode.com/problems/longest-balanced-subarray-i/description/
func longestBalanced(nums []int) int {
	n := len(nums)
	best := 0
	for i := 0; i < n; i++ {
		even := make(map[int]struct{})
		odd := make(map[int]struct{})
		for j := i; j < n; j++ {
			v := nums[j]
			if v%2 == 0 {
				even[v] = struct{}{}
			} else {
				odd[v] = struct{}{}
			}
			if len(even) == len(odd) {
				best = max(best, j-i+1)
			}
		}
	}
	return best
}
