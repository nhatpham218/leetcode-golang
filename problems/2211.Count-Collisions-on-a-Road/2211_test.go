package leetcode

import (
	"fmt"
	"testing"
)

type question2211 struct {
	name string
	para2211
	ans2211
}

// para is parameter
// one represents the first parameter
type para2211 struct {
	one string
}

// ans is answer
// one represents the first answer
type ans2211 struct {
	one int
}

func Test_Problem2211(t *testing.T) {

	qs := []question2211{
		{
			name: "example 1",
			para2211: para2211{
				one: "RLRSLL",
			},
			ans2211: ans2211{5},
		},
		{
			name: "example 2",
			para2211: para2211{
				one: "LLRR",
			},
			ans2211: ans2211{0},
		},
		{
			name: "example 3",
			para2211: para2211{
				one: "SSRSSRLLRSLLRSRSSRLRRRRLLRRLSSRR",
			},
			ans2211: ans2211{20},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2211------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2211, q.para2211
			output := countCollisions(p.one)
			fmt.Printf("[input]: directions=%v       [output]:%v\n", p.one, output)
			if output != q.ans2211.one {
				t.Errorf("Expected %v, got %v", q.ans2211.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
