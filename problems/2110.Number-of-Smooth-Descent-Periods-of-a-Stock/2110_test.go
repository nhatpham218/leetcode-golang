package leetcode

import (
	"fmt"
	"testing"
)

type question2110 struct {
	name string
	para2110
	ans2110
}

// para is parameter
// one represents the first parameter
type para2110 struct {
	prices []int
}

// ans is answer
// one represents the first answer
type ans2110 struct {
	one int64
}

func Test_Problem2110(t *testing.T) {

	qs := []question2110{
		{
			name: "example 1",
			para2110: para2110{
				prices: []int{3, 2, 1, 4},
			},
			ans2110: ans2110{7},
		},
		{
			name: "example 2",
			para2110: para2110{
				prices: []int{8, 6, 7, 7},
			},
			ans2110: ans2110{4},
		},
		{
			name: "example 3",
			para2110: para2110{
				prices: []int{1},
			},
			ans2110: ans2110{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2110------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2110, q.para2110
			output := getDescentPeriods(p.prices)
			fmt.Printf("[input]: prices=%v       [output]:%v\n", p.prices, output)
			if output != q.ans2110.one {
				t.Errorf("Expected %v, got %v", q.ans2110.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
