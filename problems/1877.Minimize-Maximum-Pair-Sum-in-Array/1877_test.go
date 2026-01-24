package leetcode

import (
	"fmt"
	"testing"
)

type question1877 struct {
	name string
	para1877
	ans1877
}

// para is parameter
// one represents the first parameter
type para1877 struct {
	nums []int
}

// ans is answer
// one represents the first answer
type ans1877 struct {
	one int
}

func Test_Problem1877(t *testing.T) {

	qs := []question1877{
		{
			name: "example 1",
			para1877: para1877{
				nums: []int{3, 5, 2, 3},
			},
			ans1877: ans1877{7},
		},
		{
			name: "example 2",
			para1877: para1877{
				nums: []int{3, 5, 4, 2, 4, 6},
			},
			ans1877: ans1877{8},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1877------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1877, q.para1877
			output := minPairSum(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans1877.one {
				t.Errorf("Expected %v, got %v", q.ans1877.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
