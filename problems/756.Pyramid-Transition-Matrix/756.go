package leetcode

// 756. Pyramid Transition Matrix
// https://leetcode.com/problems/pyramid-transition-matrix/description/
func pyramidTransition(bottom string, allowed []string) bool {
	nexts := map[string][]byte{}
	for _, v := range allowed {
		nexts[v[:2]] = append(nexts[v[:2]], v[2])
	}

	var dfs func(cur, next []byte) bool
	dfs = func(cur, next []byte) bool {
		if len(cur) == 1 {
			return true
		}
		if len(cur) == len(next)+1 {
			return dfs(next, nil)
		}
		i := len(next)
		s := string(cur[i : i+2])
		for _, c := range nexts[s] {
			if dfs(cur, append(next, c)) {
				return true
			}
		}
		return false
	}

	return dfs([]byte(bottom), nil)
}
