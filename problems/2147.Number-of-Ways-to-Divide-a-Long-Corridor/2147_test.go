package leetcode

import (
	"fmt"
	"testing"
)

type question2147 struct {
	name string
	para2147
	ans2147
}

type para2147 struct {
	corridor string
}

type ans2147 struct {
	one int
}

func Test_Problem2147(t *testing.T) {

	qs := []question2147{
		{
			name: "example 1",
			para2147: para2147{
				corridor: "SSPPSPS",
			},
			ans2147: ans2147{3},
		},
		{
			name: "example 2",
			para2147: para2147{
				corridor: "PPSPSP",
			},
			ans2147: ans2147{1},
		},
		{
			name: "example 3",
			para2147: para2147{
				corridor: "S",
			},
			ans2147: ans2147{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2147------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2147, q.para2147
			output := numberOfWays(p.corridor)
			fmt.Printf("[input]: corridor=%v       [output]:%v\n", p.corridor, output)
			if output != q.ans2147.one {
				t.Errorf("Expected %v, got %v", q.ans2147.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

