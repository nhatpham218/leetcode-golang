package leetcode

import (
	"fmt"
	"testing"
)

type question1895 struct {
	name string
	para1895
	ans1895
}

// para is parameter
// one represents the first parameter
type para1895 struct {
	grid [][]int
}

// ans is answer
// one represents the first answer
type ans1895 struct {
	one int
}

func Test_Problem1895(t *testing.T) {

	qs := []question1895{
		{
			name: "example 1",
			para1895: para1895{
				grid: [][]int{{7, 1, 4, 5, 6}, {2, 5, 1, 6, 4}, {1, 5, 4, 3, 2}, {1, 2, 7, 3, 4}},
			},
			ans1895: ans1895{3},
		},
		{
			name: "example 2",
			para1895: para1895{
				grid: [][]int{{5, 1, 3, 1}, {9, 3, 3, 1}, {1, 3, 3, 8}},
			},
			ans1895: ans1895{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1895------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1895, q.para1895
			output := largestMagicSquare(p.grid)
			fmt.Printf("[input]: grid=%v       [output]:%v\n", p.grid, output)
			if output != q.ans1895.one {
				t.Errorf("Expected %v, got %v", q.ans1895.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
