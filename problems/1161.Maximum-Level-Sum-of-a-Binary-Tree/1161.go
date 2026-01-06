package leetcode

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1161. Maximum Level Sum of a Binary Tree
// https://leetcode.com/problems/maximum-level-sum-of-a-binary-tree/
func maxLevelSum(root *TreeNode) int {
	queue := []*TreeNode{root}
	maxSum := math.MinInt
	maxLevel := 0
	level := 0
	for len(queue) > 0 {
		level++
		sum := 0
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]
			sum += node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if sum > maxSum {
			maxSum = sum
			maxLevel = level
		}
	}
	return maxLevel
}
