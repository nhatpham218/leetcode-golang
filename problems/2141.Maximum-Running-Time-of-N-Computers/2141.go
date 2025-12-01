package leetcode

import "sort"

// 2141. Maximum Running Time of N Computers
// https://leetcode.com/problems/maximum-running-time-of-n-computers/
func maxRunTime(n int, batteries []int) int64 {
	sort.Ints(batteries)

	extra := int64(0)
	for i := 0; i < len(batteries)-n; i++ {
		extra += int64(batteries[i])
	}

	live := batteries[len(batteries)-n:]

	for i := 0; i < n-1; i++ {
		if extra/int64(i+1) < int64(live[i+1]-live[i]) {
			return int64(live[i]) + extra/int64(i+1)
		}

		extra -= int64(i+1) * int64(live[i+1]-live[i])
	}

	return int64(live[n-1]) + extra/int64(n)
}
