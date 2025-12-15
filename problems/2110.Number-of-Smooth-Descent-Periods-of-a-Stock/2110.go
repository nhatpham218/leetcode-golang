package leetcode

// 2110. Number of Smooth Descent Periods of a Stock
// https://leetcode.com/problems/number-of-smooth-descent-periods-of-a-stock/description/
func getDescentPeriods(prices []int) int64 {
	start := 0
	end := 0
	answer := 0
	for i := 1; i < len(prices); i++ {
		if prices[i]-prices[i-1] == -1 {
			end = i
		} else {
			answer += (end - start + 1) * (end - start + 2) / 2

			start = i
			end = i
		}
	}

	answer += (end - start + 1) * (end - start + 2) / 2
	return int64(answer)
}
