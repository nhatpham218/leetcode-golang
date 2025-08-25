package leetcode

import (
	"fmt"
	"testing"
)

type question1277 struct {
	para1277
	ans1277
}

// para is parameter
// one represents the first parameter
type para1277 struct {
	matrix [][]int
}

// ans is answer
// one represents the first answer
type ans1277 struct {
	one int
}

func Test_Problem1277(t *testing.T) {

	qs := []question1277{
		{
			para1277{[][]int{
				{0, 1, 1, 1},
				{1, 1, 1, 1},
				{0, 1, 1, 1},
			}},
			ans1277{15},
		},

		{
			para1277{[][]int{
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 0},
			}},
			ans1277{7},
		},
		// If you need more tests, you can copy the elements above.
		{
			para1277{[][]int{
				{1, 0},
				{1, 0},
				{0, 1},
			}},
			ans1277{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1277------------------------\n")

	for i, q := range qs {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			_, p := q.ans1277, q.para1277
			output := countSquares(p.matrix)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			if output != q.ans1277.one {
				t.Errorf("Expected %d, got %d", q.ans1277.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
