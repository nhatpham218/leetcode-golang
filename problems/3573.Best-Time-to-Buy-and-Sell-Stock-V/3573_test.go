package leetcode

import (
	"fmt"
	"testing"
)

type question3573 struct {
	name string
	para3573
	ans3573
}

// para is parameter
// one represents the first parameter
type para3573 struct {
	prices []int
	k      int
}

// ans is answer
// one represents the first answer
type ans3573 struct {
	one int64
}

func Test_Problem3573(t *testing.T) {

	qs := []question3573{
		{
			name: "example 1",
			para3573: para3573{
				prices: []int{1, 7, 9, 8, 2},
				k:      2,
			},
			ans3573: ans3573{14},
		},
		{
			name: "example 2",
			para3573: para3573{
				prices: []int{12, 16, 19, 19, 8, 1, 19, 13, 9},
				k:      3,
			},
			ans3573: ans3573{36},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3573------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3573, q.para3573
			output := maximumProfit(p.prices, p.k)
			fmt.Printf("[input]: prices=%v, k=%v       [output]:%v\n", p.prices, p.k, output)
			if output != q.ans3573.one {
				t.Errorf("Expected %v, got %v", q.ans3573.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
