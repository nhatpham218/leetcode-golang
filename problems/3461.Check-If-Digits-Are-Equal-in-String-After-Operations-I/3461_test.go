package leetcode

import (
	"fmt"
	"testing"
)

type question3461 struct {
	name string
	para3461
	ans3461
}

type para3461 struct {
	s string
}

type ans3461 struct {
	one bool
}

func Test_Problem3461(t *testing.T) {

	qs := []question3461{
		{
			name: "Example 1",
			para3461: para3461{
				s: "3902",
			},
			ans3461: ans3461{
				one: true,
			},
		},
		{
			name: "Example 2",
			para3461: para3461{
				s: "34789",
			},
			ans3461: ans3461{
				one: false,
			},
		},
		{
			name: "Example 3",
			para3461: para3461{
				s: "323",
			},
			ans3461: ans3461{
				one: true,
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3461------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans3461, q.para3461
			output := hasSameDigits(p.s)
			fmt.Printf("[input]: s=%v       [output]:%v\n", p.s, output)
			if output != a.one {
				t.Errorf("expected %v, got %v", a.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
