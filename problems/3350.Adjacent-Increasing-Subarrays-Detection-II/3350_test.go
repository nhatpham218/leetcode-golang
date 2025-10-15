package leetcode

import (
	"fmt"
	"testing"
)

type question3350 struct {
	name string
	para3350
	ans3350
}

type para3350 struct {
	nums []int
}

type ans3350 struct {
	one int
}

func Test_Problem3350(t *testing.T) {

	qs := []question3350{
		{
			name: "Example 1",
			para3350: para3350{
				nums: []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1},
			},
			ans3350: ans3350{
				one: 3,
			},
		},
		{
			name: "Example 2",
			para3350: para3350{
				nums: []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7},
			},
			ans3350: ans3350{
				one: 2,
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3350------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans3350, q.para3350
			output := maxIncreasingSubarrays(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != a.one {
				t.Errorf("expected %v, got %v", a.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
