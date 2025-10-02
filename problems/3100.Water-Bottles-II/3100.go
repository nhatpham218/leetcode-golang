package leetcode

// 3100. Water Bottles II
// https://leetcode.com/problems/water-bottles-ii/
func maxBottlesDrunk(numBottles int, numExchange int) int {
	drunk := 0
	empty := 0
	exchange := numExchange

	for {
		if numBottles > 0 {
			drunk += numBottles
			empty += numBottles
			numBottles = 0
		}

		if empty >= exchange {
			numBottles = 1
			empty -= exchange
			exchange++
		} else {
			break
		}
	}

	return drunk
}
