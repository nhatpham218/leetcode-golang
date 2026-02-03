package leetcode

import (
	"fmt"
	"testing"
)

type question3637 struct {
	name string
	para3637
	ans3637
}

// para is parameter
// one represents the first parameter
type para3637 struct {
	nums []int
}

// ans is answer
// one represents the first answer
type ans3637 struct {
	one bool
}

func Test_Problem3637(t *testing.T) {

	qs := []question3637{
		{
			name: "example 1",
			para3637: para3637{
				nums: []int{1, 3, 5, 4, 2, 6},
			},
			ans3637: ans3637{true},
		},
		{
			name: "example 2",
			para3637: para3637{
				nums: []int{2, 1, 3},
			},
			ans3637: ans3637{false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3637------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3637, q.para3637
			output := isTrionic(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans3637.one {
				t.Errorf("Expected %v, got %v", q.ans3637.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
