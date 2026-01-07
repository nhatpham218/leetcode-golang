package leetcode

import (
	"fmt"
	"testing"
)

type question1339 struct {
	name string
	para1339
	ans1339
}

type para1339 struct {
	root []interface{}
}

type ans1339 struct {
	one int
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

func Test_Problem1339(t *testing.T) {

	qs := []question1339{
		{
			name: "example 1",
			para1339: para1339{
				root: []interface{}{1, 2, 3, 4, 5, 6},
			},
			ans1339: ans1339{110},
		},
		{
			name: "example 2",
			para1339: para1339{
				root: []interface{}{1, nil, 2, 3, 4, nil, nil, 5, 6},
			},
			ans1339: ans1339{90},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1339------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1339, q.para1339
			root := buildTree(p.root)
			output := maxProduct(root)
			fmt.Printf("[input]: root=%v       [output]:%v\n", p.root, output)
			if output != q.ans1339.one {
				t.Errorf("Expected %v, got %v", q.ans1339.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

