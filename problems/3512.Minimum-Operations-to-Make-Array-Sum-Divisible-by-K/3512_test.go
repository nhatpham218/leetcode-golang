package leetcode

import (
	"fmt"
	"testing"
)

type question3512 struct {
	name string
	para3512
	ans3512
}

type para3512 struct {
	nums []int
	k    int
}

type ans3512 struct {
	output int
}

func Test_Problem3512(t *testing.T) {

	qs := []question3512{
		{
			name: "Example 1",
			para3512: para3512{
				nums: []int{3, 9, 7},
				k:    5,
			},
			ans3512: ans3512{output: 4},
		},
		{
			name: "Example 2",
			para3512: para3512{
				nums: []int{4, 1, 3},
				k:    4,
			},
			ans3512: ans3512{output: 0},
		},
		{
			name: "Example 3",
			para3512: para3512{
				nums: []int{3, 2},
				k:    6,
			},
			ans3512: ans3512{output: 5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3512------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans3512, q.para3512
			output := minOperations(p.nums, p.k)

			fmt.Printf("[input]: nums=%v k=%v       [output]:%v\n", p.nums, p.k, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

