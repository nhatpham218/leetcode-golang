package leetcode

import (
	"fmt"
	"testing"
)

type question2598 struct {
	name string
	para2598
	ans2598
}

type para2598 struct {
	nums  []int
	value int
}

type ans2598 struct {
	one int
}

func Test_Problem2598(t *testing.T) {

	qs := []question2598{
		{
			name: "Example 1",
			para2598: para2598{
				nums:  []int{1, -10, 7, 13, 6, 8},
				value: 5,
			},
			ans2598: ans2598{
				one: 4,
			},
		},
		{
			name: "Example 2",
			para2598: para2598{
				nums:  []int{1, -10, 7, 13, 6, 8},
				value: 7,
			},
			ans2598: ans2598{
				one: 2,
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2598------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans2598, q.para2598
			output := findSmallestInteger(p.nums, p.value)
			fmt.Printf("[input]: nums=%v, value=%v       [output]:%v\n", p.nums, p.value, output)
			if output != a.one {
				t.Errorf("expected %v, got %v", a.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
