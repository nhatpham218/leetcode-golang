package leetcode

import (
	"fmt"
	"testing"
)

type question0013 struct {
	name string
	para0013
	ans0013
}

type para0013 struct {
	one string
}

type ans0013 struct {
	one int
}

func Test_Problem0013(t *testing.T) {
	qs := []question0013{
		{
			name:     "example 1",
			para0013: para0013{one: "III"},
			ans0013:  ans0013{one: 3},
		},
		{
			name:     "example 2",
			para0013: para0013{one: "LVIII"},
			ans0013:  ans0013{one: 58},
		},
		{
			name:     "example 3",
			para0013: para0013{one: "MCMXCIV"},
			ans0013:  ans0013{one: 1994},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0013------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0013, q.para0013
			output := romanToInt(p.one)
			fmt.Printf("[input]: s=%v       [output]:%v\n", p.one, output)
			if output != q.ans0013.one {
				t.Errorf("Expected %v, got %v", q.ans0013.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
