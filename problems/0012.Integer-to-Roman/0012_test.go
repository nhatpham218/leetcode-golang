package leetcode

import (
	"fmt"
	"testing"
)

type question0012 struct {
	name string
	para0012
	ans0012
}

type para0012 struct {
	one int
}

type ans0012 struct {
	one string
}

func Test_Problem0012(t *testing.T) {
	qs := []question0012{
		{
			name:     "example 1",
			para0012: para0012{one: 3},
			ans0012:  ans0012{one: "III"},
		},
		{
			name:     "example 2",
			para0012: para0012{one: 4},
			ans0012:  ans0012{one: "IV"},
		},
		{
			name:     "example 3",
			para0012: para0012{one: 9},
			ans0012:  ans0012{one: "IX"},
		},
		{
			name:     "example 4",
			para0012: para0012{one: 58},
			ans0012:  ans0012{one: "LVIII"},
		},
		{
			name:     "example 5",
			para0012: para0012{one: 1994},
			ans0012:  ans0012{one: "MCMXCIV"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0012------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0012, q.para0012
			output := intToRoman(p.one)
			fmt.Printf("[input]: num=%v       [output]:%v\n", p.one, output)
			if output != q.ans0012.one {
				t.Errorf("Expected %v, got %v", q.ans0012.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
