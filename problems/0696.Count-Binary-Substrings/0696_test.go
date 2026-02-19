package leetcode

import (
	"fmt"
	"testing"
)

type question0696 struct {
	name string
	para0696
	ans0696
}

type para0696 struct {
	s string
}

type ans0696 struct {
	one int
}

func Test_Problem0696(t *testing.T) {
	qs := []question0696{
		{
			name:     "example 1",
			para0696: para0696{s: "00110011"},
			ans0696:  ans0696{one: 6},
		},
		{
			name:     "example 2",
			para0696: para0696{s: "10101"},
			ans0696:  ans0696{one: 4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0696------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0696, q.para0696
			output := countBinarySubstrings(p.s)
			fmt.Printf("[input]: s=%v       [output]:%v\n", p.s, output)
			if output != q.ans0696.one {
				t.Errorf("Expected %v, got %v", q.ans0696.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
