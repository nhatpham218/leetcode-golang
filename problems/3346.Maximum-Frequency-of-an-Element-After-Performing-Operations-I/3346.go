package leetcode

// 3346. Maximum Frequency of an Element After Performing Operations I
// https://leetcode.com/problems/maximum-frequency-of-an-element-after-performing-operations-i/
func maxFrequency(nums []int, k int, numOperations int) int {
	maxNum := 0
	for _, num := range nums {
		if num > maxNum {
			maxNum = num
		}
	}

	size := maxNum + k + 2
	freq := make([]int, size)
	for _, x := range nums {
		freq[x]++
	}

	prefixSum := make([]int, size)
	prefixSum[0] = freq[0]
	for i := 1; i < size; i++ {
		prefixSum[i] = prefixSum[i-1] + freq[i]
	}

	maxFreq := 0
	for target := 0; target < size; target++ {
		if freq[target] == 0 && numOperations == 0 {
			continue
		}
		left := max(0, target-k)
		right := min(size-1, target+k)
		total := prefixSum[right]
		if left > 0 {
			total -= prefixSum[left-1]
		}
		canConvert := total - freq[target]
		currFreq := freq[target] + min(numOperations, canConvert)
		if currFreq > maxFreq {
			maxFreq = currFreq
		}
	}
	return maxFreq
}
