package leetcode

// 3658. GCD of Odd and Even Sums
// You are given an integer n. Your task is to compute the GCD (greatest common divisor) of two values:
// sumOdd: the sum of the first n odd numbers.
// sumEven: the sum of the first n even numbers.
// Return the GCD of sumOdd and sumEven.
func gcdOfOddEvenSums(n int) int {
	sumOdd := n * n
	sumEven := n * (n + 1)
	return gcd(sumOdd, sumEven)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
