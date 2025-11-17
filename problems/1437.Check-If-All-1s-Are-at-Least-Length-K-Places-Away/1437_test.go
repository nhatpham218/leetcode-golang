package leetcode

import (
	"fmt"
	"testing"
)

type question1437 struct {
	name string
	para1437
	ans1437
}

type para1437 struct {
	nums []int
	k    int
}

type ans1437 struct {
	output bool
}

func Test_Problem1437(t *testing.T) {

	qs := []question1437{
		{
			name: "Example 1",
			para1437: para1437{
				nums: []int{1, 0, 0, 0, 1, 0, 0, 1},
				k:    2,
			},
			ans1437: ans1437{output: true},
		},
		{
			name: "Example 2",
			para1437: para1437{
				nums: []int{1, 0, 0, 1, 0, 1},
				k:    2,
			},
			ans1437: ans1437{output: false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1437------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans1437, q.para1437
			output := kLengthApart(p.nums, p.k)

			fmt.Printf("[input]: nums=%v k=%v       [output]:%v\n", p.nums, p.k, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

