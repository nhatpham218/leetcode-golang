package leetcode

import "strings"

// 2125. Number of Laser Beams in a Bank
// https://leetcode.com/problems/number-of-laser-beams-in-a-bank/
func numberOfBeams(bank []string) int {
	result := 0
	prevCount := 0
	for _, row := range bank {
		count := strings.Count(row, "1")
		if count > 0 {
			result += prevCount * count
			prevCount = count
		}
	}
	return result
}
