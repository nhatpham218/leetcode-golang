package leetcode

import (
	"fmt"
	"testing"
)

type question0865 struct {
	name string
	para0865
	ans0865
}

// para is parameter
// one represents the first parameter
type para0865 struct {
	root []any
}

// ans is answer
// one represents the first answer
type ans0865 struct {
	one int
}

// buildTree builds a binary tree from a level-order array representation
func buildTree(arr []any) *TreeNode {
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

func Test_Problem0865(t *testing.T) {

	qs := []question0865{
		{
			name: "example 1",
			para0865: para0865{
				root: []any{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4},
			},
			ans0865: ans0865{2},
		},
		{
			name: "example 2",
			para0865: para0865{
				root: []any{1},
			},
			ans0865: ans0865{1},
		},
		{
			name: "example 3",
			para0865: para0865{
				root: []any{0, 1, 3, nil, 2},
			},
			ans0865: ans0865{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0865------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0865, q.para0865
			tree := buildTree(p.root)
			output := subtreeWithAllDeepest(tree)
			outputVal := -1
			if output != nil {
				outputVal = output.Val
			}
			fmt.Printf("[input]: root=%v       [output]:%v\n", p.root, outputVal)
			if outputVal != q.ans0865.one {
				t.Errorf("Expected %v, got %v", q.ans0865.one, outputVal)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
