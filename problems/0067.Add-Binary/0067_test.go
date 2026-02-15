package leetcode

import (
	"fmt"
	"testing"
)

type question67 struct {
	name string
	para67
	ans67
}

type para67 struct {
	a string
	b string
}

type ans67 struct {
	one string
}

func Test_Problem67(t *testing.T) {

	qs := []question67{
		{
			name:   "example 1",
			para67: para67{a: "11", b: "1"},
			ans67:  ans67{one: "100"},
		},
		{
			name:   "example 2",
			para67: para67{a: "1010", b: "1011"},
			ans67:  ans67{one: "10101"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 67------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans67, q.para67
			output := addBinary(p.a, p.b)
			fmt.Printf("[input]: a=%v, b=%v       [output]:%v\n", p.a, p.b, output)
			if output != q.ans67.one {
				t.Errorf("Expected %v, got %v", q.ans67.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
