package leetcode

import (
	"fmt"
	"testing"
)

type question0693 struct {
	name    string
	para0693
	ans0693
}

type para0693 struct {
	n int
}

type ans0693 struct {
	one bool
}

func Test_Problem0693(t *testing.T) {
	qs := []question0693{
		{
			name:     "example 1",
			para0693: para0693{n: 5},
			ans0693:  ans0693{one: true},
		},
		{
			name:     "example 2",
			para0693: para0693{n: 7},
			ans0693:  ans0693{one: false},
		},
		{
			name:     "example 3",
			para0693: para0693{n: 11},
			ans0693:  ans0693{one: false},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 0693------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans0693, q.para0693
			output := hasAlternatingBits(p.n)
			fmt.Printf("[input]: n=%v       [output]:%v\n", p.n, output)
			if output != q.ans0693.one {
				t.Errorf("Expected %v, got %v", q.ans0693.one, output)
			}
		})
	}
	fmt.Printf("\n\n\n")
}
