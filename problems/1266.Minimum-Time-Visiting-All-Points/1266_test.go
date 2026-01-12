package leetcode

import (
	"fmt"
	"testing"
)

type question1266 struct {
	name string
	para1266
	ans1266
}

// para is parameter
// one represents the first parameter
type para1266 struct {
	points [][]int
}

// ans is answer
// one represents the first answer
type ans1266 struct {
	one int
}

func Test_Problem1266(t *testing.T) {

	qs := []question1266{
		{
			name: "example 1",
			para1266: para1266{
				points: [][]int{{1, 1}, {3, 4}, {-1, 0}},
			},
			ans1266: ans1266{7},
		},
		{
			name: "example 2",
			para1266: para1266{
				points: [][]int{{3, 2}, {-2, 2}},
			},
			ans1266: ans1266{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1266------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1266, q.para1266
			output := minTimeToVisitAllPoints(p.points)
			fmt.Printf("[input]: points=%v       [output]:%v\n", p.points, output)
			if output != q.ans1266.one {
				t.Errorf("Expected %v, got %v", q.ans1266.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
