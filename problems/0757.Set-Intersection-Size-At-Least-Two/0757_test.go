package leetcode

import (
	"fmt"
	"testing"
)

type question757 struct {
	name string
	para757
	ans757
}

type para757 struct {
	intervals [][]int
}

type ans757 struct {
	output int
}

func Test_Problem757(t *testing.T) {

	qs := []question757{
		{
			name: "Example 1",
			para757: para757{
				intervals: [][]int{{1, 3}, {3, 7}, {8, 9}},
			},
			ans757: ans757{output: 5},
		},
		{
			name: "Example 2",
			para757: para757{
				intervals: [][]int{{1, 3}, {1, 4}, {2, 5}, {3, 5}},
			},
			ans757: ans757{output: 3},
		},
		{
			name: "Example 3",
			para757: para757{
				intervals: [][]int{{1, 2}, {2, 3}, {2, 4}, {4, 5}},
			},
			ans757: ans757{output: 5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 757------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans757, q.para757
			output := intersectionSizeTwo(p.intervals)

			fmt.Printf("[input]: intervals=%v       [output]:%v\n", p.intervals, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}




