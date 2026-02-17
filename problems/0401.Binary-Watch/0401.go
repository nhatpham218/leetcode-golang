package leetcode

import (
	"fmt"
	"math/bits"
)

// 0401. Binary Watch
// https://leetcode.com/problems/binary-watch/description/
func readBinaryWatch(turnedOn int) []string {
	if turnedOn > 10 || turnedOn < 0 {
		return []string{}
	}
	var ans []string
	for h := 0; h < 12; h++ {
		for m := 0; m < 60; m++ {
			if bits.OnesCount(uint(h))+bits.OnesCount(uint(m)) == turnedOn {
				ans = append(ans, fmt.Sprintf("%d:%02d", h, m))
			}
		}
	}
	return ans
}
