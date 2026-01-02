package leetcode

import (
	"fmt"
	"testing"
)

type question961 struct {
	name string
	para961
	ans961
}

type para961 struct {
	nums []int
}

type ans961 struct {
	one int
}

func Test_Problem961(t *testing.T) {

	qs := []question961{
		{
			name:    "example 1",
			para961: para961{nums: []int{1, 2, 3, 3}},
			ans961:  ans961{3},
		},
		{
			name:    "example 2",
			para961: para961{nums: []int{2, 1, 2, 5, 3, 2}},
			ans961:  ans961{2},
		},
		{
			name:    "example 3",
			para961: para961{nums: []int{5, 1, 5, 2, 5, 3, 5, 4}},
			ans961:  ans961{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 961------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans961, q.para961
			output := repeatedNTimes(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans961.one {
				t.Errorf("Expected %v, got %v", q.ans961.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

