package leetcode

const inf int64 = 1e18

// 2977. Minimum Cost to Convert String II
// https://leetcode.com/problems/minimum-cost-to-convert-string-ii/description/
func minimumCost(source string, target string, original []string, changed []string, cost []int) int64 {
	n := len(source)
	m := len(original)

	// Step 1: Assign unique IDs to all strings in original and changed
	stringToID := make(map[string]int)
	idCount := 0

	getID := func(s string) int {
		if id, exists := stringToID[s]; exists {
			return id
		}
		stringToID[s] = idCount
		idCount++
		return idCount - 1
	}

	// Collect all unique strings
	for i := 0; i < m; i++ {
		getID(original[i])
		getID(changed[i])
	}

	// Step 2: Build cost matrix and apply Floyd-Warshall
	// dist[i][j] = minimum cost to convert string with id i to string with id j
	dist := make([][]int64, idCount)
	for i := range dist {
		dist[i] = make([]int64, idCount)
		for j := range dist[i] {
			if i == j {
				dist[i][j] = 0
			} else {
				dist[i][j] = inf
			}
		}
	}

	// Initialize direct edges from the cost array
	for i := 0; i < m; i++ {
		u := stringToID[original[i]]
		v := stringToID[changed[i]]
		if int64(cost[i]) < dist[u][v] {
			dist[u][v] = int64(cost[i])
		}
	}

	// Floyd-Warshall algorithm
	for k := 0; k < idCount; k++ {
		for i := 0; i < idCount; i++ {
			for j := 0; j < idCount; j++ {
				if dist[i][k]+dist[k][j] < dist[i][j] {
					dist[i][j] = dist[i][k] + dist[k][j]
				}
			}
		}
	}

	// Step 3: Build Trie for efficient substring matching
	type TrieNode struct {
		children [26]*TrieNode
		strID    int // -1 if not end of a string
	}

	newTrieNode := func() *TrieNode {
		return &TrieNode{strID: -1}
	}

	sourceRoot := newTrieNode()
	targetRoot := newTrieNode()

	// Insert string into trie and return its ID
	insertTrie := func(root *TrieNode, s string, id int) {
		node := root
		for _, c := range s {
			idx := c - 'a'
			if node.children[idx] == nil {
				node.children[idx] = newTrieNode()
			}
			node = node.children[idx]
		}
		node.strID = id
	}

	// Insert all original strings into sourceTrie
	// Insert all changed strings into targetTrie
	for i := 0; i < m; i++ {
		insertTrie(sourceRoot, original[i], stringToID[original[i]])
		insertTrie(targetRoot, changed[i], stringToID[changed[i]])
	}

	// Step 4: DP to find minimum cost
	// dp[i] = minimum cost to convert source[0..i-1] to target[0..i-1]
	dp := make([]int64, n+1)
	for i := range dp {
		dp[i] = inf
	}
	dp[0] = 0

	for i := 0; i < n; i++ {
		if dp[i] == inf {
			continue
		}

		// Case 1: source[i] == target[i], no cost needed
		if source[i] == target[i] {
			if dp[i] < dp[i+1] {
				dp[i+1] = dp[i]
			}
		}

		// Case 2: Try all substrings starting at position i
		// Use Trie to find matching substrings
		srcNode := sourceRoot
		tgtNode := targetRoot

		for j := i; j < n; j++ {
			srcIdx := source[j] - 'a'
			tgtIdx := target[j] - 'a'

			if srcNode != nil {
				srcNode = srcNode.children[srcIdx]
			}
			if tgtNode != nil {
				tgtNode = tgtNode.children[tgtIdx]
			}

			// If both paths exist and end at valid strings
			if srcNode != nil && tgtNode != nil &&
				srcNode.strID != -1 && tgtNode.strID != -1 {
				srcID := srcNode.strID
				tgtID := tgtNode.strID
				if dist[srcID][tgtID] < inf {
					newCost := dp[i] + dist[srcID][tgtID]
					if newCost < dp[j+1] {
						dp[j+1] = newCost
					}
				}
			}

			// Early termination if both tries have no path
			if srcNode == nil && tgtNode == nil {
				break
			}
		}
	}

	if dp[n] == inf {
		return -1
	}
	return dp[n]
}
