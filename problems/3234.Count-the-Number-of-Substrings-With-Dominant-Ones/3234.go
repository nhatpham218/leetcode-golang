package leetcode

// 3234. Count the Number of Substrings With Dominant Ones
// https://leetcode.com/problems/count-the-number-of-substrings-with-dominant-ones/
func numberOfSubstrings(s string) int {
	n := len(s)

	lastZeroPos := make([]int, n+1)
	lastZeroPos[0] = -1

	for i := 0; i < n; i++ {
		if i == 0 || s[i-1] == '0' {
			lastZeroPos[i+1] = i
		} else {
			lastZeroPos[i+1] = lastZeroPos[i]
		}
	}

	result := 0

	for endIdx := 1; endIdx <= n; endIdx++ {
		zeroCount := 0
		if s[endIdx-1] == '0' {
			zeroCount = 1
		}

		startIdx := endIdx

		for startIdx > 0 && zeroCount*zeroCount <= n {
			oneCount := (endIdx - lastZeroPos[startIdx]) - zeroCount

			if zeroCount*zeroCount <= oneCount {
				validCount := startIdx - lastZeroPos[startIdx]

				maxValid := oneCount - zeroCount*zeroCount + 1
				if maxValid < validCount {
					validCount = maxValid
				}

				result += validCount
			}

			startIdx = lastZeroPos[startIdx]
			zeroCount++
		}
	}

	return result
}
