package leetcode

import (
	"fmt"
	"testing"
)

type question3507 struct {
	name string
	para3507
	ans3507
}

// para is parameter
// one represents the first parameter
type para3507 struct {
	nums []int
}

// ans is answer
// one represents the first answer
type ans3507 struct {
	one int
}

func Test_Problem3507(t *testing.T) {

	qs := []question3507{
		{
			name: "example 1",
			para3507: para3507{
				nums: []int{5, 2, 3, 1},
			},
			ans3507: ans3507{2},
		},
		{
			name: "example 2",
			para3507: para3507{
				nums: []int{1, 2, 2},
			},
			ans3507: ans3507{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3507------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3507, q.para3507
			output := minimumPairRemoval(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans3507.one {
				t.Errorf("Expected %v, got %v", q.ans3507.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
