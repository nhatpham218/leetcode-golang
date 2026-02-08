package leetcode

import (
	"fmt"
	"testing"
)

type question110 struct {
	name string
	para110
	ans110
}

type para110 struct {
	root []interface{}
}

type ans110 struct {
	one bool
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

func Test_Problem110(t *testing.T) {

	qs := []question110{
		{
			name: "example 1",
			para110: para110{
				root: []interface{}{3, 9, 20, nil, nil, 15, 7},
			},
			ans110: ans110{true},
		},
		{
			name: "example 2",
			para110: para110{
				root: []interface{}{1, 2, 2, 3, 3, nil, nil, 4, 4},
			},
			ans110: ans110{false},
		},
		{
			name: "example 3",
			para110: para110{
				root: []interface{}{},
			},
			ans110: ans110{true},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 110------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans110, q.para110
			root := buildTree(p.root)
			output := isBalanced(root)
			fmt.Printf("[input]: root=%v       [output]:%v\n", p.root, output)
			if output != q.ans110.one {
				t.Errorf("Expected %v, got %v", q.ans110.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
