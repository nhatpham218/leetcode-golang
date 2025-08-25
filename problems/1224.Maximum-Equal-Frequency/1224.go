package leetcode

// 1224. Maximum Equal Frequency
func maxEqualFreq(nums []int) int {
	longest := 0
	count := make(map[int]int)
	countFreq := make(map[int]int)
	maxFreq := 0
	for i := 0; i < len(nums); i++ {
		count[nums[i]]++
		countFreq[count[nums[i]]-1]--
		countFreq[count[nums[i]]]++

		if count[nums[i]] > maxFreq {
			maxFreq = count[nums[i]]
		}
		if maxFreq == 1 || maxFreq*countFreq[maxFreq] == i || (maxFreq-1)*(countFreq[maxFreq-1]+1) == i {
			longest = i + 1
		}
	}
	return longest
}
