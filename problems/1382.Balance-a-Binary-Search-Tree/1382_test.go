package leetcode

import (
	"fmt"
	"testing"
)

type question1382 struct {
	name string
	para1382
	ans1382
}

type para1382 struct {
	root []interface{}
}

type ans1382 struct {
	values []int // expected in-order traversal (sorted values for BST)
}

func buildTree(arr []interface{}) *TreeNode {
	if len(arr) == 0 || arr[0] == nil {
		return nil
	}
	root := &TreeNode{Val: arr[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(arr) {
		node := queue[0]
		queue = queue[1:]
		if i < len(arr) && arr[i] != nil {
			node.Left = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Left)
		}
		i++
		if i < len(arr) && arr[i] != nil {
			node.Right = &TreeNode{Val: arr[i].(int)}
			queue = append(queue, node.Right)
		}
		i++
	}
	return root
}

func inOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var result []int
	var traverse func(*TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		result = append(result, node.Val)
		traverse(node.Right)
	}
	traverse(root)
	return result
}

func getHeight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(getHeight(root.Left), getHeight(root.Right)) + 1
}

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftH := getHeight(root.Left)
	rightH := getHeight(root.Right)
	if abs(leftH-rightH) > 1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func Test_Problem1382(t *testing.T) {

	qs := []question1382{
		{
			name: "example 1",
			para1382: para1382{
				root: []interface{}{1, nil, 2, nil, 3, nil, 4},
			},
			ans1382: ans1382{values: []int{1, 2, 3, 4}},
		},
		{
			name: "example 2",
			para1382: para1382{
				root: []interface{}{2, 1, 3},
			},
			ans1382: ans1382{values: []int{1, 2, 3}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1382------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1382, q.para1382
			root := buildTree(p.root)
			output := balanceBST(root)
			fmt.Printf("[input]: root=%v       [output]: balanced=%v, values=%v\n", p.root, output != nil && isBalanced(output), inOrder(output))
			if output == nil {
				t.Errorf("Expected non-nil tree")
				return
			}
			if !isBalanced(output) {
				t.Errorf("Expected balanced tree, got unbalanced")
			}
			got := inOrder(output)
			if len(got) != len(q.ans1382.values) {
				t.Errorf("Expected values %v, got %v", q.ans1382.values, got)
				return
			}
			for i := range got {
				if got[i] != q.ans1382.values[i] {
					t.Errorf("Expected values %v, got %v", q.ans1382.values, got)
					return
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}
