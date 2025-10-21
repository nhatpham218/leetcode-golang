package leetcode

import (
	"fmt"
	"testing"
)

type question3346 struct {
	name string
	para3346
	ans3346
}

type para3346 struct {
	nums          []int
	k             int
	numOperations int
}

type ans3346 struct {
	one int
}

func Test_Problem3346(t *testing.T) {

	qs := []question3346{
		{
			name: "Example 1",
			para3346: para3346{
				nums:          []int{1, 4, 5},
				k:             1,
				numOperations: 2,
			},
			ans3346: ans3346{
				one: 2,
			},
		},
		{
			name: "Example 2",
			para3346: para3346{
				nums:          []int{5, 11, 20, 20},
				k:             5,
				numOperations: 1,
			},
			ans3346: ans3346{
				one: 2,
			},
		},
		{
			name: "Example 3",
			para3346: para3346{
				nums:          []int{2, 49},
				k:             97,
				numOperations: 0,
			},
			ans3346: ans3346{
				one: 1,
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3346------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans3346, q.para3346
			output := maxFrequency(p.nums, p.k, p.numOperations)
			fmt.Printf("[input]: nums=%v k=%v numOperations=%v       [output]:%v\n", p.nums, p.k, p.numOperations, output)
			if output != a.one {
				t.Errorf("expected %v, got %v", a.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
