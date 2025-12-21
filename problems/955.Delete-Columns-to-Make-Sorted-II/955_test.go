package leetcode

import (
	"fmt"
	"testing"
)

type question955 struct {
	name string
	para955
	ans955
}

type para955 struct {
	strs []string
}

type ans955 struct {
	one int
}

func Test_Problem955(t *testing.T) {

	qs := []question955{
		{
			name: "example 1",
			para955: para955{
				strs: []string{"ca", "bb", "ac"},
			},
			ans955: ans955{1},
		},
		{
			name: "example 2",
			para955: para955{
				strs: []string{"xc", "yb", "za"},
			},
			ans955: ans955{0},
		},
		{
			name: "example 3",
			para955: para955{
				strs: []string{"zyx", "wvu", "tsr"},
			},
			ans955: ans955{3},
		},
		{
			name: "example 4",
			para955: para955{
				strs: []string{"jwkwdc", "etukoz"},
			},
			ans955: ans955{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 955------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans955, q.para955
			output := minDeletionSize(p.strs)
			fmt.Printf("[input]: strs=%v       [output]:%v\n", p.strs, output)
			if output != q.ans955.one {
				t.Errorf("Expected %v, got %v", q.ans955.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
