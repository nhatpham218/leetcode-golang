package leetcode

import (
	"fmt"
	"testing"
)

type question3721 struct {
	name string
	para3721
	ans3721
}

// para is parameter
// one represents the first parameter
type para3721 struct {
	nums []int
}

// ans is answer
// one represents the first answer
type ans3721 struct {
	one int
}

func Test_Problem3721(t *testing.T) {

	qs := []question3721{
		{
			name:     "example 1",
			para3721: para3721{nums: []int{2, 5, 4, 3}},
			ans3721:  ans3721{one: 4},
		},
		{
			name:     "example 2",
			para3721: para3721{nums: []int{3, 2, 2, 5, 4}},
			ans3721:  ans3721{one: 5},
		},
		{
			name:     "example 3",
			para3721: para3721{nums: []int{1, 2, 3, 2}},
			ans3721:  ans3721{one: 3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3721------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3721, q.para3721
			output := longestBalanced(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans3721.one {
				t.Errorf("Expected %v, got %v", q.ans3721.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
