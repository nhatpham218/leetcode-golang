package leetcode

import (
	"fmt"
	"reflect"
	"testing"
)

type question2536 struct {
	name string
	para2536
	ans2536
}

type para2536 struct {
	n       int
	queries [][]int
}

type ans2536 struct {
	output [][]int
}

func Test_Problem2536(t *testing.T) {

	qs := []question2536{
		{
			name: "Example 1",
			para2536: para2536{
				n:       3,
				queries: [][]int{{1, 1, 2, 2}, {0, 0, 1, 1}},
			},
			ans2536: ans2536{output: [][]int{{1, 1, 0}, {1, 2, 1}, {0, 1, 1}}},
		},
		{
			name: "Example 2",
			para2536: para2536{
				n:       2,
				queries: [][]int{{0, 0, 1, 1}},
			},
			ans2536: ans2536{output: [][]int{{1, 1}, {1, 1}}},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2536------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans2536, q.para2536
			output := rangeAddQueries(p.n, p.queries)

			fmt.Printf("[input]: n=%v queries=%v       [output]:%v\n", p.n, p.queries, output)
			if !reflect.DeepEqual(output, a.output) {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

