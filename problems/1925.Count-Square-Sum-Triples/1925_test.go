package leetcode

import (
	"fmt"
	"testing"
)

type question1925 struct {
	name string
	para1925
	ans1925
}

// para is parameter
// one represents the first parameter
type para1925 struct {
	n int
}

// ans is answer
// one represents the first answer
type ans1925 struct {
	one int
}

func Test_Problem1925(t *testing.T) {

	qs := []question1925{
		{
			name: "example 1",
			para1925: para1925{
				n: 5,
			},
			ans1925: ans1925{2},
		},
		{
			name: "example 2",
			para1925: para1925{
				n: 10,
			},
			ans1925: ans1925{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1925------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1925, q.para1925
			output := countTriples(p.n)
			fmt.Printf("[input]: n=%v       [output]:%v\n", p.n, output)
			if output != q.ans1925.one {
				t.Errorf("Expected %v, got %v", q.ans1925.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
