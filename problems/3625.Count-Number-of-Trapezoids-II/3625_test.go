package leetcode

import (
	"fmt"
	"testing"
)

type question3625 struct {
	name string
	para3625
	ans3625
}

type para3625 struct {
	points [][]int
}

type ans3625 struct {
	one int
}

func Test_Problem3625(t *testing.T) {

	qs := []question3625{
		{
			name: "example 1",
			para3625: para3625{
				points: [][]int{{-3, 2}, {3, 0}, {2, 3}, {3, 2}, {2, -3}},
			},
			ans3625: ans3625{2},
		},
		{
			name: "example 2",
			para3625: para3625{
				points: [][]int{{0, 0}, {1, 0}, {0, 1}, {2, 1}},
			},
			ans3625: ans3625{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3625------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3625, q.para3625
			output := countTrapezoids(p.points)
			fmt.Printf("[input]: points=%v       [output]:%v\n", p.points, output)
			if output != q.ans3625.one {
				t.Errorf("Expected %v, got %v", q.ans3625.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}


