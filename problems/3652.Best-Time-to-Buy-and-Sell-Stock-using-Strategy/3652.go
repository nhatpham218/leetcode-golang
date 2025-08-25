package leetcode

// 3652. Best Time to Buy and Sell Stock using Strategy
func maxProfit(prices []int, strategy []int, k int) int64 {
	if k%2 != 0 {
		return -1
	}

	maxProfit := 0
	for i := range len(prices) {
		if strategy[i] == 0 {
			continue
		}
		maxProfit += strategy[i] * prices[i]
	}

	currentProfit := maxProfit
	newStrategy := make([]int, len(strategy))
	for i := range len(strategy) {
		currentProfit -= strategy[i] * prices[i]
		if i < k/2 {
			newStrategy[i] = 0
		} else if i < k {
			newStrategy[i] = 1
		} else {
			newStrategy[i] = strategy[i]
		}
		currentProfit += newStrategy[i] * prices[i]
	}
	if currentProfit > maxProfit {
		maxProfit = currentProfit
	}

	for i := k; i < len(prices); i++ {
		currentProfit -= newStrategy[i-k] * prices[i-k]
		currentProfit -= newStrategy[i] * prices[i]
		currentProfit -= newStrategy[i-k/2] * prices[i-k/2]
		newStrategy[i-k/2] = 0
		newStrategy[i] = 1
		newStrategy[i-k] = strategy[i-k]

		currentProfit += newStrategy[i-k] * prices[i-k]
		currentProfit += newStrategy[i] * prices[i]
		currentProfit += newStrategy[i-k/2] * prices[i-k/2]

		if currentProfit > maxProfit {
			maxProfit = currentProfit
		}

	}

	return int64(maxProfit)
}
