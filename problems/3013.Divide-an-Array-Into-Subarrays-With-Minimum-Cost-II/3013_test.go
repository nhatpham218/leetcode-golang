package leetcode

import (
	"fmt"
	"testing"
)

type question3013 struct {
	name string
	para3013
	ans3013
}

// para is parameter
// one represents the first parameter
type para3013 struct {
	nums []int
	k    int
	dist int
}

// ans is answer
// one represents the first answer
type ans3013 struct {
	one int64
}

func Test_Problem3013(t *testing.T) {

	qs := []question3013{
		{
			name: "example 1",
			para3013: para3013{
				nums: []int{1, 3, 2, 6, 4, 2},
				k:    3,
				dist: 3,
			},
			ans3013: ans3013{5},
		},
		{
			name: "example 2",
			para3013: para3013{
				nums: []int{10, 1, 2, 2, 2, 1},
				k:    4,
				dist: 3,
			},
			ans3013: ans3013{15},
		},
		{
			name: "example 3",
			para3013: para3013{
				nums: []int{10, 8, 18, 9},
				k:    3,
				dist: 1,
			},
			ans3013: ans3013{36},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3013------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3013, q.para3013
			output := minimumCost(p.nums, p.k, p.dist)
			fmt.Printf("[input]: nums=%v, k=%v, dist=%v       [output]:%v\n", p.nums, p.k, p.dist, output)
			if output != q.ans3013.one {
				t.Errorf("Expected %v, got %v", q.ans3013.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
