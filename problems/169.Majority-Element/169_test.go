package leetcode

import (
	"fmt"
	"testing"
)

type question169 struct {
	name string
	para169
	ans169
}

type para169 struct {
	nums []int
}

type ans169 struct {
	one int
}

func Test_Problem169(t *testing.T) {

	qs := []question169{
		{
			name:    "example 1",
			para169: para169{nums: []int{3, 2, 3}},
			ans169:  ans169{3},
		},
		{
			name:    "example 2",
			para169: para169{nums: []int{2, 2, 1, 1, 1, 2, 2}},
			ans169:  ans169{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 169------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans169, q.para169
			output := majorityElement(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans169.one {
				t.Errorf("Expected %v, got %v", q.ans169.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
