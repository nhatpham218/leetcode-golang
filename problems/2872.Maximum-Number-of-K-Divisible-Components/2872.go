package leetcode

const maxNodes = 30000

var adjXor [maxNodes]uint16
var degree [maxNodes]uint8
var stk [maxNodes / 2]uint16

func isZeroMask(v uint32) uint32 {
	return ((v | -v) >> 31) - 1
}

// 2872. Maximum Number of K-Divisible Components
// https://leetcode.com/problems/maximum-number-of-k-divisible-components/
func maxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	if n == 1 {
		return 1
	}
	for i := range n {
		adjXor[i] = 0
		degree[i] = 0
	}
	for _, e := range edges {
		adjXor[e[0]] ^= uint16(e[1])
		adjXor[e[1]] ^= uint16(e[0])
		degree[e[0]]++
		degree[e[1]]++
	}
	var stkTop uint16
	var comps uint32
	pushStack := func(i uint16) {
		mask := uint16(isZeroMask(uint32(degree[i] ^ 1)))
		stk[stkTop] = (stk[stkTop] &^ mask) | (i & mask)
		stkTop += 1 & mask
	}
	for i := range uint16(n) {
		pushStack(i)
	}
	for stkTop != 0 {
		stkTop--
		node := stk[stkTop]
		nbr := adjXor[node]
		adjXor[nbr] ^= node
		degree[nbr]--
		pushStack(nbr)
		values[node] %= k
		comps += 1 & isZeroMask(uint32(values[node]))
		values[nbr] += values[node]
	}
	return int(comps)
}
