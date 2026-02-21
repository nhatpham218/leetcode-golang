package leetcode

import (
	"fmt"
	"testing"
)

type question0762 struct {
	name string
	para0762
	ans0762
}

type para0762 struct {
	left  int
	right int
}

type ans0762 struct {
	one int
}

func Test_Problem0762(t *testing.T) {
	qs := []question0762{
		{
			name:     "example 1",
			para0762: para0762{left: 6, right: 10},
			ans0762:  ans0762{one: 4},
		},
		{
			name:     "example 2",
			para0762: para0762{left: 10, right: 15},
			ans0762:  ans0762{one: 5},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0762------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0762, q.para0762
			output := countPrimeSetBits(p.left, p.right)
			fmt.Printf("[input]: left=%v, right=%v       [output]:%v\n", p.left, p.right, output)
			if output != q.ans0762.one {
				t.Errorf("Expected %v, got %v", q.ans0762.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
