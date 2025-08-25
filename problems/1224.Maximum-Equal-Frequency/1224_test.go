package leetcode

import (
	"fmt"
	"testing"
)

type question1224 struct {
	para1224
	ans1224
}

// para is parameter
// one represents the first parameter
type para1224 struct {
	nums []int
}

// ans is answer
// one represents the first answer
type ans1224 struct {
	one int
}

func Test_Problem1224(t *testing.T) {

	qs := []question1224{
		{
			para1224{[]int{2, 2, 1, 1, 5, 3, 3, 5}},
			ans1224{7},
		},

		{
			para1224{[]int{1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 5}},
			ans1224{13},
		},
		{
			para1224{[]int{1, 1, 1, 2, 2, 2}},
			ans1224{5},
		},
		{
			para1224{[]int{10, 2, 8, 9, 3, 8, 1, 5, 2, 3, 7, 6}},
			ans1224{8},
		},
		{
			para1224{[]int{1, 1}},
			ans1224{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1224------------------------\n")

	for i, q := range qs {
		t.Run(fmt.Sprintf("Case %d", i), func(t *testing.T) {
			_, p := q.ans1224, q.para1224
			output := maxEqualFreq(p.nums)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			if output != q.ans1224.one {
				t.Errorf("Expected %d, got %d", q.ans1224.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
