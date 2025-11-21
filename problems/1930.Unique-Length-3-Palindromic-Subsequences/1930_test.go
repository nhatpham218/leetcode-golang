package leetcode

import (
	"fmt"
	"testing"
)

type question1930 struct {
	name string
	para1930
	ans1930
}

type para1930 struct {
	s string
}

type ans1930 struct {
	output int
}

func Test_Problem1930(t *testing.T) {

	qs := []question1930{
		{
			name: "Example 1",
			para1930: para1930{
				s: "aabca",
			},
			ans1930: ans1930{output: 3},
		},
		{
			name: "Example 2",
			para1930: para1930{
				s: "adc",
			},
			ans1930: ans1930{output: 0},
		},
		{
			name: "Example 3",
			para1930: para1930{
				s: "bbcbaba",
			},
			ans1930: ans1930{output: 4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1930------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans1930, q.para1930
			output := countPalindromicSubsequence(p.s)

			fmt.Printf("[input]: s=%v       [output]:%v\n", p.s, output)
			if output != a.output {
				t.Errorf("expected %v, got %v", a.output, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
