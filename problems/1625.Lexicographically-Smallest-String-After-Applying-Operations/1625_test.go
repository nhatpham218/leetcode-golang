package leetcode

import (
	"fmt"
	"testing"
)

type question1625 struct {
	name string
	para1625
	ans1625
}

// para is parameter
// one represents the first parameter
type para1625 struct {
	s string
	a int
	b int
}

// ans is answer
// one represents the first answer
type ans1625 struct {
	one string
}

func Test_Problem1625(t *testing.T) {

	qs := []question1625{
		{
			name: "example 1",
			para1625: para1625{
				s: "5525",
				a: 9,
				b: 2,
			},
			ans1625: ans1625{"2050"},
		},
		{
			name: "example 2",
			para1625: para1625{
				s: "74",
				a: 5,
				b: 1,
			},
			ans1625: ans1625{"24"},
		},
		{
			name: "example 3",
			para1625: para1625{
				s: "0011",
				a: 4,
				b: 2,
			},
			ans1625: ans1625{"0011"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1625------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1625, q.para1625
			output := findLexSmallestString(p.s, p.a, p.b)
			fmt.Printf("[input]: s=%v a=%v b=%v       [output]:%v\n", p.s, p.a, p.b, output)
			if output != q.ans1625.one {
				t.Errorf("Expected %v, got %v", q.ans1625.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
