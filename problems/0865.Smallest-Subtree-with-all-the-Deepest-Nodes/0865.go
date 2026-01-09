package leetcode

// 865. Smallest Subtree with all the Deepest Nodes
// https://leetcode.com/problems/smallest-subtree-with-all-the-deepest-nodes/description/

// TreeNode Definition for a binary tree node.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	nodesToLevel := make(map[*TreeNode]int)
	maxDepth := dfs(root, nodesToLevel)
	queue := []*TreeNode{root}
	depth := 1
	for len(queue) > 0 {
		levelSize := len(queue)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			if nodesToLevel[node.Left] == maxDepth-depth && nodesToLevel[node.Right] == maxDepth-depth {
				return node
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[levelSize:]
		depth++
	}
	return nil
}

func dfs(node *TreeNode, nodesToLevel map[*TreeNode]int) int {
	if node == nil {
		return 0
	}
	leftDepth := dfs(node.Left, nodesToLevel)
	rightDepth := dfs(node.Right, nodesToLevel)
	nodesToLevel[node] = max(leftDepth, rightDepth) + 1
	return nodesToLevel[node]
}
