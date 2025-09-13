package leetcode

import "slices"

// 3541. Find Most Frequent Vowel and Consonant
// https://leetcode.com/problems/find-most-frequent-vowel-and-consonant/

func maxFreqSum(s string) int {
	mapCount := make(map[rune]int)
	vowels := []rune{'a', 'e', 'i', 'o', 'u'}
	maxVowelCount := 0
	maxConsonantCount := 0
	for _, char := range s {
		mapCount[char]++
		if slices.Contains(vowels, char) {
			maxVowelCount = max(maxVowelCount, mapCount[char])
		} else {
			maxConsonantCount = max(maxConsonantCount, mapCount[char])
		}
	}
	return maxVowelCount + maxConsonantCount

}
