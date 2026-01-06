package leetcode

import (
	"fmt"
	"testing"
)

type question1161 struct {
	name string
	para1161
	ans1161
}

type para1161 struct {
	root []interface{}
}

type ans1161 struct {
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

func Test_Problem1161(t *testing.T) {

	qs := []question1161{
		{
			name: "example 1",
			para1161: para1161{
				root: []interface{}{1, 7, 0, 7, -8, nil, nil},
			},
			ans1161: ans1161{2},
		},
		{
			name: "example 2",
			para1161: para1161{
				root: []interface{}{989, nil, 10250, 98693, -89388, nil, nil, nil, -32127},
			},
			ans1161: ans1161{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1161------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1161, q.para1161
			root := buildTree(p.root)
			output := maxLevelSum(root)
			fmt.Printf("[input]: root=%v       [output]:%v\n", p.root, output)
			if output != q.ans1161.one {
				t.Errorf("Expected %v, got %v", q.ans1161.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

