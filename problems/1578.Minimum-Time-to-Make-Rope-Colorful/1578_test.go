package leetcode

import (
	"fmt"
	"testing"
)

type question1578 struct {
	name string
	para1578
	ans1578
}

// para is parameter
// one represents the first parameter
type para1578 struct {
	colors     string
	neededTime []int
}

// ans is answer
// one represents the first answer
type ans1578 struct {
	one int
}

func Test_Problem1578(t *testing.T) {

	qs := []question1578{
		{
			name: "example 1",
			para1578: para1578{
				colors:     "abaac",
				neededTime: []int{1, 2, 3, 4, 5},
			},
			ans1578: ans1578{3},
		},
		{
			name: "example 2",
			para1578: para1578{
				colors:     "abc",
				neededTime: []int{1, 2, 3},
			},
			ans1578: ans1578{0},
		},
		{
			name: "example 3",
			para1578: para1578{
				colors:     "aabaa",
				neededTime: []int{1, 2, 3, 4, 1},
			},
			ans1578: ans1578{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1578------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1578, q.para1578
			output := minCost(p.colors, p.neededTime)
			fmt.Printf("[input]: colors=%v neededTime=%v       [output]:%v\n", p.colors, p.neededTime, output)
			if output != q.ans1578.one {
				t.Errorf("Expected %v, got %v", q.ans1578.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
