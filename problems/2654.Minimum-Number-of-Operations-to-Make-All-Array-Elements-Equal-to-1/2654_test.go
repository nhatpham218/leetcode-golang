package leetcode

import (
	"fmt"
	"testing"
)

type question2654 struct {
	name string
	para2654
	ans2654
}

type para2654 struct {
	nums []int
}

type ans2654 struct {
	output int
}

func Test_Problem2654(t *testing.T) {

	qs := []question2654{
		{
			name: "Example 1",
			para2654: para2654{
				nums: []int{2, 6, 3, 4},
			},
			ans2654: ans2654{output: 4},
		},
		{
			name: "Example 2",
			para2654: para2654{
				nums: []int{2, 10, 6, 14},
			},
			ans2654: ans2654{output: -1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2654------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans2654, q.para2654
			output := minOperations(p.nums)

			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
