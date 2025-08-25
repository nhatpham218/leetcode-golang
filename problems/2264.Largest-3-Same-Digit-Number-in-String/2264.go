package leetcode

// largestGoodInteger finds the largest 3-same-digit number in the string
func largestGoodInteger(num string) string {
	maxGood := ""

	// Check each possible 3-digit substring
	for i := 0; i <= len(num)-3; i++ {
		substr := num[i : i+3]

		// Check if all three digits are the same
		if substr[0] == substr[1] && substr[1] == substr[2] {
			// If this is the first good integer found, or if it's larger than the current max
			if maxGood == "" || substr > maxGood {
				maxGood = substr
			}
		}
	}

	return maxGood
}
