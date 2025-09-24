package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question166 struct {
	name string
	para166
	ans166
}

// para is parameter
type para166 struct {
	numerator   int
	denominator int
}

// ans is answer
type ans166 struct {
	result string
}

func Test_Problem166(t *testing.T) {
	qs := []question166{
		{
			name:    "example 1",
			para166: para166{1, 2},
			ans166:  ans166{"0.5"},
		},
		{
			name:    "example 2",
			para166: para166{2, 1},
			ans166:  ans166{"2"},
		},
		{
			name:    "example 3",
			para166: para166{4, 333},
			ans166:  ans166{"0.(012)"},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 166------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans166, q.para166
			output := fractionToDecimal(p.numerator, p.denominator)
			fmt.Printf("[input]: numerator=%d, denominator=%d       [output]:%v\n", p.numerator, p.denominator, output)
			assert.Equal(t, q.ans166.result, output, "fractionToDecimal(%d, %d) should return %v, got %v", p.numerator, p.denominator, q.ans166.result, output)
		})
	}
	fmt.Printf("\n\n\n")
}
