package leetcode

// 1390. Four Divisors
// https://leetcode.com/problems/four-divisors/description/
func sumFourDivisors(nums []int) int {
	total := 0
	for _, num := range nums {
		divisors := getDivisors(num)
		if len(divisors) == 2 {
			total += divisors[0] + divisors[1] + 1 + num
		}
	}
	return total
}

func getDivisors(num int) []int {
	divisors := make([]int, 4)
	numDivisors := 0
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			divisors[numDivisors] = i
			numDivisors++
			if i != num/i {
				divisors[numDivisors] = num / i
				numDivisors++
			}
		}
		if numDivisors > 2 {
			break
		}
	}
	return divisors[:numDivisors]
}
