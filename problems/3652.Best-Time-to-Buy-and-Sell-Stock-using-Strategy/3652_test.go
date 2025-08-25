package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3652 struct {
	name string
	para3652
	ans3652
}

// para is parameter
// prices represents the first parameter
// strategy represents the second parameter
// k represents the third parameter
type para3652 struct {
	prices   []int
	strategy []int
	k        int
}

// ans is answer
// one represents the first answer
type ans3652 struct {
	one int64
}

func Test_Problem3652(t *testing.T) {

	qs := []question3652{
		{
			name:     "example 1",
			para3652: para3652{[]int{4, 2, 8}, []int{-1, 0, 1}, 2},
			ans3652:  ans3652{10},
		},
		{
			name:     "example 2",
			para3652: para3652{[]int{5, 4, 3}, []int{1, 1, 0}, 2},
			ans3652:  ans3652{9},
		},
		{
			name:     "example 3",
			para3652: para3652{[]int{4, 7, 13}, []int{-1, -1, 0}, 2},
			ans3652:  ans3652{9},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3652------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3652, q.para3652
			output := maxProfit(p.prices, p.strategy, p.k)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			assert.Equal(t, q.ans3652.one, output, "maxProfit(%v, %v, %d) should return %d", p.prices, p.strategy, p.k, q.ans3652.one)
		})
	}
	fmt.Printf("\n\n\n")
}
