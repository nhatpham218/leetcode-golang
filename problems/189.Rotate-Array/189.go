package leetcode

// 189. Rotate Array
// https://leetcode.com/problems/rotate-array/description/

// Solution 1: Reverse - O(n) time, O(1) space
func rotate(nums []int, k int) {
	n := len(nums)
	k = k % n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, l, r int) {
	for l < r {
		nums[l], nums[r] = nums[r], nums[l]
		l++
		r--
	}
}

// Solution 2: Extra array - O(n) time, O(n) space
// func rotate(nums []int, k int) {
// 	k = k % len(nums)
// 	tmp := append(nums[len(nums)-k:], nums[:len(nums)-k]...)
// 	copy(nums, tmp)
// }

// Solution 3: Cyclic replacements - O(n) time, O(1) space
// func rotate(nums []int, k int) {
// 	n := len(nums)
// 	k = k % n
// 	count := 0
// 	for start := 0; count < n; start++ {
// 		cur := start
// 		prev := nums[cur]
// 		for {
// 			next := (cur + k) % n
// 			nums[next], prev = prev, nums[next]
// 			cur = next
// 			count++
// 			if cur == start {
// 				break
// 			}
// 		}
// 	}
// }
