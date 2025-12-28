package leetcode

import (
	"fmt"
	"testing"
)

type question1351 struct {
	name string
	para1351
	ans1351
}

type para1351 struct {
	grid [][]int
}

type ans1351 struct {
	one int
}

func Test_Problem1351(t *testing.T) {

	qs := []question1351{
		{
			name: "example 1",
			para1351: para1351{
				grid: [][]int{{4, 3, 2, -1}, {3, 2, 1, -1}, {1, 1, -1, -2}, {-1, -1, -2, -3}},
			},
			ans1351: ans1351{8},
		},
		{
			name: "example 2",
			para1351: para1351{
				grid: [][]int{{3, 2}, {1, 0}},
			},
			ans1351: ans1351{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1351------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1351, q.para1351
			output := countNegatives(p.grid)
			fmt.Printf("[input]: grid=%v       [output]:%v\n", p.grid, output)
			if output != q.ans1351.one {
				t.Errorf("Expected %v, got %v", q.ans1351.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
