package leetcode

import "math"

// 2654. Minimum Number of Operations to Make All Array Elements Equal to 1
// https://leetcode.com/problems/minimum-number-of-operations-to-make-all-array-elements-equal-to-1/
func minOperations(nums []int) int {
	count1 := 0
	g := 0
	for _, num := range nums {
		if num == 1 {
			count1++
		}
		g = gcd(g, num)
	}
	if count1 != 0 {
		return len(nums) - count1
	}
	if g != 1 {
		return -1
	}

	result := math.MaxInt64
	for i := 0; i < len(nums); i++ {
		gcd_num := nums[i]
		for j := i + 1; j < len(nums); j++ {
			gcd_num = gcd(gcd_num, nums[j])
			if gcd_num == 1 {
				result = min(result, j-i+len(nums)-1)
				break
			}
		}
	}
	return result
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
