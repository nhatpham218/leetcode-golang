package leetcode

import "math"

// judgePoint24 returns true if it's possible to arrange the four cards to get 24
// using mathematical operations (+, -, *, /) and parentheses
func judgePoint24(cards []int) bool {
	newCards := make([]float64, len(cards))
	for i := range cards {
		newCards[i] = float64(cards[i])
	}
	return check(newCards)
}

func check(cards []float64) bool {
	if len(cards) == 1 {
		return math.Abs(float64(cards[0])-24) < 1e-6
	}

	for i := 0; i < len(cards); i++ {
		for j := i + 1; j < len(cards); j++ {
			newList := make([]float64, 0, len(cards)-1)
			for k := 0; k < len(cards); k++ {
				if k != i && k != j {
					newList = append(newList, cards[k])
				}
			}
			for _, res := range compute(float64(cards[i]), float64(cards[j])) {
				newList = append(newList, res)
				if check(newList) {
					return true
				}
				newList = newList[:len(newList)-1]
			}
		}
	}
	return false

}

func compute(a, b float64) []float64 {
	res := []float64{a + b, a - b, a * b, b - a}
	if b != 0 {
		res = append(res, a/b)
	}
	if a != 0 {
		res = append(res, b/a)
	}
	return res
}
