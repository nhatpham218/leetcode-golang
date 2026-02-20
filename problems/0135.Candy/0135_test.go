package leetcode

import (
	"fmt"
	"testing"
)

type question0135 struct {
	name string
	para0135
	ans0135
}

// para is parameter
// one represents the first parameter
type para0135 struct {
	ratings []int
}

// ans is answer
// one represents the first answer
type ans0135 struct {
	one int
}

func Test_Problem0135(t *testing.T) {

	qs := []question0135{
		{
			name: "example 1",
			para0135: para0135{
				ratings: []int{1, 0, 2},
			},
			ans0135: ans0135{5},
		},
		{
			name: "example 2",
			para0135: para0135{
				ratings: []int{1, 2, 2},
			},
			ans0135: ans0135{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0135------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0135, q.para0135
			output := candy(p.ratings)
			fmt.Printf("[input]: ratings=%v       [output]:%v\n", p.ratings, output)
			if output != q.ans0135.one {
				t.Errorf("Expected %v, got %v", q.ans0135.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
