package leetcode

// 3653. XOR After Range Multiplication Queries I
func xorAfterQueries(nums []int, queries [][]int) int {
	const MOD = 1000000007

	// Process each query
	for _, query := range queries {
		l, r, k, v := query[0], query[1], query[2], query[3]

		// Apply the multiplication operation
		idx := l
		for idx <= r {
			nums[idx] = (nums[idx] * v) % MOD
			idx += k
		}
	}

	// Calculate XOR of all elements
	result := 0
	for _, num := range nums {
		result ^= num
	}

	return result
}
