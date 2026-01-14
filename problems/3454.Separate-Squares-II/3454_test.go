package leetcode

import (
	"fmt"
	"math"
	"testing"
)

type question3454 struct {
	name string
	para3454
	ans3454
}

// para is parameter
// one represents the first parameter
type para3454 struct {
	squares [][]int
}

// ans is answer
// one represents the first answer
type ans3454 struct {
	one float64
}

func Test_Problem3454(t *testing.T) {

	qs := []question3454{
		{
			name: "example 1",
			para3454: para3454{
				squares: [][]int{{0, 0, 1}, {2, 2, 1}},
			},
			ans3454: ans3454{1.00000},
		},
		{
			name: "example 2",
			para3454: para3454{
				squares: [][]int{{0, 0, 2}, {1, 1, 1}},
			},
			ans3454: ans3454{1.00000},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3454------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3454, q.para3454
			output := separateSquares(p.squares)
			fmt.Printf("[input]: squares=%v       [output]:%v\n", p.squares, output)
			if math.Abs(output-q.ans3454.one) > 1e-5 {
				t.Errorf("Expected %v, got %v", q.ans3454.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
