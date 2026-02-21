package leetcode

// 0012. Integer to Roman
// https://leetcode.com/problems/integer-to-roman/description/
func intToRoman(num int) string {
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	values := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	result := ""
	for i := range values {
		for num >= values[i] {
			num -= values[i]
			result += roman[i]
		}
	}
	return result
}
