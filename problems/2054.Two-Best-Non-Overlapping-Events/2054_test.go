package leetcode

import (
	"fmt"
	"testing"
)

type question2054 struct {
	name string
	para2054
	ans2054
}

type para2054 struct {
	events [][]int
}

type ans2054 struct {
	one int
}

func Test_Problem2054(t *testing.T) {

	qs := []question2054{
		{
			name: "example 1",
			para2054: para2054{
				events: [][]int{{1, 3, 2}, {4, 5, 2}, {2, 4, 3}},
			},
			ans2054: ans2054{4},
		},
		{
			name: "example 2",
			para2054: para2054{
				events: [][]int{{1, 3, 2}, {4, 5, 2}, {1, 5, 5}},
			},
			ans2054: ans2054{5},
		},
		{
			name: "example 3",
			para2054: para2054{
				events: [][]int{{1, 5, 3}, {1, 5, 1}, {6, 6, 5}},
			},
			ans2054: ans2054{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2054------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2054, q.para2054
			output := maxTwoEvents(p.events)
			fmt.Printf("[input]: events=%v       [output]:%v\n", p.events, output)
			if output != q.ans2054.one {
				t.Errorf("Expected %v, got %v", q.ans2054.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

