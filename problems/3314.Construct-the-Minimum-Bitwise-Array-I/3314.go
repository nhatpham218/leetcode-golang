package leetcode

// 3314. Construct the Minimum Bitwise Array I
// https://leetcode.com/problems/construct-the-minimum-bitwise-array-i/description/
func minBitwiseArray(nums []int) []int {
	ans := make([]int, len(nums))

	for i, p := range nums {
		found := -1

		for x := 0; x <= p; x++ {
			if (x | (x + 1)) == p {
				found = x
				break
			}
		}

		ans[i] = found
	}

	return ans
}
