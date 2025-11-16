package leetcode

import (
	"fmt"
	"testing"
)

type question1513 struct {
	name string
	para1513
	ans1513
}

type para1513 struct {
	s string
}

type ans1513 struct {
	output int
}

func Test_Problem1513(t *testing.T) {

	qs := []question1513{
		{
			name: "Example 1",
			para1513: para1513{
				s: "0110111",
			},
			ans1513: ans1513{output: 9},
		},
		{
			name: "Example 2",
			para1513: para1513{
				s: "101",
			},
			ans1513: ans1513{output: 2},
		},
		{
			name: "Example 3",
			para1513: para1513{
				s: "111111",
			},
			ans1513: ans1513{output: 21},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1513------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans1513, q.para1513
			output := numSub(p.s)

			fmt.Printf("[input]: s=%v       [output]:%v\n", p.s, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

