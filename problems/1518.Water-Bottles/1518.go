package leetcode

// 1518. Water Bottles
// https://leetcode.com/problems/water-bottles/
func numWaterBottles(numBottles int, numExchange int) int {
	return numBottles + (numBottles-1)/(numExchange-1)
}
