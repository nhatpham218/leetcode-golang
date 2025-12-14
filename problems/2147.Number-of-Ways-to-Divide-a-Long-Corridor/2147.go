package leetcode

// 2147. Number of Ways to Divide a Long Corridor
// https://leetcode.com/problems/number-of-ways-to-divide-a-long-corridor/description/
func numberOfWays(corridor string) int {
	seatCount := 0
	plantCount := 0
	answer := 1
	for _, ch := range corridor {
		if ch == 'S' {
			seatCount++
			if seatCount%2 == 1 && seatCount > 1 {
				answer = (answer * (plantCount + 1)) % 1000000007
			}

			if seatCount%2 == 0 {
				plantCount = 0
			}
		}
		if ch == 'P' {
			plantCount++
		}

	}

	if seatCount%2 != 0 && seatCount > 1 {
		return 0
	}

	return answer
}
