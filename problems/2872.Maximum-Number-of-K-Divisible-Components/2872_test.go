package leetcode

import (
	"fmt"
	"testing"
)

type question2872 struct {
	name string
	para2872
	ans2872
}

type para2872 struct {
	n      int
	edges  [][]int
	values []int
	k      int
}

type ans2872 struct {
	output int
}

func Test_Problem2872(t *testing.T) {

	qs := []question2872{
		{
			name: "Example 1",
			para2872: para2872{
				n:      5,
				edges:  [][]int{{0, 2}, {1, 2}, {1, 3}, {2, 4}},
				values: []int{1, 8, 1, 4, 4},
				k:      6,
			},
			ans2872: ans2872{output: 2},
		},
		{
			name: "Example 2",
			para2872: para2872{
				n:      7,
				edges:  [][]int{{0, 1}, {0, 2}, {1, 3}, {1, 4}, {2, 5}, {2, 6}},
				values: []int{3, 0, 6, 1, 5, 2, 1},
				k:      3,
			},
			ans2872: ans2872{output: 3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2872------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans2872, q.para2872
			output := maxKDivisibleComponents(p.n, p.edges, p.values, p.k)

			fmt.Printf("[input]: n=%v edges=%v values=%v k=%v       [output]:%v\n", p.n, p.edges, p.values, p.k, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
