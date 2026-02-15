package leetcode

import (
	"fmt"
	"strings"
)

// 67. Add Binary
// https://leetcode.com/problems/add-binary/description/
func addBinary(a string, b string) string {
	result := ""
	carry := 0
	// add padding to the shorter string
	if len(a) < len(b) {
		a = strings.Repeat("0", len(b)-len(a)) + a
	} else {
		b = strings.Repeat("0", len(a)-len(b)) + b
	}
	// add binary
	for i := len(a) - 1; i >= 0; i-- {
		sum := carry
		sum += int(a[i]-'0') + int(b[i]-'0')
		carry = sum / 2
		result = fmt.Sprintf("%d%s", sum%2, result)
	}

	if carry == 1 {
		result = "1" + result
	}
	return result
}
