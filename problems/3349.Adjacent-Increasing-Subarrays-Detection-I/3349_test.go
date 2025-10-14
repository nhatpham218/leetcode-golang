package leetcode

import (
	"fmt"
	"testing"
)

type question3349 struct {
	name string
	para3349
	ans3349
}

type para3349 struct {
	nums []int
	k    int
}

type ans3349 struct {
	one bool
}

func Test_Problem3349(t *testing.T) {

	qs := []question3349{
		{
			name: "Example 1",
			para3349: para3349{
				nums: []int{2, 5, 7, 8, 9, 2, 3, 4, 3, 1},
				k:    3,
			},
			ans3349: ans3349{
				one: true,
			},
		},
		{
			name: "Example 2",
			para3349: para3349{
				nums: []int{1, 2, 3, 4, 4, 4, 4, 5, 6, 7},
				k:    5,
			},
			ans3349: ans3349{
				one: false,
			},
		},
		{
			name: "Edge Case - k=1",
			para3349: para3349{
				nums: []int{-15, 19},
				k:    1,
			},
			ans3349: ans3349{
				one: true,
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3349------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans3349, q.para3349
			output := hasIncreasingSubarrays(p.nums, p.k)
			fmt.Printf("[input]: nums=%v, k=%v       [output]:%v\n", p.nums, p.k, output)
			if output != a.one {
				t.Errorf("expected %v, got %v", a.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
