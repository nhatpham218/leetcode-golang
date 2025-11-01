package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

// 3217. Delete Nodes From Linked List Present in Array
// https://leetcode.com/problems/delete-nodes-from-linked-list-present-in-array/
func modifiedList(nums []int, head *ListNode) *ListNode {
	toRemove := make(map[int]bool)
	for _, n := range nums {
		toRemove[n] = true
	}

	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil {
		if toRemove[prev.Next.Val] {
			prev.Next = prev.Next.Next
		} else {
			prev = prev.Next
		}
	}

	return dummy.Next
}
