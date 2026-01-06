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
		newQueue := []*TreeNode{}
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			sum += node.Val
			if node.Left != nil {
				newQueue = append(newQueue, node.Left)
			}
			if node.Right != nil {
				newQueue = append(newQueue, node.Right)
			}
		}
		if sum > maxSum {
			maxSum = sum
			maxLevel = level
		}
		queue = newQueue
	}
	return maxLevel
}
