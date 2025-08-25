package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3646 struct {
	name string
	para3646
	ans3646
}

// para is parameter
type para3646 struct {
	n int64
}

// ans is answer
type ans3646 struct {
	result int64
}

func Test_Problem3646(t *testing.T) {

	qs := []question3646{
		{
			name: "example 1",
			para3646: para3646{
				n: 2,
			},
			ans3646: ans3646{
				result: 22,
			},
		},
		{
			name: "example 2",
			para3646: para3646{
				n: 33,
			},
			ans3646: ans3646{
				result: 212,
			},
		},
		{
			name: "edge case - n = 0",
			para3646: para3646{
				n: 0,
			},
			ans3646: ans3646{
				result: 1,
			},
		},
		{
			name: "edge case - n = 1",
			para3646: para3646{
				n: 1,
			},
			ans3646: ans3646{
				result: 22,
			},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3646------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3646, q.para3646
			output := specialPalindrome(p.n)
			fmt.Printf("[input]: n=%v       [output]:%v\n", p.n, output)
			assert.Equal(t, q.ans3646.result, output, "specialPalindrome(%v) should return %v", p.n, q.ans3646.result)
		})
	}
	fmt.Printf("\n\n\n")
}
