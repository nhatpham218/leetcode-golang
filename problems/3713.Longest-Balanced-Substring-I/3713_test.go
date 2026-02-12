package leetcode

import (
	"fmt"
	"testing"
)

type question3713 struct {
	name string
	para3713
	ans3713
}

// para is parameter
// one represents the first parameter
type para3713 struct {
	s string
}

// ans is answer
// one represents the first answer
type ans3713 struct {
	one int
}

func Test_Problem3713(t *testing.T) {

	qs := []question3713{
		{
			name:     "example 1",
			para3713: para3713{s: "abbac"},
			ans3713:  ans3713{4},
		},
		{
			name:     "example 2",
			para3713: para3713{s: "zzabccy"},
			ans3713:  ans3713{4},
		},
		{
			name:     "example 3",
			para3713: para3713{s: "aba"},
			ans3713:  ans3713{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3713------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3713, q.para3713
			output := longestBalanced(p.s)
			fmt.Printf("[input]: s=%v       [output]:%v\n", p.s, output)
			if output != q.ans3713.one {
				t.Errorf("Expected %v, got %v", q.ans3713.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
