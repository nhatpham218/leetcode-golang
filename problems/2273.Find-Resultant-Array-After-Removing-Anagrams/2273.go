package leetcode

import "sort"

// 2273. Find Resultant Array After Removing Anagrams
// https://leetcode.com/problems/find-resultant-array-after-removing-anagrams/
func removeAnagrams(words []string) []string {
	if len(words) == 0 {
		return []string{}
	}

	res := []string{words[0]}

	for i := 1; i < len(words); i++ {
		if !isAnagram(res[len(res)-1], words[i]) {
			res = append(res, words[i])
		}
	}

	return res
}

func isAnagram(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	return sortStr(s1) == sortStr(s2)
}

func sortStr(s string) string {
	runes := []rune(s)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}
