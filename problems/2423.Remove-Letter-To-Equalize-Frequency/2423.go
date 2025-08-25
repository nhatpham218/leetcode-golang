package leetcode

// 2423. Remove Letter To Equalize Frequency
func equalFrequency(word string) bool {
	count := make(map[rune]int)
	countFreq := make(map[int]int)
	maxFreq := 0
	for i := 0; i < len(word); i++ {
		count[rune(word[i])]++
		countFreq[count[rune(word[i])]]++
		countFreq[count[rune(word[i])]-1]--
		if maxFreq < count[rune(word[i])] {
			maxFreq = count[rune(word[i])]
		}
	}

	if maxFreq == 1 {
		return true
	}

	if maxFreq*countFreq[maxFreq] == len(word)-1 {
		return true
	}

	if (maxFreq-1)*(countFreq[maxFreq-1]+1) == len(word)-1 {
		return true
	}

	return false
}
