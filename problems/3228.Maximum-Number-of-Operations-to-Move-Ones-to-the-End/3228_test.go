package leetcode

import (
	"fmt"
	"testing"
)

type question3228 struct {
	name string
	para3228
	ans3228
}

type para3228 struct {
	s string
}

type ans3228 struct {
	output int
}

func Test_Problem3228(t *testing.T) {

	qs := []question3228{
		{
			name: "Example 1",
			para3228: para3228{
				s: "1001101",
			},
			ans3228: ans3228{output: 4},
		},
		{
			name: "Example 2",
			para3228: para3228{
				s: "00111",
			},
			ans3228: ans3228{output: 0},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3228------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans3228, q.para3228
			output := maxOperations(p.s)

			fmt.Printf("[input]: s=%v       [output]:%v\n", p.s, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

