package leetcode

import (
	"fmt"
	"testing"
)

type question3010 struct {
	name string
	para3010
	ans3010
}

// para is parameter
// one represents the first parameter
type para3010 struct {
	nums []int
}

// ans is answer
// one represents the first answer
type ans3010 struct {
	one int
}

func Test_Problem3010(t *testing.T) {

	qs := []question3010{
		{
			name: "example 1",
			para3010: para3010{
				nums: []int{1, 2, 3, 12},
			},
			ans3010: ans3010{6},
		},
		{
			name: "example 2",
			para3010: para3010{
				nums: []int{5, 4, 3},
			},
			ans3010: ans3010{12},
		},
		{
			name: "example 3",
			para3010: para3010{
				nums: []int{10, 3, 1, 1},
			},
			ans3010: ans3010{12},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3010------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3010, q.para3010
			output := minimumCost(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans3010.one {
				t.Errorf("Expected %v, got %v", q.ans3010.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
