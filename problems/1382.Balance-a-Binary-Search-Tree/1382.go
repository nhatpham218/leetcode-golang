package leetcode

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1382. Balance a Binary Search Tree
// https://leetcode.com/problems/balance-a-binary-search-tree/description/
func balanceBST(root *TreeNode) *TreeNode {
	//Convert the tree to a sorted array using an in-order traversal.
	values := []int{}
	inOrderTraversal(root, &values)
	return buildBalancedBST(values)
}

func buildBalancedBST(values []int) *TreeNode {
	if len(values) == 0 {
		return nil
	}
	mid := len(values) / 2
	root := &TreeNode{Val: values[mid]}
	root.Left = buildBalancedBST(values[:mid])
	root.Right = buildBalancedBST(values[mid+1:])
	return root
}

func inOrderTraversal(root *TreeNode, values *[]int) {
	if root == nil {
		return
	}
	inOrderTraversal(root.Left, values)
	*values = append(*values, root.Val)
	inOrderTraversal(root.Right, values)
}
