package leetcode

func sortVowels(s string) string {
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
	}

	// Count frequency of each vowel
	count := [128]int{}
	for i := 0; i < len(s); i++ {
		if isVowel(s[i]) {
			count[s[i]]++
		}
	}

	// Build result by replacing vowels in sorted order
	result := []byte(s)
	vowelIdx := 0

	// Iterate through ASCII values to get sorted order
	for ascii := 0; ascii < 128; ascii++ {
		for count[ascii] > 0 {
			// Find next vowel position
			for vowelIdx < len(s) && !isVowel(s[vowelIdx]) {
				vowelIdx++
			}
			if vowelIdx < len(s) {
				result[vowelIdx] = byte(ascii)
				vowelIdx++
				count[ascii]--
			}
		}
	}

	return string(result)
}
