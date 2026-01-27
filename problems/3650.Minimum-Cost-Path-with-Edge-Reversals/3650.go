package leetcode

// 3650. Minimum Cost Path with Edge Reversals
// https://leetcode.com/problems/minimum-cost-path-with-edge-reversals/description/
func minCost(n int, edges [][]int) int {
	graph := make(map[int][][]int)
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], []int{edge[1], edge[2]})
		graph[edge[1]] = append(graph[edge[1]], []int{edge[0], 2 * edge[2]})
	}

	dist := make([]int, n)
	for i := range dist {
		dist[i] = -1
	}
	dist[0] = 0

	queue := []int{0}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		for _, nb := range graph[curr] {
			newDist := dist[curr] + nb[1]
			if dist[nb[0]] == -1 || dist[nb[0]] > newDist {
				dist[nb[0]] = newDist
				queue = append(queue, nb[0])
			}
		}
	}

	return dist[n-1]
}
