package leetcode

import (
	"fmt"
	"testing"
)

type question3607 struct {
	name string
	para3607
	ans3607
}

type para3607 struct {
	c           int
	connections [][]int
	queries     [][]int
}

type ans3607 struct {
	output []int
}

func Test_Problem3607(t *testing.T) {

	qs := []question3607{
		{
			name: "Example 1",
			para3607: para3607{
				c:           5,
				connections: [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}},
				queries:     [][]int{{1, 3}, {2, 1}, {1, 1}, {2, 2}, {1, 2}},
			},
			ans3607: ans3607{output: []int{3, 2, 3}},
		},
		{
			name: "Example 2",
			para3607: para3607{
				c:           3,
				connections: [][]int{},
				queries:     [][]int{{1, 1}, {2, 1}, {1, 1}},
			},
			ans3607: ans3607{output: []int{1, -1}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3607------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans3607, q.para3607
			output := processQueries(p.c, p.connections, p.queries)

			fmt.Printf("[input]: c=%v, connections=%v, queries=%v       [output]:%v\n", p.c, p.connections, p.queries, output)
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
