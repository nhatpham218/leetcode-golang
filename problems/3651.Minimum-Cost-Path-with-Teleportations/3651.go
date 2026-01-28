package leetcode

import "sort"

// 3651. Minimum Cost Path with Teleportations
// https://leetcode.com/problems/minimum-cost-path-with-teleportations/description/
func minCost(grid [][]int, k int) int {
	m, n := len(grid), len(grid[0])
	const INF = 1 << 30

	// dp[i][j] = min cost to reach cell (i,j)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}

	// Base case: start at (0,0) with cost 0, only normal moves
	dp[0][0] = 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			}
			if i > 0 && dp[i-1][j] < INF {
				dp[i][j] = min(dp[i][j], dp[i-1][j]+grid[i][j])
			}
			if j > 0 && dp[i][j-1] < INF {
				dp[i][j] = min(dp[i][j], dp[i][j-1]+grid[i][j])
			}
		}
	}

	// Process up to k teleportations
	for t := 0; t < k; t++ {
		// Collect (grid value, dp cost) for all reachable cells
		// We can teleport FROM cell with value v TO any cell with value <= v
		type pair struct{ val, cost int }
		pairs := make([]pair, 0, m*n)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if dp[i][j] < INF {
					pairs = append(pairs, pair{grid[i][j], dp[i][j]})
				}
			}
		}

		if len(pairs) == 0 {
			break
		}

		// Sort by value descending (highest values first)
		sort.Slice(pairs, func(a, b int) bool {
			return pairs[a].val > pairs[b].val
		})

		// Compute prefix minimum: minCostPrefix[i] = min cost among pairs[0..i]
		// Since sorted desc, minCostPrefix[i] = min cost among cells with value >= pairs[i].val
		minCostPrefix := make([]int, len(pairs))
		minCostPrefix[0] = pairs[0].cost
		for i := 1; i < len(pairs); i++ {
			minCostPrefix[i] = min(minCostPrefix[i-1], pairs[i].cost)
		}

		// New dp: start with previous values (option to not teleport)
		newDp := make([][]int, m)
		for i := range newDp {
			newDp[i] = make([]int, n)
			for j := range newDp[i] {
				newDp[i][j] = dp[i][j]
			}
		}

		// Apply teleportation to each cell
		// To teleport to cell (i,j), we need a source cell with value >= grid[i][j]
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				targetVal := grid[i][j]
				// Binary search: find first index where pairs[idx].val < targetVal
				lo, hi := 0, len(pairs)
				for lo < hi {
					mid := (lo + hi) / 2
					if pairs[mid].val >= targetVal {
						lo = mid + 1
					} else {
						hi = mid
					}
				}
				// pairs[0..lo-1] all have val >= targetVal
				if lo > 0 {
					newDp[i][j] = min(newDp[i][j], minCostPrefix[lo-1])
				}
			}
		}

		// Propagate normal moves after teleportation
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if i > 0 && newDp[i-1][j] < INF {
					newDp[i][j] = min(newDp[i][j], newDp[i-1][j]+grid[i][j])
				}
				if j > 0 && newDp[i][j-1] < INF {
					newDp[i][j] = min(newDp[i][j], newDp[i][j-1]+grid[i][j])
				}
			}
		}

		dp = newDp
	}

	return dp[m-1][n-1]
}
