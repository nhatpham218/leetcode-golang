package leetcode

// 1018. Binary Prefix Divisible By 5
// https://leetcode.com/problems/binary-prefix-divisible-by-5/
func prefixesDivBy5(nums []int) []bool {
	result := make([]bool, len(nums))
	num := 0
	for i := 0; i < len(nums); i++ {
		num = (num*2 + nums[i]) % 5
		result[i] = num == 0
	}
	return result
}
