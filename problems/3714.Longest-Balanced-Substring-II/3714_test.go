package leetcode

import (
	"fmt"
	"testing"
)

type question3714 struct {
	name string
	para3714
	ans3714
}

// para is parameter
// one represents the first parameter
type para3714 struct {
	s string
}

// ans is answer
// one represents the first answer
type ans3714 struct {
	one int
}

func Test_Problem3714(t *testing.T) {

	qs := []question3714{
		{
			name:     "example 1",
			para3714: para3714{s: "abbac"},
			ans3714:  ans3714{one: 4},
		},
		{
			name:     "example 2",
			para3714: para3714{s: "aabcc"},
			ans3714:  ans3714{one: 3},
		},
		{
			name:     "example 3",
			para3714: para3714{s: "aba"},
			ans3714:  ans3714{one: 2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3714------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3714, q.para3714
			output := longestBalanced(p.s)
			fmt.Printf("[input]: s=%v       [output]:%v\n", p.s, output)
			if output != q.ans3714.one {
				t.Errorf("Expected %v, got %v", q.ans3714.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
