package leetcode

import (
	"fmt"
	"testing"
)

type question2402 struct {
	name string
	para2402
	ans2402
}

type para2402 struct {
	n        int
	meetings [][]int
}

type ans2402 struct {
	one int
}

func Test_Problem2402(t *testing.T) {

	qs := []question2402{
		{
			name: "example 1",
			para2402: para2402{
				n:        2,
				meetings: [][]int{{0, 10}, {1, 5}, {2, 7}, {3, 4}},
			},
			ans2402: ans2402{0},
		},
		{
			name: "example 2",
			para2402: para2402{
				n:        3,
				meetings: [][]int{{1, 20}, {2, 10}, {3, 5}, {4, 9}, {6, 8}},
			},
			ans2402: ans2402{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2402------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2402, q.para2402
			output := mostBooked(p.n, p.meetings)
			fmt.Printf("[input]: n=%v, meetings=%v       [output]:%v\n", p.n, p.meetings, output)
			if output != q.ans2402.one {
				t.Errorf("Expected %v, got %v", q.ans2402.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

