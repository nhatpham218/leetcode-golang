package leetcode

import (
	"fmt"
	"testing"
)

type question1523 struct {
	name string
	para1523
	ans1523
}

type para1523 struct {
	low  int
	high int
}

type ans1523 struct {
	one int
}

func Test_Problem1523(t *testing.T) {

	qs := []question1523{
		{
			name: "example 1",
			para1523: para1523{
				low:  3,
				high: 7,
			},
			ans1523: ans1523{3},
		},
		{
			name: "example 2",
			para1523: para1523{
				low:  8,
				high: 10,
			},
			ans1523: ans1523{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1523------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1523, q.para1523
			output := countOdds(p.low, p.high)
			fmt.Printf("[input]: low=%v, high=%v       [output]:%v\n", p.low, p.high, output)
			if output != q.ans1523.one {
				t.Errorf("Expected %v, got %v", q.ans1523.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

