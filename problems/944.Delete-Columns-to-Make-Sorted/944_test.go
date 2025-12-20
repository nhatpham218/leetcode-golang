package leetcode

import (
	"fmt"
	"testing"
)

type question944 struct {
	name string
	para944
	ans944
}

// para is parameter
// one represents the first parameter
type para944 struct {
	strs []string
}

// ans is answer
// one represents the first answer
type ans944 struct {
	one int
}

func Test_Problem944(t *testing.T) {

	qs := []question944{
		{
			name: "example 1",
			para944: para944{
				strs: []string{"cba", "daf", "ghi"},
			},
			ans944: ans944{1},
		},
		{
			name: "example 2",
			para944: para944{
				strs: []string{"a", "b"},
			},
			ans944: ans944{0},
		},
		{
			name: "example 3",
			para944: para944{
				strs: []string{"zyx", "wvu", "tsr"},
			},
			ans944: ans944{3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 944------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans944, q.para944
			output := minDeletionSize(p.strs)
			fmt.Printf("[input]: strs=%v       [output]:%v\n", p.strs, output)
			if output != q.ans944.one {
				t.Errorf("Expected %v, got %v", q.ans944.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
