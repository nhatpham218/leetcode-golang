package leetcode

// 1578. Minimum Time to Make Rope Colorful
// https://leetcode.com/problems/minimum-time-to-make-rope-colorful/
func minCost(colors string, neededTime []int) int {
	totalCost := 0
	start := 0

	// For each group of consecutive balloons of the same color:
	// Keep the most expensive balloon and remove the others, add the cost to the total cost
	for start < len(colors) {
		end := start
		groupSum := 0
		maxTime := 0

		for end < len(colors) && colors[end] == colors[start] {
			groupSum += neededTime[end]
			if neededTime[end] > maxTime {
				maxTime = neededTime[end]
			}
			end++
		}

		if end-start > 1 {
			totalCost += groupSum - maxTime
		}

		start = end
	}

	return totalCost
}
