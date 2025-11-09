package leetcode

// 2169. Count Operations to Obtain Zero
// https://leetcode.com/problems/count-operations-to-obtain-zero/
func countOperations(num1 int, num2 int) int {
	res := 0 //
	for num1 != 0 && num2 != 0 {
		res += num1 / num2
		num1 %= num2
		num1, num2 = num2, num1
	}
	return res
}
