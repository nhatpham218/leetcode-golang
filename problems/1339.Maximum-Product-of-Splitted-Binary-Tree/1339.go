package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

const MOD = 1000000007

// 1339. Maximum Product of Splitted Binary Tree
// https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/
func maxProduct(root *TreeNode) int {
	dfsSum(root)
	total := root.Val
	queue := []*TreeNode{root}
	maxProduct := 0
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		product := (total - node.Val) * node.Val
		if product > maxProduct {
			maxProduct = product
		}
		if node.Left != nil {
			queue = append(queue, node.Left)
		}
		if node.Right != nil {
			queue = append(queue, node.Right)
		}
	}
	return maxProduct % MOD
}

func dfsSum(node *TreeNode) int {
	if node == nil {
		return 0
	}
	sum := node.Val + dfsSum(node.Left) + dfsSum(node.Right)
	node.Val = sum
	return sum
}
