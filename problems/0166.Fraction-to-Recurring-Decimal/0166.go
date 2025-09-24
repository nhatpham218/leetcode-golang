package leetcode

import (
	"math"
	"strconv"
)

// 166. Fraction to Recurring Decimal
// https://leetcode.com/problems/fraction-to-recurring-decimal
func fractionToDecimal(numerator int, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	result := ""

	// Handle sign
	if (numerator < 0) != (denominator < 0) {
		result += "-"
	}

	// Work with absolute values
	num := int64(math.Abs(float64(numerator)))
	den := int64(math.Abs(float64(denominator)))

	// Add integer part
	result += strconv.FormatInt(int64(num/den), 10)
	num = num % den

	if num == 0 {
		return result
	}

	// Add decimal point
	result += "."

	// Track remainders and their positions for cycle detection
	remainderMap := make(map[int64]int)

	for num != 0 {
		// If we've seen this remainder before, we have a cycle
		if pos, exists := remainderMap[num]; exists {
			res := result
			return res[:pos] + "(" + res[pos:] + ")"
		}

		// Record this remainder and its position
		remainderMap[num] = len(result)

		num *= 10
		result += strconv.FormatInt(int64(num/den), 10)
		num = num % den
	}

	return result
}
