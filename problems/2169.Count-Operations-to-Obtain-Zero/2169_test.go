package leetcode

import (
	"fmt"
	"testing"
)

type question2169 struct {
	name string
	para2169
	ans2169
}

type para2169 struct {
	num1 int
	num2 int
}

type ans2169 struct {
	one int
}

func Test_Problem2169(t *testing.T) {

	qs := []question2169{
		{
			name: "example 1",
			para2169: para2169{
				num1: 2,
				num2: 3,
			},
			ans2169: ans2169{3},
		},
		{
			name: "example 2",
			para2169: para2169{
				num1: 10,
				num2: 10,
			},
			ans2169: ans2169{1},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2169------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2169, q.para2169
			output := countOperations(p.num1, p.num2)
			fmt.Printf("[input]: num1=%v, num2=%v       [output]:%d\n", p.num1, p.num2, output)
			if output != q.ans2169.one {
				t.Errorf("Expected %d, got %d", q.ans2169.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

