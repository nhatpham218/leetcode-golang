package leetcode

import "slices"

// 3227. Vowels Game in a String
// https://leetcode.com/problems/vowels-game-in-a-string/

func doesAliceWin(s string) bool {
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	for _, char := range s {
		if slices.Contains(vowels, char) {
			return true
		}
	}
	return false
}
