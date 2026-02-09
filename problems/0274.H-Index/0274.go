package leetcode

// 274. H-Index
// https://leetcode.com/problems/h-index/description/
func hIndex(citations []int) int {
	n := len(citations)
	// count[i] = papers with i citations (i < n); count[n] = papers with >= n citations
	count := make([]int, n+1)
	for _, c := range citations {
		if c >= n {
			count[n]++
		} else {
			count[c]++
		}
	}
	// papers = count of papers with >= h citations
	var papers int
	for h := n; h >= 0; h-- {
		papers += count[h]
		if papers >= h {
			return h
		}
	}
	return 0
}
