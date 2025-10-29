package leetcode

import (
	"fmt"
	"testing"
)

type question3370 struct {
	name string
	para3370
	ans3370
}

type para3370 struct {
	n int
}

type ans3370 struct {
	output int
}

func Test_Problem3370(t *testing.T) {

	qs := []question3370{
		{
			name:     "Example 1",
			para3370: para3370{n: 5},
			ans3370:  ans3370{output: 7},
		},
		{
			name:     "Example 2",
			para3370: para3370{n: 10},
			ans3370:  ans3370{output: 15},
		},
		{
			name:     "Example 3",
			para3370: para3370{n: 3},
			ans3370:  ans3370{output: 3},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3370------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans3370, q.para3370
			output := smallestNumber(p.n)

			fmt.Printf("[input]: %v       [output]:%v\n", p.n, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
