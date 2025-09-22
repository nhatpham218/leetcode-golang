package leetcode

// 3005. Count Elements With Maximum Frequency
// You are given an array nums consisting of positive integers.
// Return the total frequencies of elements in nums such that those elements all have the maximum frequency.
// The frequency of an element is the number of occurrences of that element in the array.
func maxFrequencyElements(nums []int) int {
	count := make([]int, 101)
	maxFreq := 0
	countMaxFreq := 0
	for _, num := range nums {
		count[num]++
		if count[num] > maxFreq {
			maxFreq = count[num]
			countMaxFreq = 1
		} else if count[num] == maxFreq {
			countMaxFreq++
		}
	}
	return countMaxFreq * maxFreq
}
