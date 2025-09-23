package leetcode

import (
	"strconv"
	"strings"
)

// 165. Compare Version Numbers
// https://leetcode.com/problems/compare-version-numbers
func compareVersion(version1 string, version2 string) int {
	version1Parts := strings.Split(version1, ".")
	version2Parts := strings.Split(version2, ".")
	maxLen := max(len(version1Parts), len(version2Parts))
	for i := 0; i < maxLen; i++ {
		version1PartInt, version2PartInt := 0, 0
		if i < len(version1Parts) {
			version1PartInt, _ = strconv.Atoi(version1Parts[i])
		}
		if i < len(version2Parts) {
			version2PartInt, _ = strconv.Atoi(version2Parts[i])
		}
		if version1PartInt > version2PartInt {
			return 1
		}
		if version1PartInt < version2PartInt {
			return -1
		}
	}
	return 0
}
