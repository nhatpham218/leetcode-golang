package leetcode

import (
	"fmt"
	"testing"
)

type question2141 struct {
	name string
	para2141
	ans2141
}

type para2141 struct {
	n         int
	batteries []int
}

type ans2141 struct {
	one int64
}

func Test_Problem2141(t *testing.T) {

	qs := []question2141{
		{
			name: "example 1",
			para2141: para2141{
				n:         2,
				batteries: []int{3, 3, 3},
			},
			ans2141: ans2141{4},
		},
		{
			name: "example 2",
			para2141: para2141{
				n:         2,
				batteries: []int{1, 1, 1, 1},
			},
			ans2141: ans2141{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2141------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2141, q.para2141
			output := maxRunTime(p.n, p.batteries)
			fmt.Printf("[input]: n=%v, batteries=%v       [output]:%v\n", p.n, p.batteries, output)
			if output != q.ans2141.one {
				t.Errorf("Expected %v, got %v", q.ans2141.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
