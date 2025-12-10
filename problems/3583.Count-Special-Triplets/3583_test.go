package leetcode

import (
	"fmt"
	"testing"
)

type question3583 struct {
	name string
	para3583
	ans3583
}

// para is parameter
// one represents the first parameter
type para3583 struct {
	nums []int
}

// ans is answer
// one represents the first answer
type ans3583 struct {
	one int
}

func Test_Problem3583(t *testing.T) {

	qs := []question3583{
		{
			name: "example 1",
			para3583: para3583{
				nums: []int{6, 3, 6},
			},
			ans3583: ans3583{1},
		},
		{
			name: "example 2",
			para3583: para3583{
				nums: []int{0, 1, 0, 0},
			},
			ans3583: ans3583{1},
		},
		{
			name: "example 3",
			para3583: para3583{
				nums: []int{8, 4, 2, 8, 4},
			},
			ans3583: ans3583{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3583------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3583, q.para3583
			output := specialTriplets(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans3583.one {
				t.Errorf("Expected %v, got %v", q.ans3583.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
