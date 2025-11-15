package leetcode

import (
	"fmt"
	"testing"
)

type question3234 struct {
	name string
	para3234
	ans3234
}

type para3234 struct {
	s string
}

type ans3234 struct {
	output int
}

func Test_Problem3234(t *testing.T) {

	qs := []question3234{
		{
			name:     "Example 1",
			para3234: para3234{s: "00011"},
			ans3234:  ans3234{output: 5},
		},
		{
			name:     "Example 2",
			para3234: para3234{s: "101101"},
			ans3234:  ans3234{output: 16},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3234------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans3234, q.para3234
			output := numberOfSubstrings(p.s)

			fmt.Printf("[input]: s=%v       [output]:%v\n", p.s, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

