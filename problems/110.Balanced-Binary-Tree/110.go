package leetcode

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 110. Balanced Binary Tree
// https://leetcode.com/problems/balanced-binary-tree/description/
func isBalanced(root *TreeNode) bool {
	queue := []*TreeNode{root}
	for len(queue) > 0 && root != nil {
		node := queue[0]
		queue = queue[1:]
		leftHeight := getHeight(node.Left)
		rightHeight := getHeight(node.Right)
		if abs(leftHeight-rightHeight) > 1 {
			return false
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return true
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
