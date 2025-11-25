package leetcode

// 1015. Smallest Integer Divisible by K
// https://leetcode.com/problems/smallest-integer-divisible-by-k/
func smallestRepunitDivByK(k int) int {
	if k%2 == 0 || k%5 == 0 {
		return -1
	}
	remainder := 0
	for length_N := 1; length_N <= k; length_N++ {
		remainder = (remainder*10 + 1) % k
		if remainder == 0 {
			return length_N
		}
	}
	return -1

}
