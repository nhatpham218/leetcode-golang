package leetcode

// 717. 1-bit and 2-bit Characters
// https://leetcode.com/problems/1-bit-and-2-bit-characters/
func isOneBitCharacter(bits []int) bool {
	if bits[len(bits)-1] != 0 {
		return false
	}
	count1 := 0
	for i := len(bits) - 2; i >= 0 && bits[i] == 1; i-- {
		count1++
	}
	return count1%2 == 0
}
