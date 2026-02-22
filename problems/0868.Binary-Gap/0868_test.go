package leetcode

import (
	"fmt"
	"testing"
)

type question0868 struct {
	name string
	para0868
	ans0868
}

type para0868 struct {
	n int
}

type ans0868 struct {
	one int
}

func Test_Problem0868(t *testing.T) {
	qs := []question0868{
		{
			name:     "example 1",
			para0868: para0868{n: 22},
			ans0868:  ans0868{one: 2},
		},
		{
			name:     "example 2",
			para0868: para0868{n: 8},
			ans0868:  ans0868{one: 0},
		},
		{
			name:     "example 3",
			para0868: para0868{n: 5},
			ans0868:  ans0868{one: 2},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0868------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0868, q.para0868
			output := binaryGap(p.n)
			fmt.Printf("[input]: n=%v       [output]:%v\n", p.n, output)
			if output != q.ans0868.one {
				t.Errorf("Expected %v, got %v", q.ans0868.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
