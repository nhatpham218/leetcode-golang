package leetcode

import (
	"sort"
	"strings"
)

// 0761. Special Binary String
// https://leetcode.com/problems/special-binary-string/description/
//
// Approach: Treat as balanced parentheses (1 = "(", 0 = ")"). Decompose into
// "mountains" (minimal special substrings) by tracking balance. For each mountain,
// recursively process its interior, then sort all mountains in descending order
// to get the lexicographically largest result.
func makeLargestSpecial(s string) string {
	var mountains []string
	bal, start := 0, 0

	for i := 0; i < len(s); i++ {
		if s[i] == '1' {
			bal++
		} else {
			bal--
		}
		if bal == 0 {
			inner := makeLargestSpecial(s[start+1 : i])
			mountains = append(mountains, "1"+inner+"0")
			start = i + 1
		}
	}

	sort.Slice(mountains, func(i, j int) bool { return mountains[i] > mountains[j] })
	return strings.Join(mountains, "")
}
