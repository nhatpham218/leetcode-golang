package leetcode

import (
	"fmt"
	"testing"
)

type question1015 struct {
	name string
	para1015
	ans1015
}

type para1015 struct {
	k int
}

type ans1015 struct {
	output int
}

func Test_Problem1015(t *testing.T) {

	qs := []question1015{
		{
			name: "Example 1",
			para1015: para1015{
				k: 1,
			},
			ans1015: ans1015{output: 1},
		},
		{
			name: "Example 2",
			para1015: para1015{
				k: 2,
			},
			ans1015: ans1015{output: -1},
		},
		{
			name: "Example 3",
			para1015: para1015{
				k: 3,
			},
			ans1015: ans1015{output: 3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1015------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans1015, q.para1015
			output := smallestRepunitDivByK(p.k)

			fmt.Printf("[input]: k=%v       [output]:%v\n", p.k, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

