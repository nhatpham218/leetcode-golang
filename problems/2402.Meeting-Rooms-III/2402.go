package leetcode

import "sort"

func mostBooked(n int, meetings [][]int) int {
	sort.Slice(meetings, func(i, j int) bool {
		return meetings[i][0] < meetings[j][0]
	})

	count := make([]int, n)
	endAt := make([]int64, n)

	for _, m := range meetings {
		start, end := int64(m[0]), int64(m[1])
		dur := end - start

		freeRoom := -1
		minEnd, minRoom := int64(1<<62), 0

		for i := 0; i < n; i++ {
			if endAt[i] < minEnd {
				minEnd, minRoom = endAt[i], i
			}
			if endAt[i] <= start {
				freeRoom = i
				break
			}
		}

		if freeRoom != -1 {
			endAt[freeRoom] = end
			count[freeRoom]++
		} else {
			endAt[minRoom] += dur
			count[minRoom]++
		}
	}

	maxCnt, res := 0, 0
	for i, c := range count {
		if c > maxCnt {
			maxCnt, res = c, i
		}
	}
	return res
}
