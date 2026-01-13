package leetcode

import (
	"fmt"
	"math"
	"testing"
)

type question3453 struct {
	name string
	para3453
	ans3453
}

// para is parameter
// one represents the first parameter
type para3453 struct {
	squares [][]int
}

// ans is answer
// one represents the first answer
type ans3453 struct {
	one float64
}

func Test_Problem3453(t *testing.T) {

	qs := []question3453{
		{
			name: "example 1",
			para3453: para3453{
				squares: [][]int{{0, 0, 1}, {2, 2, 1}},
			},
			ans3453: ans3453{1.00000},
		},
		{
			name: "example 2",
			para3453: para3453{
				squares: [][]int{{0, 0, 2}, {1, 1, 1}},
			},
			ans3453: ans3453{1.16667},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3453------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3453, q.para3453
			output := separateSquares(p.squares)
			fmt.Printf("[input]: squares=%v       [output]:%v\n", p.squares, output)
			if math.Abs(output-q.ans3453.one) > 1e-5 {
				t.Errorf("Expected %v, got %v", q.ans3453.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
