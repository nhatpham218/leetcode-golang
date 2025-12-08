package leetcode

import "math"

// 1925. Count Square Sum Triples
// https://leetcode.com/problems/count-square-sum-triples/description/
func countTriples(n int) int {
	ans := 0
	for a := 1; a <= n; a++ {
		for b := 1; b <= n; b++ {
			sum := a*a + b*b
			if sum > n*n {
				break
			}
			sqrt := int(math.Sqrt(float64(sum)))
			if sqrt*sqrt == sum {
				ans++
			}
		}
	}
	return ans
}
