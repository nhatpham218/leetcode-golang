package leetcode

import (
	"math"
	"sort"
)

// 1200. Minimum Absolute Difference
// https://leetcode.com/problems/minimum-absolute-difference/description/
func minimumAbsDifference(arr []int) [][]int {
	sort.Ints(arr)
	res := [][]int{}

	minDiff := math.MaxInt32
	for i := 1; i < len(arr); i++ {
		if minDiff > arr[i]-arr[i-1] {
			minDiff = arr[i] - arr[i-1]
			res = [][]int{}
		}
		if minDiff == arr[i]-arr[i-1] {
			res = append(res, []int{arr[i-1], arr[i]})
		}
	}

	return res
}
