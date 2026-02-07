package leetcode

import (
	"fmt"
	"testing"
)

type question1653 struct {
	name string
	para1653
	ans1653
}

type para1653 struct {
	s string
}

type ans1653 struct {
	one int
}

func Test_Problem1653(t *testing.T) {
	qs := []question1653{
		{
			name:    "example 1",
			para1653: para1653{s: "aababbab"},
			ans1653: ans1653{one: 2},
		},
		{
			name:    "example 2",
			para1653: para1653{s: "bbaaaaabb"},
			ans1653: ans1653{one: 2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1653------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1653, q.para1653
			output := minimumDeletions(p.s)
			fmt.Printf("[input]: s=%v       [output]:%v\n", p.s, output)
			if output != q.ans1653.one {
				t.Errorf("Expected %v, got %v", q.ans1653.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
