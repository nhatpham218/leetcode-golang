package leetcode

// 1437. Check If All 1's Are at Least Length K Places Away
// https://leetcode.com/problems/check-if-all-1s-are-at-least-length-k-places-away/
func kLengthApart(nums []int, k int) bool {
	lastKpos := -1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			if lastKpos != -1 && i-lastKpos <= k {
				return false
			}
			lastKpos = i
		}
	}
	return true
}
