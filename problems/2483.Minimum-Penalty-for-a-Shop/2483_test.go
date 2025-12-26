package leetcode

import (
	"fmt"
	"testing"
)

type question2483 struct {
	name string
	para2483
	ans2483
}

type para2483 struct {
	customers string
}

type ans2483 struct {
	one int
}

func Test_Problem2483(t *testing.T) {

	qs := []question2483{
		{
			name:     "example 1",
			para2483: para2483{customers: "YYNY"},
			ans2483:  ans2483{2},
		},
		{
			name:     "example 2",
			para2483: para2483{customers: "NNNNN"},
			ans2483:  ans2483{0},
		},
		{
			name:     "example 3",
			para2483: para2483{customers: "YYYY"},
			ans2483:  ans2483{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 2483------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans2483, q.para2483
			output := bestClosingTime(p.customers)
			fmt.Printf("[input]: customers=%v       [output]:%v\n", p.customers, output)
			if output != q.ans2483.one {
				t.Errorf("Expected %v, got %v", q.ans2483.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}

