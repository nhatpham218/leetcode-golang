package leetcode

import (
	"fmt"
	"testing"
)

type question2011 struct {
	name string
	para2011
	ans2011
}

type para2011 struct {
	operations []string
}

type ans2011 struct {
	one int
}

func Test_Problem2011(t *testing.T) {

	qs := []question2011{
		{
			name: "Example 1",
			para2011: para2011{
				operations: []string{"--X", "X++", "X++"},
			},
			ans2011: ans2011{
				one: 1,
			},
		},
		{
			name: "Example 2",
			para2011: para2011{
				operations: []string{"++X", "++X", "X++"},
			},
			ans2011: ans2011{
				one: 3,
			},
		},
		{
			name: "Example 3",
			para2011: para2011{
				operations: []string{"X++", "++X", "--X", "X--"},
			},
			ans2011: ans2011{
				one: 0,
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2011------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans2011, q.para2011
			output := finalValueAfterOperations(p.operations)
			fmt.Printf("[input]: operations=%v       [output]:%v\n", p.operations, output)
			if output != a.one {
				t.Errorf("expected %v, got %v", a.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
