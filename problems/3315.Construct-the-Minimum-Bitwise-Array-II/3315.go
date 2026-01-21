package leetcode

// 3315. Construct the Minimum Bitwise Array II
// https://leetcode.com/problems/construct-the-minimum-bitwise-array-ii/description/
func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))
	for i, v := range nums {
		if v == 2 {
			ans[i] = -1
			continue
		}
		t := v
		c := 0 // length of right part
		rightPart := 0
		for t&1 == 1 {
			c += 1
			t >>= 1
			rightPart = rightPart<<1 | 1
		}
		leftPart := v & (^0 << (c + 1))
		ans[i] = leftPart | rightPart>>1
	}
	return ans
}
