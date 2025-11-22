package leetcode

import (
	"fmt"
	"testing"
)

type question3190 struct {
	name string
	para3190
	ans3190
}

type para3190 struct {
	nums []int
}

type ans3190 struct {
	one int
}

func Test_Problem3190(t *testing.T) {

	qs := []question3190{
		{
			name: "example 1",
			para3190: para3190{
				nums: []int{1, 2, 3, 4},
			},
			ans3190: ans3190{3},
		},
		{
			name: "example 2",
			para3190: para3190{
				nums: []int{3, 6, 9},
			},
			ans3190: ans3190{0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3190------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3190, q.para3190
			output := minimumOperations(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%v\n", p.nums, output)
			if output != q.ans3190.one {
				t.Errorf("Expected %v, got %v", q.ans3190.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

