package leetcode

import "strings"

// 1935. Maximum Number of Words You Can Type
// https://leetcode.com/problems/maximum-number-of-words-you-can-type/

func canBeTypedWords(text string, brokenLetters string) int {
	brokenLettersMap := make(map[rune]bool)
	for _, letter := range brokenLetters {
		brokenLettersMap[letter] = true
	}
	words := strings.Split(text, " ")
	count := 0
	for _, word := range words {
		canBeTyped := true
		for _, letter := range word {
			if brokenLettersMap[letter] {
				canBeTyped = false
				break
			}
		}
		if canBeTyped {
			count++
		}
	}
	return count
}
