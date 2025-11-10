package leetcode

// 3542. Minimum Operations to Convert All Elements to Zero
// https://leetcode.com/problems/minimum-operations-to-convert-all-elements-to-zero/
func minOperations(nums []int) int {
	stack := []int{}
	ops := 0

	for _, num := range nums {
		for len(stack) > 0 && stack[len(stack)-1] > num {
			stack = stack[:len(stack)-1]
		}

		if num == 0 {
			continue
		}

		if len(stack) == 0 || stack[len(stack)-1] < num {
			ops++
			stack = append(stack, num)
		}
	}

	return ops
}
