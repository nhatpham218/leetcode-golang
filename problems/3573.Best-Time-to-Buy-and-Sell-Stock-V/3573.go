package leetcode

// 3573. Best Time to Buy and Sell Stock V
// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-v/description/

const (
	noPosition = iota // State 0: No active position
	holding           // State 1: Holding a long position (bought)
	shorting          // State 2: Holding a short position (sold)
)

func maximumProfit(prices []int, k int) int64 {
	numDays := len(prices)

	// memo[day][remainingTransactions][positionState] = max profit
	memo := make([][][]int64, numDays)
	for day := range memo {
		memo[day] = make([][]int64, k+1)
		for transactionsLeft := range memo[day] {
			memo[day][transactionsLeft] = make([]int64, 3)
			for state := range memo[day][transactionsLeft] {
				memo[day][transactionsLeft][state] = -1
			}
		}
	}

	var dfs func(dayIndex, remainingTransactions, positionState int) int64
	dfs = func(dayIndex, remainingTransactions, positionState int) int64 {
		// Base case: no transactions left
		if remainingTransactions == 0 {
			return 0
		}

		// Base case: first day
		if dayIndex == 0 {
			switch positionState {
			case noPosition:
				return 0
			case holding:
				// Buy on first day
				return -int64(prices[0])
			case shorting:
				// Short sell on first day
				return int64(prices[0])
			}
		}

		// Check memoization
		if memo[dayIndex][remainingTransactions][positionState] != -1 {
			return memo[dayIndex][remainingTransactions][positionState]
		}

		currentPrice := int64(prices[dayIndex])
		var maxProfit int64

		switch positionState {
		case noPosition:
			// No position: can stay idle, close a long, or close a short
			stayIdle := dfs(dayIndex-1, remainingTransactions, noPosition)
			closeLong := dfs(dayIndex-1, remainingTransactions, holding) + currentPrice
			closeShort := dfs(dayIndex-1, remainingTransactions, shorting) - currentPrice
			maxProfit = max(stayIdle, max(closeLong, closeShort))

		case holding:
			// Holding long: can keep holding or buy (start new transaction)
			keepHolding := dfs(dayIndex-1, remainingTransactions, holding)
			buyNew := dfs(dayIndex-1, remainingTransactions-1, noPosition) - currentPrice
			maxProfit = max(keepHolding, buyNew)

		case shorting:
			// Holding short: can keep shorting or short sell (start new transaction)
			keepShorting := dfs(dayIndex-1, remainingTransactions, shorting)
			shortSell := dfs(dayIndex-1, remainingTransactions-1, noPosition) + currentPrice
			maxProfit = max(keepShorting, shortSell)
		}

		memo[dayIndex][remainingTransactions][positionState] = maxProfit
		return maxProfit
	}

	return dfs(numDays-1, k, noPosition)
}
