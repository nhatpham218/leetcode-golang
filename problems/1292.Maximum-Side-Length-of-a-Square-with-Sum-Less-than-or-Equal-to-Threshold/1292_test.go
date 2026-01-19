package leetcode

import (
	"fmt"
	"testing"
)

type question1292 struct {
	name string
	para1292
	ans1292
}

// para is parameter
// one represents the first parameter
type para1292 struct {
	mat       [][]int
	threshold int
}

// ans is answer
// one represents the first answer
type ans1292 struct {
	one int
}

func Test_Problem1292(t *testing.T) {

	qs := []question1292{
		{
			name: "example 1",
			para1292: para1292{
				mat:       [][]int{{1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}, {1, 1, 3, 2, 4, 3, 2}},
				threshold: 4,
			},
			ans1292: ans1292{2},
		},
		{
			name: "example 2",
			para1292: para1292{
				mat:       [][]int{{2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}, {2, 2, 2, 2, 2}},
				threshold: 1,
			},
			ans1292: ans1292{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1292------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1292, q.para1292
			output := maxSideLength(p.mat, p.threshold)
			fmt.Printf("[input]: mat=%v, threshold=%v       [output]:%v\n", p.mat, p.threshold, output)
			if output != q.ans1292.one {
				t.Errorf("Expected %v, got %v", q.ans1292.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
