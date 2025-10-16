package leetcode

// 2598. Smallest Missing Non-negative Integer After Operations
// https://leetcode.com/problems/smallest-missing-non-negative-integer-after-operations/
func findSmallestInteger(nums []int, value int) int {
	freq := make(map[int]int)

	for _, num := range nums {
		rem := ((num % value) + value) % value
		freq[rem]++
	}

	mex := 0
	for {
		rem := mex % value
		if freq[rem] == 0 {
			return mex
		}
		freq[rem]--
		mex++
	}
}
