package leetcode

import (
	"fmt"
	"testing"
)

type question3217 struct {
	name string
	para3217
	ans3217
}

type para3217 struct {
	nums []int
	head *ListNode
}

type ans3217 struct {
	output *ListNode
}

func Test_Problem3217(t *testing.T) {

	qs := []question3217{
		{
			name: "Example 1",
			para3217: para3217{
				nums: []int{1, 2, 3},
				head: createList([]int{1, 2, 3, 4, 5}),
			},
			ans3217: ans3217{output: createList([]int{4, 5})},
		},
		{
			name: "Example 2",
			para3217: para3217{
				nums: []int{1},
				head: createList([]int{1, 2, 1, 2, 1, 2}),
			},
			ans3217: ans3217{output: createList([]int{2, 2, 2})},
		},
		{
			name: "Example 3",
			para3217: para3217{
				nums: []int{5},
				head: createList([]int{1, 2, 3, 4}),
			},
			ans3217: ans3217{output: createList([]int{1, 2, 3, 4})},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3217------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans3217, q.para3217
			output := modifiedList(p.nums, p.head)

			fmt.Printf("[input]: nums=%v, head=%v       [output]:%v\n", p.nums, listToArray(p.head), listToArray(output))
			if !equalList(output, a.output) {
				t.Errorf("expected %v, got %v", listToArray(a.output), listToArray(output))
			}
		})
	}
	fmt.Printf("\n\n\n")
}

func createList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}
	head := &ListNode{Val: vals[0]}
	curr := head
	for i := 1; i < len(vals); i++ {
		curr.Next = &ListNode{Val: vals[i]}
		curr = curr.Next
	}
	return head
}

func listToArray(head *ListNode) []int {
	result := []int{}
	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}
	return result
}

func equalList(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}
