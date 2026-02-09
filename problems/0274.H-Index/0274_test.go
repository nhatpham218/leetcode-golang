package leetcode

import (
	"fmt"
	"testing"
)

type question0274 struct {
	name  string
	para0274
	ans0274
}

// para is parameter
// one represents the first parameter
type para0274 struct {
	citations []int
}

// ans is answer
// one represents the first answer
type ans0274 struct {
	one int
}

func Test_Problem0274(t *testing.T) {

	qs := []question0274{
		{
			name:      "example 1",
			para0274:  para0274{citations: []int{3, 0, 6, 1, 5}},
			ans0274:   ans0274{one: 3},
		},
		{
			name:      "example 2",
			para0274:  para0274{citations: []int{1, 3, 1}},
			ans0274:   ans0274{one: 1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 274------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0274, q.para0274
			output := hIndex(p.citations)
			fmt.Printf("[input]: citations=%v       [output]:%v\n", p.citations, output)
			if output != q.ans0274.one {
				t.Errorf("Expected %v, got %v", q.ans0274.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
