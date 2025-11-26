package leetcode

import (
	"fmt"
	"testing"
)

type question2435 struct {
	name string
	para2435
	ans2435
}

type para2435 struct {
	grid [][]int
	k    int
}

type ans2435 struct {
	output int
}

func Test_Problem2435(t *testing.T) {

	qs := []question2435{
		{
			name: "Example 1",
			para2435: para2435{
				grid: [][]int{{5, 2, 4}, {3, 0, 5}, {0, 7, 2}},
				k:    3,
			},
			ans2435: ans2435{output: 2},
		},
		{
			name: "Example 2",
			para2435: para2435{
				grid: [][]int{{0, 0}},
				k:    5,
			},
			ans2435: ans2435{output: 1},
		},
		{
			name: "Example 3",
			para2435: para2435{
				grid: [][]int{{7, 3, 4, 9}, {2, 3, 6, 2}, {2, 3, 7, 0}},
				k:    1,
			},
			ans2435: ans2435{output: 10},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2435------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans2435, q.para2435
			output := numberOfPaths(p.grid, p.k)

			fmt.Printf("[input]: grid=%v k=%v       [output]:%v\n", p.grid, p.k, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

