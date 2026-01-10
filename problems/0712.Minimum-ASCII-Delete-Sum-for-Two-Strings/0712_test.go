package leetcode

import (
	"fmt"
	"testing"
)

type question0712 struct {
	name string
	para0712
	ans0712
}

// para is parameter
// one represents the first parameter
type para0712 struct {
	s1 string
	s2 string
}

// ans is answer
// one represents the first answer
type ans0712 struct {
	one int
}

func Test_Problem0712(t *testing.T) {

	qs := []question0712{
		{
			name: "example 1",
			para0712: para0712{
				s1: "sea",
				s2: "eat",
			},
			ans0712: ans0712{231},
		},
		{
			name: "example 2",
			para0712: para0712{
				s1: "delete",
				s2: "leet",
			},
			ans0712: ans0712{403},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0712------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0712, q.para0712
			output := minimumDeleteSum(p.s1, p.s2)
			fmt.Printf("[input]: s1=%v, s2=%v       [output]:%v\n", p.s1, p.s2, output)
			if output != q.ans0712.one {
				t.Errorf("Expected %v, got %v", q.ans0712.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
