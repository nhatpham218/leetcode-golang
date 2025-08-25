package leetcode

func productQueries(n int, queries [][]int) []int {
	mod := 1000000007

	// Generate powers of 2 based on binary representation of n
	var powers []int
	for i := 0; i < 32; i++ {
		if (n>>i)&1 == 1 {
			powers = append(powers, 1<<i)
		}
	}

	ans := make([]int, len(queries))
	for i, query := range queries {
		left, right := query[0], query[1]
		ans[i] = 1
		for j := left; j <= right; j++ {
			ans[i] = (ans[i] * powers[j]) % mod
		}
	}
	return ans
}
