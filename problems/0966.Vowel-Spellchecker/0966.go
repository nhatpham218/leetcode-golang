package leetcode

import (
	"strings"
)

// 0966. Vowel Spellchecker
// https://leetcode.com/problems/vowel-spellchecker/

func spellchecker(wordlist []string, queries []string) []string {
	perfectMatch := make(map[string]string)
	capitalizationMatch := make(map[string]string)
	vowelMatch := make(map[string]string)

	for _, word := range wordlist {
		perfectMatch[word] = word
		lower := strings.ToLower(word)
		if _, exists := capitalizationMatch[lower]; !exists {
			capitalizationMatch[lower] = word
		}

		vowelKey := replaceVowels(lower)
		if _, exists := vowelMatch[vowelKey]; !exists {
			vowelMatch[vowelKey] = word
		}
	}

	ans := make([]string, len(queries))
	for i, query := range queries {
		if match, exists := perfectMatch[query]; exists {
			ans[i] = match
		} else {
			lower := strings.ToLower(query)
			if match, exists := capitalizationMatch[lower]; exists {
				ans[i] = match
			} else if match, exists := vowelMatch[replaceVowels(lower)]; exists {
				ans[i] = match
			}
		}
	}
	return ans
}

func replaceVowels(word string) string {
	result := make([]byte, len(word))
	for i, ch := range []byte(word) {
		if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
			result[i] = '*'
		} else {
			result[i] = ch
		}
	}
	return string(result)
}
