package leetcode

import (
	"fmt"
	"testing"
)

type question3195 struct {
	para3195
	ans3195
}

// para is parameter
// one represents the first parameter
type para3195 struct {
	grid [][]int
}

// ans is answer
// one represents the first answer
type ans3195 struct {
	one int
}

func Test_Problem3195(t *testing.T) {

	qs := []question3195{
		{
			para3195{[][]int{{0, 1, 0}, {1, 0, 1}}},
			ans3195{6},
		},
		{
			para3195{[][]int{{1, 0}, {0, 0}}},
			ans3195{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3195------------------------\n")

	for i, q := range qs {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			_, p := q.ans3195, q.para3195
			output := minimumArea(p.grid)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			if output != q.ans3195.one {
				t.Errorf("Expected %v, got %v", q.ans3195.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
