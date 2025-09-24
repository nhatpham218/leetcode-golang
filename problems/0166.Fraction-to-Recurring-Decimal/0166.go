package leetcode

import (
	"math"
	"strconv"
	"strings"
)

// 166. Fraction to Recurring Decimal
// https://leetcode.com/problems/fraction-to-recurring-decimal
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	var result strings.Builder

	// Handle sign
	if (numerator < 0) != (denominator < 0) {
		result.WriteByte('-')
	}

	// Work with absolute values
	num := int64(math.Abs(float64(numerator)))
	den := int64(math.Abs(float64(denominator)))

	// Add integer part
	result.WriteString(strconv.FormatInt(int64(num/den), 10))
	num = num % den

	if num == 0 {
		return result.String()
	}

	// Add decimal point
	result.WriteByte('.')

	// Track remainders and their positions for cycle detection
	remainderMap := make(map[int64]int)

	for num != 0 {
		// If we've seen this remainder before, we have a cycle
		if pos, exists := remainderMap[num]; exists {
			res := result.String()
			return res[:pos] + "(" + res[pos:] + ")"
		}

		// Record this remainder and its position
		remainderMap[num] = result.Len()

		num *= 10
		result.WriteByte(byte('0' + num/den))
		num = num % den
	}

	return result.String()
}
