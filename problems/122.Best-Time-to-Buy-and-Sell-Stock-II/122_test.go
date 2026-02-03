package leetcode

import (
	"fmt"
	"testing"
)

type question122 struct {
	name string
	para122
	ans122
}

type para122 struct {
	prices []int
}

type ans122 struct {
	one int
}

func Test_Problem122(t *testing.T) {

	qs := []question122{
		{
			name:    "example 1",
			para122: para122{prices: []int{7, 1, 5, 3, 6, 4}},
			ans122:  ans122{7},
		},
		{
			name:    "example 2",
			para122: para122{prices: []int{1, 2, 3, 4, 5}},
			ans122:  ans122{4},
		},
		{
			name:    "example 3",
			para122: para122{prices: []int{7, 6, 4, 3, 1}},
			ans122:  ans122{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 122------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans122, q.para122
			output := maxProfit(p.prices)
			fmt.Printf("[input]: prices=%v       [output]:%v\n", p.prices, output)
			if output != q.ans122.one {
				t.Errorf("Expected %v, got %v", q.ans122.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
