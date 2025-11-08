package leetcode

// 1611. Minimum One Bit Operations to Make Integers Zero
// https://leetcode.com/problems/minimum-one-bit-operations-to-make-integers-zero/
func minimumOneBitOperations(n int) int {
	if n == 0 {
		return 0
	}

	k := 0
	curr := 1
	for curr*2 <= n {
		curr *= 2
		k++
	}

	return (1 << (k + 1)) - 1 - minimumOneBitOperations(n^curr)
}
