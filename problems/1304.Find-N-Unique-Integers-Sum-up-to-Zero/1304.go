package leetcode

// 1304. Find N Unique Integers Sum up to Zero
// Given an integer n, return any array containing n unique integers such that they add up to 0.
func sumZero(n int) []int {
	res := make([]int, n)
	for i := 0; i < n-1; i++ {
		res[i] = i + 1
	}
	// The sum of the first n-1 integers is n * (n - 1) / 2
	// The last element is the negative of the sum of the first n-1 integers
	res[n-1] = -n * (n - 1) / 2
	return res
}
