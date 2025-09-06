package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type question3495 struct {
	name string
	para3495
	ans3495
}

// para is parameter
// one represents the first parameter
type para3495 struct {
	one [][]int
}

// ans is answer
// one represents the first answer
type ans3495 struct {
	one int64
}

func Test_Problem3495(t *testing.T) {

	qs := []question3495{
		{
			name:     "example 1",
			para3495: para3495{[][]int{{1, 2}, {2, 4}}},
			ans3495:  ans3495{3},
		},
		{
			name:     "example 2",
			para3495: para3495{[][]int{{2, 6}}},
			ans3495:  ans3495{4},
		},
	}

	fmt.Printf("------------------------Leetcode Problem 3495------------------------\n")

	for _, q := range qs {
		t.Run(q.name, func(t *testing.T) {
			_, p := q.ans3495, q.para3495
			output := minOperations(p.one)
			fmt.Printf("[input]:%v       [output]:%v\n", p, output)
			assert.Equal(t, q.ans3495.one, output, "minOperations(%v) should return %v, got %v", p.one, q.ans3495.one, output)
		})
	}
	fmt.Printf("\n\n\n")
}
