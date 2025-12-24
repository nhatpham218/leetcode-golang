package leetcode

import "sort"

// 3074. Apple Redistribution into Boxes
// https://leetcode.com/problems/apple-redistribution-into-boxes/description/
func minimumBoxes(apple []int, capacity []int) int {
	sort.Ints(capacity)
	totalApple := 0
	for i := 0; i < len(apple); i++ {
		totalApple += apple[i]
	}

	for i := 1; i < len(capacity); i++ {
		totalApple -= capacity[len(capacity)-i]
		if totalApple <= 0 {
			return i
		}
	}
	return len(capacity)
}
