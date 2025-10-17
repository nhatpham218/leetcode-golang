package leetcode

import (
	"fmt"
	"testing"
)

type question3003 struct {
	name string
	para3003
	ans3003
}

type para3003 struct {
	s string
	k int
}

type ans3003 struct {
	one int
}

func Test_Problem3003(t *testing.T) {

	qs := []question3003{
		{
			name: "Example 1",
			para3003: para3003{
				s: "accca",
				k: 2,
			},
			ans3003: ans3003{
				one: 3,
			},
		},
		{
			name: "Example 2",
			para3003: para3003{
				s: "aabaab",
				k: 3,
			},
			ans3003: ans3003{
				one: 1,
			},
		},
		{
			name: "Example 3",
			para3003: para3003{
				s: "xxyz",
				k: 1,
			},
			ans3003: ans3003{
				one: 4,
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3003------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans3003, q.para3003
			output := maxPartitionsAfterOperations(p.s, p.k)
			fmt.Printf("[input]: s=%v, k=%v       [output]:%v\n", p.s, p.k, output)
			if output != a.one {
				t.Errorf("expected %v, got %v", a.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
