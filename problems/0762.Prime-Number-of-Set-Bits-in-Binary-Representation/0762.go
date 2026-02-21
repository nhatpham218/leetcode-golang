package leetcode

import (
	"math/bits"
	"slices"
)

// 0762. Prime Number of Set Bits in Binary Representation
// https://leetcode.com/problems/prime-number-of-set-bits-in-binary-representation/description/
func countPrimeSetBits(left int, right int) int {
	count := 0
	for i := left; i <= right; i++ {
		if isPrime(bits.OnesCount(uint(i))) {
			count++
		}
	}
	return count
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	primes := []int{2, 3, 5, 7, 11, 13, 17, 19}
	return slices.ContainsFunc(primes, func(p int) bool {
		return p == n
	})
}
