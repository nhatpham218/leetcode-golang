package leetcode

// 3562. Maximum Profit from Trading Stocks with Discounts
// https://leetcode.com/problems/maximum-profit-from-trading-stocks-with-discounts/description/
func maxProfit(n int, present []int, future []int, hierarchy [][]int, budget int) int {
	// Build tree structure (0-based indexing)
	tree := make([][]int, n)
	for _, e := range hierarchy {
		parent, child := e[0]-1, e[1]-1 // Convert to 0-based
		tree[parent] = append(tree[parent], child)
	}

	// dp[node][parentBought][budget] = max profit
	// parentBought: 0 = false, 1 = true
	dp := make([][][]int, n)
	for i := range dp {
		dp[i] = make([][]int, 2)
		for j := 0; j < 2; j++ {
			dp[i][j] = make([]int, budget+1)
			for k := range dp[i][j] {
				dp[i][j][k] = -1e9 // Initialize as invalid
			}
		}
	}

	// Merge function: combines DP arrays from children using knapsack merging
	merge := func(A, B []int) []int {
		C := make([]int, budget+1)
		for i := range C {
			C[i] = -1e9
		}
		for i := 0; i <= budget; i++ {
			if A[i] < 0 {
				continue
			}
			for j := 0; i+j <= budget; j++ {
				if B[j] >= 0 && C[i+j] < A[i]+B[j] {
					C[i+j] = A[i] + B[j]
				}
			}
		}
		return C
	}

	var dfs func(int)
	dfs = func(u int) {
		// Process children first (post-order DFS)
		for _, v := range tree[u] {
			dfs(v)
		}

		// For each parentBought state (0 = false, 1 = true)
		for pb := 0; pb <= 1; pb++ {
			// Calculate price and profit for current node
			price := present[u]
			if pb == 1 {
				price /= 2 // Discount if parent bought
			}
			profit := future[u] - price

			// Option 1: Skip buying current node
			// Merge children's DP arrays where they don't get discount
			skip := make([]int, budget+1)
			for i := range skip {
				skip[i] = -1e9
			}
			skip[0] = 0 // Base case: 0 budget gives 0 profit
			for _, v := range tree[u] {
				skip = merge(skip, dp[v][0])
			}

			// Option 2: Buy current node
			take := make([]int, budget+1)
			for i := range take {
				take[i] = -1e9
			}

			if price <= budget {
				// Merge children's DP arrays where they get discount
				base := make([]int, budget+1)
				for i := range base {
					base[i] = -1e9
				}
				base[0] = 0 // Base case
				for _, v := range tree[u] {
					base = merge(base, dp[v][1])
				}
				// If we buy current node, we spend 'price' and get 'profit'
				for b := price; b <= budget; b++ {
					if base[b-price] >= 0 {
						take[b] = base[b-price] + profit
					}
				}
			}

			// Take the maximum of skip and take for each budget
			for b := 0; b <= budget; b++ {
				dp[u][pb][b] = max(skip[b], take[b])
			}
		}
	}

	dfs(0) // Start from root (employee 1, which is index 0)

	// Find maximum profit from root with no parent (pb=0)
	ans := 0
	for _, v := range dp[0][0] {
		if v > ans {
			ans = v
		}
	}
	return ans
}
