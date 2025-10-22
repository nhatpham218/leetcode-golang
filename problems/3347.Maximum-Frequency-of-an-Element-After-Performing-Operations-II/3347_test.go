package leetcode

import (
	"fmt"
	"testing"
)

type question3347 struct {
	name string
	para3347
	ans3347
}

type para3347 struct {
	nums          []int
	k             int
	numOperations int
}

type ans3347 struct {
	one int
}

func Test_Problem3347(t *testing.T) {

	qs := []question3347{
		{
			name: "Example 1",
			para3347: para3347{
				nums:          []int{1, 4, 5},
				k:             1,
				numOperations: 2,
			},
			ans3347: ans3347{
				one: 2,
			},
		},
		{
			name: "Example 2",
			para3347: para3347{
				nums:          []int{5, 11, 20, 20},
				k:             5,
				numOperations: 1,
			},
			ans3347: ans3347{
				one: 2,
			},
		},
		{
			name: "Failing Case",
			para3347: para3347{
				nums:          []int{5, 64},
				k:             42,
				numOperations: 2,
			},
			ans3347: ans3347{
				one: 2,
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3347------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans3347, q.para3347
			output := maxFrequency(p.nums, p.k, p.numOperations)
			fmt.Printf("[input]: nums=%v k=%v numOperations=%v       [output]:%v\n", p.nums, p.k, p.numOperations, output)
			if output != a.one {
				t.Errorf("expected %v, got %v", a.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
