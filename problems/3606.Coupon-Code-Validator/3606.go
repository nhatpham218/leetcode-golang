package leetcode

import "sort"

// 3606. Coupon Code Validator
// https://leetcode.com/problems/coupon-code-validator/description/
func validateCoupons(code []string, businessLine []string, isActive []bool) []string {
	priority := map[string]int{
		"electronics": 0,
		"grocery":     1,
		"pharmacy":    2,
		"restaurant":  3,
	}

	indices := []int{}
	for i := range code {
		if isValidCoupon(code[i], businessLine[i], isActive[i]) {
			indices = append(indices, i)
		}
	}

	sort.Slice(indices, func(i, j int) bool {
		bi, bj := businessLine[indices[i]], businessLine[indices[j]]
		if bi == bj {
			return code[indices[i]] < code[indices[j]]
		}
		return priority[bi] < priority[bj]
	})

	res := make([]string, len(indices))
	for i, idx := range indices {
		res[i] = code[idx]
	}
	return res
}

func isValidCoupon(code string, businessLine string, isActive bool) bool {
	if businessLine != "electronics" && businessLine != "grocery" && businessLine != "pharmacy" && businessLine != "restaurant" {
		return false
	}
	if !isActive {
		return false
	}

	if len(code) == 0 {
		return false
	}

	for _, ch := range code {
		if ch != '_' && (ch < 'a' || ch > 'z') && (ch < 'A' || ch > 'Z') && (ch < '0' || ch > '9') {
			return false
		}
	}

	return true
}
