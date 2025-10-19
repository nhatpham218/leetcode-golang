package leetcode

// 1625. Lexicographically Smallest String After Applying Operations
// https://leetcode.com/problems/lexicographically-smallest-string-after-applying-operations/
func findLexSmallestString(s string, a int, b int) string {
	vis := make(map[string]bool)
	smallest := s
	q := []string{s}
	vis[s] = true

	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		if cur < smallest {
			smallest = cur
		}

		added := []byte(cur)
		for i := 1; i < len(added); i += 2 {
			added[i] = byte((int(added[i]-'0')+a)%10) + '0'
		}
		addedStr := string(added)
		if !vis[addedStr] {
			vis[addedStr] = true
			q = append(q, addedStr)
		}

		n := len(cur)
		rotated := cur[n-b:] + cur[:n-b]
		if !vis[rotated] {
			vis[rotated] = true
			q = append(q, rotated)
		}
	}
	return smallest
}
