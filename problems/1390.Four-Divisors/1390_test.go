package leetcode

import (
	"fmt"
	"testing"
)

type question1390 struct {
	name string
	para1390
	ans1390
}

type para1390 struct {
	nums []int
}

type ans1390 struct {
	one int
}

func Test_Problem1390(t *testing.T) {

	qs := []question1390{
		{
			name:     "example 1",
			para1390: para1390{nums: []int{21, 4, 7}},
			ans1390:  ans1390{32},
		},
		{
			name:     "example 2",
			para1390: para1390{nums: []int{21, 21}},
			ans1390:  ans1390{64},
		},
		{
			name:     "example 3",
			para1390: para1390{nums: []int{1, 2, 3, 4, 5}},
			ans1390:  ans1390{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1390------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1390, q.para1390
			output := sumFourDivisors(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans1390.one {
				t.Errorf("Expected %v, got %v", q.ans1390.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

