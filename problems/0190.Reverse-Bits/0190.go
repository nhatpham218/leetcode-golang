package leetcode

// 190. Reverse Bits
// https://leetcode.com/problems/reverse-bits/description/
func reverseBits(n uint32) uint32 {
	reversed := uint32(0)
	for i := 0; i < 32; i++ {
		reversed <<= 1
		reversed |= n & 1
		n >>= 1
	}
	return reversed
}
