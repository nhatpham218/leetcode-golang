package leetcode

import (
	"fmt"
	"testing"
)

type question474 struct {
	name string
	para474
	ans474
}

type para474 struct {
	strs []string
	m    int
	n    int
}

type ans474 struct {
	output int
}

func Test_Problem474(t *testing.T) {

	qs := []question474{
		{
			name: "Example 1",
			para474: para474{
				strs: []string{"10", "0001", "111001", "1", "0"},
				m:    5,
				n:    3,
			},
			ans474: ans474{output: 4},
		},
		{
			name: "Example 2",
			para474: para474{
				strs: []string{"10", "0", "1"},
				m:    1,
				n:    1,
			},
			ans474: ans474{output: 2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 474------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans474, q.para474
			output := findMaxForm(p.strs, p.m, p.n)

			fmt.Printf("[input]: strs=%v, m=%v, n=%v       [output]:%v\n", p.strs, p.m, p.n, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
