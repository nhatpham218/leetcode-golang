package leetcode

import (
	"fmt"
	"testing"
)

type question3381 struct {
	name string
	para3381
	ans3381
}

type para3381 struct {
	nums []int
	k    int
}

type ans3381 struct {
	output int64
}

func Test_Problem3381(t *testing.T) {

	qs := []question3381{
		{
			name: "Example 1",
			para3381: para3381{
				nums: []int{1, 2},
				k:    1,
			},
			ans3381: ans3381{output: 3},
		},
		{
			name: "Example 2",
			para3381: para3381{
				nums: []int{-1, -2, -3, -4, -5},
				k:    4,
			},
			ans3381: ans3381{output: -10},
		},
		{
			name: "Example 3",
			para3381: para3381{
				nums: []int{-5, 1, 2, -3, 4},
				k:    2,
			},
			ans3381: ans3381{output: 4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3381------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans3381, q.para3381
			output := maxSubarraySum(p.nums, p.k)

			fmt.Printf("[input]: nums=%v k=%v       [output]:%v\n", p.nums, p.k, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

