package leetcode

import (
	"fmt"
	"testing"
)

type question121 struct {
	name string
	para121
	ans121
}

type para121 struct {
	prices []int
}

type ans121 struct {
	one int
}

func Test_Problem121(t *testing.T) {

	qs := []question121{
		{
			name:    "example 1",
			para121: para121{prices: []int{7, 1, 5, 3, 6, 4}},
			ans121:  ans121{5},
		},
		{
			name:    "example 2",
			para121: para121{prices: []int{7, 6, 4, 3, 1}},
			ans121:  ans121{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 121------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans121, q.para121
			output := maxProfit(p.prices)
			fmt.Printf("[input]: prices=%v       [output]:%v\n", p.prices, output)
			if output != q.ans121.one {
				t.Errorf("Expected %v, got %v", q.ans121.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
