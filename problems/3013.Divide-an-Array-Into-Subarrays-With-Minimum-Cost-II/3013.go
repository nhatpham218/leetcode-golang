package leetcode

import (
	"github.com/emirpasic/gods/trees/redblacktree"
)

// 3013. Divide an Array Into Subarrays With Minimum Cost II
// https://leetcode.com/problems/divide-an-array-into-subarrays-with-minimum-cost-ii/description/
func minimumCost(nums []int, k int, dist int) int64 {
	n := len(nums)
	need := k - 2 // number of elements we need to select from the window

	// Use pair [value, index] to handle duplicates
	// Sort by value first, then by index
	comparator := func(a, b interface{}) int {
		pa := a.([2]int)
		pb := b.([2]int)
		if pa[0] != pb[0] {
			return pa[0] - pb[0]
		}
		return pa[1] - pb[1]
	}

	// small: max heap of k-2 smallest elements (using tree, we access the largest via Right())
	// large: min heap of remaining elements (using tree, we access the smallest via Left())
	small := redblacktree.NewWith(comparator)
	large := redblacktree.NewWith(comparator)

	var sumSmall int64 = 0

	// Move largest from small to large
	moveSmallToLarge := func() {
		if small.Size() == 0 {
			return
		}
		it := small.Right()
		key := it.Key.([2]int)
		small.Remove(it.Key)
		large.Put(key, nil)
		sumSmall -= int64(key[0])
	}

	// Move smallest from large to small
	moveLargeToSmall := func() {
		if large.Size() == 0 {
			return
		}
		it := large.Left()
		key := it.Key.([2]int)
		large.Remove(it.Key)
		small.Put(key, nil)
		sumSmall += int64(key[0])
	}

	// Balance: ensure small has exactly 'need' elements (if possible)
	balance := func() {
		for small.Size() < need && large.Size() > 0 {
			moveLargeToSmall()
		}
		for small.Size() > need {
			moveSmallToLarge()
		}
	}

	// Add element to the data structure
	add := func(val, idx int) {
		key := [2]int{val, idx}
		if small.Size() < need {
			small.Put(key, nil)
			sumSmall += int64(val)
		} else if small.Size() > 0 {
			largest := small.Right().Key.([2]int)
			if val < largest[0] || (val == largest[0] && idx < largest[1]) {
				small.Put(key, nil)
				sumSmall += int64(val)
				moveSmallToLarge()
			} else {
				large.Put(key, nil)
			}
		} else {
			large.Put(key, nil)
		}
	}

	// Remove element from the data structure
	remove := func(val, idx int) {
		key := [2]int{val, idx}
		if _, found := small.Get(key); found {
			small.Remove(key)
			sumSmall -= int64(val)
		} else {
			large.Remove(key)
		}
		balance()
	}

	// Initialize: for i=1, window is [2, min(1+dist, n-1)]
	right := 1 + dist
	if right > n-1 {
		right = n - 1
	}
	for j := 2; j <= right; j++ {
		add(nums[j], j)
	}
	balance()

	var ans int64 = 1<<63 - 1
	if small.Size() >= need {
		ans = int64(nums[0]) + int64(nums[1]) + sumSmall
	}

	// Iterate i from 2 to n-k+1
	// For each i, window is [i+1, min(i+dist, n-1)]
	for i := 2; i <= n-k+1; i++ {
		// Remove element at index i (was left boundary of previous window)
		remove(nums[i], i)

		// Add new element entering the window (new right boundary)
		newIdx := i + dist
		if newIdx < n {
			add(nums[newIdx], newIdx)
		}
		balance()

		if small.Size() >= need {
			cost := int64(nums[0]) + int64(nums[i]) + sumSmall
			if cost < ans {
				ans = cost
			}
		}
	}

	return ans
}
