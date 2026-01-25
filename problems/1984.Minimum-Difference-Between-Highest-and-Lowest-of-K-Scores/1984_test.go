package leetcode

import (
	"fmt"
	"testing"
)

type question1984 struct {
	name string
	para1984
	ans1984
}

// para is parameter
// one represents the first parameter
type para1984 struct {
	nums []int
	k    int
}

// ans is answer
// one represents the first answer
type ans1984 struct {
	one int
}

func Test_Problem1984(t *testing.T) {

	qs := []question1984{
		{
			name: "example 1",
			para1984: para1984{
				nums: []int{90},
				k:    1,
			},
			ans1984: ans1984{0},
		},
		{
			name: "example 2",
			para1984: para1984{
				nums: []int{9, 4, 1, 7},
				k:    2,
			},
			ans1984: ans1984{2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1984------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans1984, q.para1984
			output := minimumDifference(p.nums, p.k)
			fmt.Printf("[input]: nums=%v, k=%v       [output]:%v\n", p.nums, p.k, output)
			if output != q.ans1984.one {
				t.Errorf("Expected %v, got %v", q.ans1984.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
