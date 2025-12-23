package leetcode

import "sort"

// 2054. Two Best Non-Overlapping Events
// https://leetcode.com/problems/two-best-non-overlapping-events/description/
func maxTwoEvents(events [][]int) int {
	sort.Slice(events, func(i, j int) bool {
		return events[i][0] < events[j][0]
	})
	dp := make([][]int, len(events))
	for i := range dp {
		dp[i] = make([]int, 3)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}
	return findEvents(events, 0, 0, dp)
}

func findEvents(events [][]int, idx, cnt int, dp [][]int) int {
	if cnt == 2 || idx >= len(events) {
		return 0
	}
	if dp[idx][cnt] == -1 {
		end := events[idx][1]
		lo, hi := idx+1, len(events)-1
		for lo < hi {
			mid := lo + ((hi - lo) >> 1)
			if events[mid][0] > end {
				hi = mid
			} else {
				lo = mid + 1
			}
		}
		include := events[idx][2]
		if lo < len(events) && events[lo][0] > end {
			include += findEvents(events, lo, cnt+1, dp)
		}
		exclude := findEvents(events, idx+1, cnt, dp)
		dp[idx][cnt] = max(include, exclude)
	}
	return dp[idx][cnt]
}
