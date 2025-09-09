package leetcode

func getNoZeroIntegers(n int) []int {
	for a := 1; a < n; a++ {
		b := n - a
		if hasNoZero(a) && hasNoZero(b) {
			return []int{a, b}
		}
	}
	return nil
}

func hasNoZero(num int) bool {
	if num == 0 {
		return false
	}
	for num > 0 {
		if num%10 == 0 {
			return false
		}
		num /= 10
	}
	return true
}
