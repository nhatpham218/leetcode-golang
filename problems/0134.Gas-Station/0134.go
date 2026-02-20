package leetcode

// 0134. Gas Station
// https://leetcode.com/problems/gas-station/description/
func canCompleteCircuit(gas []int, cost []int) int {
	var totalGas, totalCost, tank int
	start := 0
	for i := range gas {
		totalGas += gas[i]
		totalCost += cost[i]
		tank += gas[i] - cost[i]
		if tank < 0 {
			start = i + 1
			tank = 0
		}
	}
	if totalGas < totalCost {
		return -1
	}
	return start
}
