package leetcode

import "math/bits"

func makeTheIntegerZero(num1 int, num2 int) int {
	k := 1
	for {
		x := num1 - num2*k
		if x < k {
			return -1
		}
		if k >= bits.OnesCount(uint(x)) {
			return k
		}
		k++
	}
}
