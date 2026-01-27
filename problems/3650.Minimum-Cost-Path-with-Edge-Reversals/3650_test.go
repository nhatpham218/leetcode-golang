package leetcode

import (
	"fmt"
	"testing"
)

type question3650 struct {
	name string
	para3650
	ans3650
}

type para3650 struct {
	n     int
	edges [][]int
}

type ans3650 struct {
	one int
}

func Test_Problem3650(t *testing.T) {

	qs := []question3650{
		{
			name: "example 1",
			para3650: para3650{
				n:     4,
				edges: [][]int{{0, 1, 3}, {3, 1, 1}, {2, 3, 4}, {0, 2, 2}},
			},
			ans3650: ans3650{5},
		},
		{
			name: "example 2",
			para3650: para3650{
				n:     4,
				edges: [][]int{{0, 2, 1}, {2, 1, 1}, {1, 3, 1}, {2, 3, 3}},
			},
			ans3650: ans3650{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3650------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3650, q.para3650
			output := minCost(p.n, p.edges)
			fmt.Printf("[input]: n=%v, edges=%v       [output]:%v\n", p.n, p.edges, output)
			if output != q.ans3650.one {
				t.Errorf("Expected %v, got %v", q.ans3650.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
