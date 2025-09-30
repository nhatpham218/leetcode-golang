package leetcode

import (
	"fmt"
	"testing"
)

type question2221 struct {
	name string
	para2221
	ans2221
}

// para is parameter
// one represents the first parameter
type para2221 struct {
	nums []int
}

// ans is answer
// one represents the first answer
type ans2221 struct {
	one int
}

func Test_Problem2221(t *testing.T) {

	qs := []question2221{
		{
			name: "example 1",
			para2221: para2221{
				nums: []int{1, 2, 3, 4, 5},
			},
			ans2221: ans2221{8},
		},
		{
			name: "example 2",
			para2221: para2221{
				nums: []int{5},
			},
			ans2221: ans2221{5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2221------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2221, q.para2221
			// Create a copy of nums to avoid modifying the original
			numsCopy := make([]int, len(p.nums))
			copy(numsCopy, p.nums)
			output := triangularSum(numsCopy)
			fmt.Printf("[input]: nums=%v       [output]:%d\n", p.nums, output)
			if output != q.ans2221.one {
				t.Errorf("Expected %d, got %d", q.ans2221.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
