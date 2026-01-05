package leetcode

import (
	"fmt"
	"testing"
)

type question1975 struct {
	name string
	para1975
	ans1975
}

type para1975 struct {
	matrix [][]int
}

type ans1975 struct {
	one int64
}

func Test_Problem1975(t *testing.T) {

	qs := []question1975{
		{
			name: "example 1",
			para1975: para1975{
				matrix: [][]int{{1, -1}, {-1, 1}},
			},
			ans1975: ans1975{4},
		},
		{
			name: "example 2",
			para1975: para1975{
				matrix: [][]int{{1, 2, 3}, {-1, -2, -3}, {1, 2, 3}},
			},
			ans1975: ans1975{16},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1975------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1975, q.para1975
			output := maxMatrixSum(p.matrix)
			fmt.Printf("[input]: matrix=%v       [output]:%v\n", p.matrix, output)
			if output != q.ans1975.one {
				t.Errorf("Expected %v, got %v", q.ans1975.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

