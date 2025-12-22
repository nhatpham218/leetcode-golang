package leetcode

import (
	"fmt"
	"testing"
)

type question960 struct {
	name string
	para960
	ans960
}

// para is parameter
// one represents the first parameter
type para960 struct {
	strs []string
}

// ans is answer
// one represents the first answer
type ans960 struct {
	one int
}

func Test_Problem960(t *testing.T) {

	qs := []question960{
		{
			name: "example 1",
			para960: para960{
				strs: []string{"babca", "bbazb"},
			},
			ans960: ans960{3},
		},
		{
			name: "example 2",
			para960: para960{
				strs: []string{"edcba"},
			},
			ans960: ans960{4},
		},
		{
			name: "example 3",
			para960: para960{
				strs: []string{"ghi", "def", "abc"},
			},
			ans960: ans960{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 960------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans960, q.para960
			output := minDeletionSize(p.strs)
			fmt.Printf("[input]: strs=%v       [output]:%v\n", p.strs, output)
			if output != q.ans960.one {
				t.Errorf("Expected %v, got %v", q.ans960.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

