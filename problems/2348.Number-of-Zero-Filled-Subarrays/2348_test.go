package leetcode

import (
	"fmt"
	"testing"
)

type question2348 struct {
	para2348
	ans2348
}

// para is parameter
// one represents the first parameter
type para2348 struct {
	nums []int
}

// ans is answer
// one represents the first answer
type ans2348 struct {
	one int64
}

func Test_Problem2348(t *testing.T) {

	qs := []question2348{
		{
			para2348{[]int{1, 3, 0, 0, 2, 0, 0, 4}},
			ans2348{6},
		},

		{
			para2348{[]int{0, 0, 0, 2, 0, 0}},
			ans2348{9},
		},

		{
			para2348{[]int{2, 10, 2019}},
			ans2348{0},
		},
		// If you need more tests, you can copy the elements above.
	}

	fmt.Printf("------------------------Leetcode Problem 2348------------------------\n")

	for _, q := range qs {
		_, p := q.ans2348, q.para2348
		fmt.Printf("[input]:%v       [output]:%v\n", p, zeroFilledSubarray(p.nums))
	}
	fmt.Printf("\n\n\n")
}
