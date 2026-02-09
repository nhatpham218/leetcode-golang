package leetcode

import (
	"fmt"
	"testing"
)

type question0055 struct {
	name string
	para0055
	ans0055
}

// para is parameter
// one represents the first parameter
type para0055 struct {
	nums []int
}

// ans is answer
// one represents the first answer
type ans0055 struct {
	one bool
}

func Test_Problem0055(t *testing.T) {

	qs := []question0055{
		{
			name:    "example 1",
			para0055: para0055{nums: []int{2, 3, 1, 1, 4}},
			ans0055: ans0055{one: true},
		},
		{
			name:    "example 2",
			para0055: para0055{nums: []int{3, 2, 1, 0, 4}},
			ans0055: ans0055{one: false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 55------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0055, q.para0055
			output := canJump(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans0055.one {
				t.Errorf("Expected %v, got %v", q.ans0055.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
