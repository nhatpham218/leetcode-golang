package leetcode

import (
	"fmt"
	"testing"
)

type question3651 struct {
	name string
	para3651
	ans3651
}

// para is parameter
// one represents the first parameter
type para3651 struct {
	grid [][]int
	k    int
}

// ans is answer
// one represents the first answer
type ans3651 struct {
	one int
}

func Test_Problem3651(t *testing.T) {

	qs := []question3651{
		{
			name: "example 1",
			para3651: para3651{
				grid: [][]int{{1, 3, 3}, {2, 5, 4}, {4, 3, 5}},
				k:    2,
			},
			ans3651: ans3651{7},
		},
		{
			name: "example 2",
			para3651: para3651{
				grid: [][]int{{1, 2}, {2, 3}, {3, 4}},
				k:    1,
			},
			ans3651: ans3651{9},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3651------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3651, q.para3651
			output := minCost(p.grid, p.k)
			fmt.Printf("[input]: grid=%v, k=%v       [output]:%v\n", p.grid, p.k, output)
			if output != q.ans3651.one {
				t.Errorf("Expected %v, got %v", q.ans3651.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
