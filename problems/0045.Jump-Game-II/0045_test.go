package leetcode

import (
	"fmt"
	"testing"
)

type question0045 struct {
	name string
	para0045
	ans0045
}

// para is parameter
// one represents the first parameter
type para0045 struct {
	nums []int
}

// ans is answer
// one represents the first answer
type ans0045 struct {
	one int
}

func Test_Problem0045(t *testing.T) {

	qs := []question0045{
		{
			name:     "example 1",
			para0045: para0045{nums: []int{2, 3, 1, 1, 4}},
			ans0045:  ans0045{one: 2},
		},
		{
			name:     "example 2",
			para0045: para0045{nums: []int{2, 3, 0, 1, 4}},
			ans0045:  ans0045{one: 2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 45------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0045, q.para0045
			output := jump(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans0045.one {
				t.Errorf("Expected %v, got %v", q.ans0045.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
