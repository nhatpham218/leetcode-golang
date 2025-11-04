package leetcode

import "sort"

// 3318. Find X-Sum of All K-Long Subarrays I
// https://leetcode.com/problems/find-x-sum-of-all-k-long-subarrays-i/
func findXSum(nums []int, k int, x int) []int {
	res := make([]int, len(nums)-k+1)
	count := make(map[int]int)
	for i := 0; i < k; i++ {
		count[nums[i]]++
	}
	res[0] = sum(count, x)

	for i := k; i < len(nums); i++ {
		count[nums[i]]++
		count[nums[i-k]]--
		if count[nums[i-k]] == 0 {
			delete(count, nums[i-k])
		}
		res[i-k+1] = sum(count, x)
	}
	return res
}

func sum(count map[int]int, x int) int {
	keys := make([]int, 0, len(count))
	for k := range count {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		if count[keys[i]] == count[keys[j]] {
			return keys[i] > keys[j]
		}
		return count[keys[i]] > count[keys[j]]
	})

	sum := 0
	for i := 0; i < min(x, len(keys)); i++ {
		sum += keys[i] * count[keys[i]]
	}
	return sum
}
