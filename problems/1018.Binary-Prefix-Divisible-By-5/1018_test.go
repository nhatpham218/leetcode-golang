package leetcode

import (
	"fmt"
	"testing"
)

type question1018 struct {
	name string
	para1018
	ans1018
}

type para1018 struct {
	nums []int
}

type ans1018 struct {
	output []bool
}

func Test_Problem1018(t *testing.T) {

	qs := []question1018{
		{
			name: "Example 1",
			para1018: para1018{
				nums: []int{0, 1, 1},
			},
			ans1018: ans1018{output: []bool{true, false, false}},
		},
		{
			name: "Example 2",
			para1018: para1018{
				nums: []int{1, 1, 1},
			},
			ans1018: ans1018{output: []bool{false, false, false}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1018------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans1018, q.para1018
			output := prefixesDivBy5(p.nums)

			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if len(output) != len(a.output) {
				t.Errorf("expected %v, got %v", a.output, output)
				return
			}
			for i := range output {
				if output[i] != a.output[i] {
					t.Errorf("expected %v, got %v", a.output, output)
					return
				}
			}
		})
	}
	fmt.Printf("\n\n\n")
}
