package leetcode

func zeroFilledSubarray(nums []int) int64 {
	count := 0
	res := 0
	for i := 0; i <= len(nums); i++ {
		if i < len(nums) && nums[i] == 0 {
			count++
		} else {
			res += count * (count + 1) / 2
			count = 0
		}
	}
	return int64(res)
}
