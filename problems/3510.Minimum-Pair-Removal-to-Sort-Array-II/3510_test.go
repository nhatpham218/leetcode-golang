package leetcode

import (
	"fmt"
	"testing"
)

type question3510 struct {
	name string
	para3510
	ans3510
}

// para is parameter
// one represents the first parameter
type para3510 struct {
	nums []int
}

// ans is answer
// one represents the first answer
type ans3510 struct {
	one int
}

func Test_Problem3510(t *testing.T) {

	qs := []question3510{
		{
			name: "example 1",
			para3510: para3510{
				nums: []int{5, 2, 3, 1},
			},
			ans3510: ans3510{2},
		},
		{
			name: "example 2",
			para3510: para3510{
				nums: []int{1, 2, 2},
			},
			ans3510: ans3510{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3510------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3510, q.para3510
			output := minimumPairRemoval(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans3510.one {
				t.Errorf("Expected %v, got %v", q.ans3510.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
