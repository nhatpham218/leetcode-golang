package leetcode

// 2976. Minimum Cost to Convert String I
// https://leetcode.com/problems/minimum-cost-to-convert-string-i/description/
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	const inf = int64(1e18)

	// Initialize distance matrix for 26 letters
	dist := [26][26]int64{}
	for i := 0; i < 26; i++ {
		for j := 0; j < 26; j++ {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = inf
			}
		}
	}

	// Build edges from the transformation rules (keep minimum cost)
	for i := 0; i < len(original); i++ {
		from := int(original[i] - 'a')
		to := int(changed[i] - 'a')
		if int64(cost[i]) < dist[from][to] {
			dist[from][to] = int64(cost[i])
		}
	}

	// Floyd-Warshall: compute all-pairs shortest paths
	for k := 0; k < 26; k++ {
		for i := 0; i < 26; i++ {
			for j := 0; j < 26; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	// Calculate total cost
	var totalCost int64
	for i := 0; i < len(source); i++ {
		from := int(source[i] - 'a')
		to := int(target[i] - 'a')
		if dist[from][to] == inf {
			return -1
		}
		totalCost += dist[from][to]
	}

	return totalCost
}
