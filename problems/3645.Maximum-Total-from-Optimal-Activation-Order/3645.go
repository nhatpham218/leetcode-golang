package leetcode

// 3645. Maximum Total from Optimal Activation Order
// You are given two integer arrays value and limit, both of length n.
// Initially, all elements are inactive. You may activate them in any order.
// To activate an inactive element at index i, the number of currently active elements must be strictly less than limit[i].
// When you activate the element at index i, it adds value[i] to the total activation value (i.e., the sum of value[i] for all elements that have undergone activation operations).
// After each activation, if the number of currently active elements becomes x, then all elements j with limit[j] <= x become permanently inactive, even if they are already active.
// Return the maximum total you can obtain by choosing the activation order optimally.
func maxTotal(value []int, limit []int) int64 {
	return 0
}
