package leetcode

// 1262. Greatest Sum Divisible by Three
// Given an integer array nums, return the maximum possible sum of elements of the array such that it is divisible by three.
func maxSumDivThree(nums []int) int {
	f := []int{0, -1, -1}

	for _, num := range nums {
		temp := make([]int, 3)
		for i := range 3 {
			temp[(i+num)%3] = max(f[(i+num)%3], f[i]+num)
		}
		f = temp

	}

	return f[0]
}
