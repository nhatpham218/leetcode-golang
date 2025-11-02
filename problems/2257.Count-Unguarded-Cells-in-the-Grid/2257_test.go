package leetcode

import (
	"fmt"
	"testing"
)

type question2257 struct {
	name string
	para2257
	ans2257
}

type para2257 struct {
	m      int
	n      int
	guards [][]int
	walls  [][]int
}

type ans2257 struct {
	output int
}

func Test_Problem2257(t *testing.T) {

	qs := []question2257{
		{
			name: "Example 1",
			para2257: para2257{
				m:      4,
				n:      6,
				guards: [][]int{{0, 0}, {1, 1}, {2, 3}},
				walls:  [][]int{{0, 1}, {2, 2}, {1, 4}},
			},
			ans2257: ans2257{output: 7},
		},
		{
			name: "Example 2",
			para2257: para2257{
				m:      3,
				n:      3,
				guards: [][]int{{1, 1}},
				walls:  [][]int{{0, 1}, {1, 0}, {2, 1}, {1, 2}},
			},
			ans2257: ans2257{output: 4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2257------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans2257, q.para2257
			output := countUnguarded(p.m, p.n, p.guards, p.walls)

			fmt.Printf("[input]: m=%v, n=%v, guards=%v, walls=%v       [output]:%v\n", p.m, p.n, p.guards, p.walls, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
