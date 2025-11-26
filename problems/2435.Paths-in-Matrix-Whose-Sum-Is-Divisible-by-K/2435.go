package leetcode

// 2435. Paths in Matrix Whose Sum Is Divisible by K
// https://leetcode.com/problems/paths-in-matrix-whose-sum-is-divisible-by-k/
func numberOfPaths(grid [][]int, k int) int {
	MOD := 1000000007
	m, n := len(grid), len(grid[0])

	prev := make([][]int, n)
	curr := make([][]int, n)
	for i := 0; i < n; i++ {
		prev[i] = make([]int, k)
		curr[i] = make([]int, k)
	}

	sum := 0
	for j := 0; j < n; j++ {
		sum = (sum + grid[0][j]) % k
		prev[j][sum] = 1
	}

	sum = grid[0][0] % k
	for i := 1; i < m; i++ {
		sum = (sum + grid[i][0]) % k
		for idx := 0; idx < k; idx++ {
			curr[0][idx] = 0
		}
		curr[0][sum] = 1

		for j := 1; j < n; j++ {
			for idx := 0; idx < k; idx++ {
				curr[j][idx] = 0
			}
			val := grid[i][j]
			for rem := 0; rem < k; rem++ {
				newRem := (rem + val) % k
				curr[j][newRem] = (prev[j][rem] + curr[j-1][rem]) % MOD
			}
		}
		prev, curr = curr, prev
	}

	return prev[n-1][0]
}
