package leetcode

import "sort"

// 2975. Maximum Square Area by Removing Fences From a Field
// https://leetcode.com/problems/maximum-square-area-by-removing-fences-from-a-field/description/
func maximizeSquareArea(totalHeight int, totalWidth int, horizontalCuts []int, verticalCuts []int) int {
	const MOD = 1000000007

	prep := func(cuts []int, limit int) []int {
		sort.Ints(cuts)
		out := []int{1}
		out = append(out, cuts...)
		out = append(out, limit)
		return out
	}

	h := prep(horizontalCuts, totalHeight)
	v := prep(verticalCuts, totalWidth)

	gapSet := make(map[int]bool)

	for i := 0; i < len(h); i++ {
		for j := i + 1; j < len(h); j++ {
			gapSet[h[j]-h[i]] = true
		}
	}

	best := 0
	for i := 0; i < len(v); i++ {
		for j := i + 1; j < len(v); j++ {
			d := v[j] - v[i]
			if d > best && gapSet[d] {
				best = d
			}
		}
	}

	if best == 0 {
		return -1
	}
	return int((int64(best) * int64(best)) % MOD)
}
