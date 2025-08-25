package leetcode

import "strconv"

// maximum69Number returns the maximum number you can get by changing at most one digit
// (6 becomes 9, and 9 becomes 6) in the given positive integer num
func maximum69Number(num int) int {
	numStrBytes := []byte(strconv.Itoa(num))
	for i := 0; i < len(numStrBytes); i++ {
		if numStrBytes[i] == '6' {
			numStrBytes[i] = '9'
			break
		}
	}
	res, _ := strconv.Atoi(string(numStrBytes))
	return res
}
