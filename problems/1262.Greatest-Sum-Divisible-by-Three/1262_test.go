package leetcode

import (
	"fmt"
	"testing"
)

type question1262 struct {
	name string
	para1262
	ans1262
}

type para1262 struct {
	nums []int
}

type ans1262 struct {
	one int
}

func Test_Problem1262(t *testing.T) {

	qs := []question1262{
		{
			name: "example 1",
			para1262: para1262{
				nums: []int{3, 6, 5, 1, 8},
			},
			ans1262: ans1262{18},
		},
		{
			name: "example 2",
			para1262: para1262{
				nums: []int{4},
			},
			ans1262: ans1262{0},
		},
		{
			name: "example 3",
			para1262: para1262{
				nums: []int{1, 2, 3, 4, 4},
			},
			ans1262: ans1262{12},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1262------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1262, q.para1262
			output := maxSumDivThree(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans1262.one {
				t.Errorf("Expected %v, got %v", q.ans1262.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

