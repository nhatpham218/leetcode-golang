package leetcode

// 3354. Make Array Elements Equal to Zero
// https://leetcode.com/problems/make-array-elements-equal-to-zero/
func countValidSelections(nums []int) int {
	total := 0
	for _, num := range nums {
		total += num
	}

	leftSum := 0
	count := 0
	for _, num := range nums {
		rightSum := total - leftSum - num
		if num == 0 {
			if leftSum == rightSum {
				count += 2
			} else if leftSum == rightSum-1 || leftSum-1 == rightSum {
				count++
			}
		}
		leftSum += num
	}
	return count
}
