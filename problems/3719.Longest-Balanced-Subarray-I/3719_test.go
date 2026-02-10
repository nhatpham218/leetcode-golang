package leetcode

import (
	"fmt"
	"testing"
)

type question3719 struct {
	name string
	para3719
	ans3719
}

// para is parameter
// one represents the first parameter
type para3719 struct {
	nums []int
}

// ans is answer
// one represents the first answer
type ans3719 struct {
	one int
}

func Test_Problem3719(t *testing.T) {

	qs := []question3719{
		{
			name:    "example 1",
			para3719: para3719{nums: []int{2, 5, 4, 3}},
			ans3719:  ans3719{one: 4},
		},
		{
			name:    "example 2",
			para3719: para3719{nums: []int{3, 2, 2, 5, 4}},
			ans3719:  ans3719{one: 5},
		},
		{
			name:    "example 3",
			para3719: para3719{nums: []int{1, 2, 3, 2}},
			ans3719:  ans3719{one: 3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3719------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3719, q.para3719
			output := longestBalanced(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans3719.one {
				t.Errorf("Expected %v, got %v", q.ans3719.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
