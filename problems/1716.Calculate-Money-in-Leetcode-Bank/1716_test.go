package leetcode

import (
	"fmt"
	"testing"
)

type question1716 struct {
	name string
	para1716
	ans1716
}

type para1716 struct {
	n int
}

type ans1716 struct {
	one int
}

func Test_Problem1716(t *testing.T) {

	qs := []question1716{
		{
			name: "Example 1",
			para1716: para1716{
				n: 4,
			},
			ans1716: ans1716{
				one: 10,
			},
		},
		{
			name: "Example 2",
			para1716: para1716{
				n: 10,
			},
			ans1716: ans1716{
				one: 37,
			},
		},
		{
			name: "Example 3",
			para1716: para1716{
				n: 20,
			},
			ans1716: ans1716{
				one: 96,
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 1716------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			a, p := q.ans1716, q.para1716
			output := totalMoney(p.n)
			fmt.Printf("[input]: n=%v       [output]:%v\n", p.n, output)
			if output != a.one {
				t.Errorf("expected %v, got %v", a.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
