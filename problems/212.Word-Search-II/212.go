package leetcode

// 212. Word Search II
// https://leetcode.com/problems/word-search-ii/description/
func findWords(board [][]byte, words []string) []string {
	if len(board) == 0 || len(board[0]) == 0 {
		return nil
	}

	// Build Trie from words
	trie := &trieNode{children: [26]*trieNode{}}
	for _, w := range words {
		trie.insert(w)
	}

	m, n := len(board), len(board[0])
	seen := make(map[string]bool)
	var result []string

	var dfs func(node *trieNode, r, c int, path []byte)
	dfs = func(node *trieNode, r, c int, path []byte) {
		if r < 0 || r >= m || c < 0 || c >= n || board[r][c] == 0 {
			return
		}
		ch := board[r][c]
		idx := ch - 'a'
		if node.children[idx] == nil {
			return
		}
		next := node.children[idx]
		path = append(path, ch)

		if next.isEnd {
			w := string(path)
			if !seen[w] {
				seen[w] = true
				result = append(result, w)
			}
		}

		// Mark as visited
		board[r][c] = 0
		dfs(next, r-1, c, path)
		dfs(next, r+1, c, path)
		dfs(next, r, c-1, path)
		dfs(next, r, c+1, path)
		board[r][c] = ch // backtrack
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			dfs(trie, i, j, nil)
		}
	}
	return result
}

type trieNode struct {
	children [26]*trieNode
	isEnd    bool
}

func (t *trieNode) insert(word string) {
	node := t
	for _, c := range word {
		idx := c - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &trieNode{children: [26]*trieNode{}}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}
