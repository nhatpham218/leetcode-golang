package leetcode

import "sort"

// 757. Set Intersection Size At Least Two
// https://leetcode.com/problems/set-intersection-size-at-least-two/
func intersectionSizeTwo(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][1] != intervals[j][1] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] > intervals[j][0]
	})

	const INF = -1 << 60
	smaller, larger := INF, INF
	count := 0

	for _, interval := range intervals {
		start, end := interval[0], interval[1]
		if start > larger {
			count += 2
			smaller = end - 1
			larger = end
		} else if start > smaller {
			count += 1
			smaller = larger
			larger = end
		}
	}
	return count
}
