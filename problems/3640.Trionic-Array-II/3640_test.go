package leetcode

import (
	"fmt"
	"testing"
)

type question3640 struct {
	name string
	para3640
	ans3640
}

type para3640 struct {
	nums []int
}

type ans3640 struct {
	one int64
}

func Test_Problem3640(t *testing.T) {

	qs := []question3640{
		{
			name:     "example 1",
			para3640: para3640{nums: []int{0, -2, -1, -3, 0, 2, -1}},
			ans3640:  ans3640{-4},
		},
		{
			name:     "example 2",
			para3640: para3640{nums: []int{1, 4, 2, 7}},
			ans3640:  ans3640{14},
		},
		{
			name:     "test 506",
			para3640: para3640{nums: []int{-754, 167, -172, 202, 735, -941, 992}},
			ans3640:  ans3640{988},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3640------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3640, q.para3640
			output := maxSumTrionic(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans3640.one {
				t.Errorf("Expected %v, got %v", q.ans3640.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
