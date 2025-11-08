package leetcode

import (
	"fmt"
	"testing"
)

type question1611 struct {
	name string
	para1611
	ans1611
}

type para1611 struct {
	n int
}

type ans1611 struct {
	one int
}

func Test_Problem1611(t *testing.T) {

	qs := []question1611{
		{
			name: "example 1",
			para1611: para1611{
				n: 3,
			},
			ans1611: ans1611{2},
		},
		{
			name: "example 2",
			para1611: para1611{
				n: 6,
			},
			ans1611: ans1611{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1611------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1611, q.para1611
			output := minimumOneBitOperations(p.n)
			fmt.Printf("[input]: n=%v       [output]:%d\n", p.n, output)
			if output != q.ans1611.one {
				t.Errorf("Expected %d, got %d", q.ans1611.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
