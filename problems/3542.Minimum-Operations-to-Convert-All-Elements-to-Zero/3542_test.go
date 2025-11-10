package leetcode

import (
	"fmt"
	"testing"
)

type question3542 struct {
	name string
	para3542
	ans3542
}

type para3542 struct {
	nums []int
}

type ans3542 struct {
	one int
}

func Test_Problem3542(t *testing.T) {

	qs := []question3542{
		{
			name: "example 1",
			para3542: para3542{
				nums: []int{0, 2},
			},
			ans3542: ans3542{1},
		},
		{
			name: "example 2",
			para3542: para3542{
				nums: []int{3, 1, 2, 1},
			},
			ans3542: ans3542{3},
		},
		{
			name: "example 3",
			para3542: para3542{
				nums: []int{1, 2, 1, 2, 1, 2},
			},
			ans3542: ans3542{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3542------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3542, q.para3542
			output := minOperations(p.nums)
			fmt.Printf("[input]: nums=%v       [output]:%d\n", p.nums, output)
			if output != q.ans3542.one {
				t.Errorf("Expected %d, got %d", q.ans3542.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

