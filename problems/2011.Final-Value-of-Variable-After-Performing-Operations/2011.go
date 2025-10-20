package leetcode

// 2011. Final Value of Variable After Performing Operations
// https://leetcode.com/problems/final-value-of-variable-after-performing-operations/
func finalValueAfterOperations(operations []string) int {
	result := 0
	for _, operation := range operations {
		if operation == "X++" || operation == "++X" {
			result++
		} else {
			result--
		}
	}
	return result
}
