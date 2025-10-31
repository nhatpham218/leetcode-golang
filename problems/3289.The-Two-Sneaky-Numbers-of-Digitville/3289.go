package leetcode

// 3289. The Two Sneaky Numbers of Digitville
// https://leetcode.com/problems/the-two-sneaky-numbers-of-digitville/
func getSneakyNumbers(nums []int) []int {
	result := []int{}
	exist := make([]bool, 100)
	for _, num := range nums {
		if exist[num] {
			result = append(result, num)
		}
		exist[num] = true
	}
	return result
}
