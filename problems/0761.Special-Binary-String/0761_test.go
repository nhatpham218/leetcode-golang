package leetcode

import (
	"fmt"
	"testing"
)

type question0761 struct {
	name string
	para0761
	ans0761
}

type para0761 struct {
	s string
}

type ans0761 struct {
	one string
}

func Test_Problem0761(t *testing.T) {
	qs := []question0761{
		{
			name:     "example 1",
			para0761: para0761{s: "11011000"},
			ans0761:  ans0761{one: "11100100"},
		},
		{
			name:     "example 2",
			para0761: para0761{s: "10"},
			ans0761:  ans0761{one: "10"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0761------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0761, q.para0761
			output := makeLargestSpecial(p.s)
			fmt.Printf("[input]: s=%v       [output]:%v\n", p.s, output)
			if output != q.ans0761.one {
				t.Errorf("Expected %v, got %v", q.ans0761.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
