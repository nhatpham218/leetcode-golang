package leetcode

import (
	"fmt"
	"testing"
)

type question3623 struct {
	name string
	para3623
	ans3623
}

type para3623 struct {
	points [][]int
}

type ans3623 struct {
	one int
}

func Test_Problem3623(t *testing.T) {

	qs := []question3623{
		{
			name: "example 1",
			para3623: para3623{
				points: [][]int{{1, 0}, {2, 0}, {3, 0}, {2, 2}, {3, 2}},
			},
			ans3623: ans3623{3},
		},
		{
			name: "example 2",
			para3623: para3623{
				points: [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}},
			},
			ans3623: ans3623{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3623------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3623, q.para3623
			output := countTrapezoids(p.points)
			fmt.Printf("[input]: points=%v       [output]:%v\n", p.points, output)
			if output != q.ans3623.one {
				t.Errorf("Expected %v, got %v", q.ans3623.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

