package leetcode

import (
	"fmt"
	"testing"
)

type question1590 struct {
	name string
	para1590
	ans1590
}

type para1590 struct {
	nums []int
	p    int
}

type ans1590 struct {
	output int
}

func Test_Problem1590(t *testing.T) {

	qs := []question1590{
		{
			name: "Example 1",
			para1590: para1590{
				nums: []int{3, 1, 4, 2},
				p:    6,
			},
			ans1590: ans1590{output: 1},
		},
		{
			name: "Example 2",
			para1590: para1590{
				nums: []int{6, 3, 5, 2},
				p:    9,
			},
			ans1590: ans1590{output: 2},
		},
		{
			name: "Example 3",
			para1590: para1590{
				nums: []int{1, 2, 3},
				p:    3,
			},
			ans1590: ans1590{output: 0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1590------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans1590, q.para1590
			output := minSubarray(p.nums, p.p)

			fmt.Printf("[input]: nums=%v p=%v       [output]:%v\n", p.nums, p.p, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

