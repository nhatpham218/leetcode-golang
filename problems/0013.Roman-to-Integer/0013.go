package leetcode

import "strings"

// 0013. Roman to Integer
// https://leetcode.com/problems/roman-to-integer/description/
func romanToInt(s string) int {
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := 0
	for i := range values {
		for strings.HasPrefix(s, roman[i]) {
			s = s[len(roman[i]):]
			result += values[i]
		}
	}
	return result
}
