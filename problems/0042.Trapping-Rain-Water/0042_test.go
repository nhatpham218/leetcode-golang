package leetcode

import (
	"fmt"
	"testing"
)

type question0042 struct {
	name string
	para0042
	ans0042
}

// para is parameter
// one represents the first parameter
type para0042 struct {
	height []int
}

// ans is answer
// one represents the first answer
type ans0042 struct {
	one int
}

func Test_Problem0042(t *testing.T) {

	qs := []question0042{
		{
			name: "example 1",
			para0042: para0042{
				height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			},
			ans0042: ans0042{6},
		},
		{
			name: "example 2",
			para0042: para0042{
				height: []int{4, 2, 0, 3, 2, 5},
			},
			ans0042: ans0042{9},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0042------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0042, q.para0042
			output := trap(p.height)
			fmt.Printf("[input]: height=%v       [output]:%v\n", p.height, output)
			if output != q.ans0042.one {
				t.Errorf("Expected %v, got %v", q.ans0042.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
