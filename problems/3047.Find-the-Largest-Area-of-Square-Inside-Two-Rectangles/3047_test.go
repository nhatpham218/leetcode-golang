package leetcode

import (
	"fmt"
	"testing"
)

type question3047 struct {
	name string
	para3047
	ans3047
}

// para is parameter
// one represents the first parameter
type para3047 struct {
	bottomLeft [][]int
	topRight   [][]int
}

// ans is answer
// one represents the first answer
type ans3047 struct {
	one int64
}

func Test_Problem3047(t *testing.T) {

	qs := []question3047{
		{
			name: "example 1",
			para3047: para3047{
				bottomLeft: [][]int{{1, 1}, {2, 2}, {3, 1}},
				topRight:   [][]int{{3, 3}, {4, 4}, {6, 6}},
			},
			ans3047: ans3047{1},
		},
		{
			name: "example 2",
			para3047: para3047{
				bottomLeft: [][]int{{1, 1}, {2, 2}, {1, 2}},
				topRight:   [][]int{{3, 3}, {4, 4}, {3, 4}},
			},
			ans3047: ans3047{1},
		},
		{
			name: "example 3",
			para3047: para3047{
				bottomLeft: [][]int{{1, 1}, {3, 3}, {3, 1}},
				topRight:   [][]int{{2, 2}, {4, 4}, {4, 2}},
			},
			ans3047: ans3047{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3047------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3047, q.para3047
			output := largestSquareArea(p.bottomLeft, p.topRight)
			fmt.Printf("[input]: bottomLeft=%v, topRight=%v       [output]:%v\n", p.bottomLeft, p.topRight, output)
			if output != q.ans3047.one {
				t.Errorf("Expected %v, got %v", q.ans3047.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
