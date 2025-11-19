package leetcode

import (
	"fmt"
	"testing"
)

type question2154 struct {
	name string
	para2154
	ans2154
}

type para2154 struct {
	nums     []int
	original int
}

type ans2154 struct {
	output int
}

func Test_Problem2154(t *testing.T) {

	qs := []question2154{
		{
			name: "Example 1",
			para2154: para2154{
				nums:     []int{5, 3, 6, 1, 12},
				original: 3,
			},
			ans2154: ans2154{output: 24},
		},
		{
			name: "Example 2",
			para2154: para2154{
				nums:     []int{2, 7, 9},
				original: 4,
			},
			ans2154: ans2154{output: 4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2154------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans2154, q.para2154
			output := findFinalValue(p.nums, p.original)

			fmt.Printf("[input]: nums=%v, original=%v       [output]:%v\n", p.nums, p.original, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}



